package eventsourcing

import (
	"SebStudy/internal/infrastructure"
	"SebStudy/logger"
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type JetStreamEventStore struct {
	EventStore

	log    logger.Logger
	nc     *nats.Conn
	js     jetstream.JetStream
	serde  EventSerde
	prefix string
}

func NewJetStreamEventStore(appLogger logger.Logger, nc *nats.Conn, serde EventSerde, prefix string) *JetStreamEventStore {
	js, err := jetstream.New(nc)
	if err != nil {
		appLogger.Fatalf("failed to get nats jetstream: %v", err)
		return nil
	}

	if js == nil {
		appLogger.Fatalf("JetStream is nil")
		return nil
	}

	if serde == nil {
		appLogger.Fatalf("Serde is nil")
		return nil
	}

	return &JetStreamEventStore{
		log:    appLogger,
		nc:     nc,
		js:     js,
		serde:  serde,
		prefix: prefix,
	}
}

func (es *JetStreamEventStore) GetFullStreamName(streamName string) string {
	return fmt.Sprintf("%s_%s", es.prefix, streamName)
}

func (es *JetStreamEventStore) GetFullSubject(streamName string) string {
	return fmt.Sprintf("%s.%s.>", es.prefix, streamName)
}

func (es *JetStreamEventStore) LoadEvents(streamName string) ([]interface{}, error) {
	cfg := jetstream.ConsumerConfig{
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckPolicy:     jetstream.AckNonePolicy,
	}

	return es.loadEvents(streamName, cfg)
}

func (es *JetStreamEventStore) loadEvents(streamName string, cfg jetstream.ConsumerConfig) ([]interface{}, error) {

	stream, err := es.js.Stream(context.Background(), es.GetFullStreamName(streamName))
	if err != nil {
		return nil, fmt.Errorf("failed to get stream info: %w", err)
	}

	streamInfo, err := stream.Info(context.TODO())
	if err != nil {
		return nil, err
	}

	totalMes := int(streamInfo.State.Msgs)

	cons, err := es.js.CreateOrUpdateConsumer(context.Background(), es.GetFullStreamName(streamName), cfg)
	if err != nil {
		es.log.Debugf("Failed to create or update consumer: %v", err)
		return nil, err
	}

	defer func() {
		err := es.js.DeleteConsumer(context.Background(), es.GetFullStreamName(streamName), cons.CachedInfo().Name)
		if err != nil {
			es.log.Debugf("Failed to delete ephemeral consumer: %v", err)
		}
	}()

	batchSize := 100
	var events []interface{}

	for i := totalMes; i > 0; i -= batchSize {
		fetchSize := batchSize

		if i < batchSize {
			fetchSize = i
		}

		batch, err := cons.Fetch(fetchSize, jetstream.FetchMaxWait(3*time.Second))
		if err != nil {
			return nil, err
		}

		for msg := range batch.Messages() {
			event, _, err := es.serde.Deserialize(msg)
			if err != nil {
				return nil, err
			}

			events = append(events, event)

			ackMsg(cfg, msg)
		}

		if err := batch.Error(); err != nil {
			es.log.Debugf("Batch fetch error: %v", err)
			return nil, fmt.Errorf("batch fetch error: %v", err)
		}
	}

	return events, nil
}

func (es *JetStreamEventStore) AppendEvents(streamName string, version int, m infrastructure.CommandMetadata, events ...interface{}) error {
	subj := fmt.Sprintf("%s.>", es.GetFullStreamName(streamName))
	options := jetstream.StreamConfig{
		Name:      es.GetFullStreamName(streamName),
		Retention: jetstream.LimitsPolicy,
		Storage:   jetstream.FileStorage,
		Subjects:  []string{subj},
		RePublish: &jetstream.RePublish{
			Source:      subj,
			Destination: "projection.>",
		},
	}

	if events == nil {
		return nil
	}

	// if version > 1 {
	// 	events = events[version:]
	// }

	return es.appendEvents(streamName, options, m, events...)
}

func (es *JetStreamEventStore) appendEvents(streamName string, o jetstream.StreamConfig, m infrastructure.CommandMetadata, events ...interface{}) error {

	_, err := es.js.CreateOrUpdateStream(context.Background(), o)
	if err != nil {
		es.log.Debugf("Failed to create or update stream: %v", err)
	}

	var msgs []*nats.Msg
	for _, i := range events {
		msg, err := es.serde.Serialize(streamName, i, infrastructure.NewEventMetadataFrom(m))
		msg.Subject = es.prefix + "_" + msg.Subject
		if err != nil {
			es.log.Debugf("Failed to serialize to nats msg: %v", err)
			return err
		}

		msgs = append(msgs, msg)
	}

	for _, msg := range msgs {
		_, err := es.js.PublishMsgAsync(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func ackMsg(cfg jetstream.ConsumerConfig, msg jetstream.Msg) {
	if cfg.AckPolicy == jetstream.AckExplicitPolicy {
		msg.Ack()
	}
}
