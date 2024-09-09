package eventsourcing

import (
	"SebStudy/infrastructure"
	"SebStudy/logger"
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/sirupsen/logrus"
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
	return es.prefix + "." + streamName
}

func (es *JetStreamEventStore) LoadEvents(streamName string) ([]interface{}, error) {
	cfg := jetstream.ConsumerConfig{
		Name:          "consumer." + streamName,
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckPolicy:     jetstream.AckNonePolicy,
	}

	return es.loadEvents(streamName, cfg)
}

func (es *JetStreamEventStore) loadEvents(streamName string, cfg jetstream.ConsumerConfig) ([]interface{}, error) {

	streamInfo, err := es.js.Stream(context.Background(), es.GetFullStreamName(streamName))
	if err != nil {
		return nil, fmt.Errorf("failed to get stream info: %w", err)
	}

	totalMes := int(streamInfo.CachedInfo().State.Msgs)
	logrus.Debugf("Ивентов в потоке %v", totalMes)

	cons, err := es.js.CreateOrUpdateConsumer(context.Background(), es.GetFullStreamName(streamName), cfg)
	if err != nil {
		return nil, err
	}

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
			logrus.Printf("Пока так: %v", msg) // Исправить

			event, _, err := es.serde.Deserialize(msg)
			if err != nil {
				return nil, err
			}

			events = append(events, event)

			ackMsg(cfg, msg)
		}

		if err := batch.Error(); err != nil {
			logrus.Infof("Ошибка при батче: %v", err) // Исправить

			return nil, fmt.Errorf("batch fetch error: %w", err) // Исправить
		}
	}

	logrus.Printf("В итоге на выходе из event store: %v", events) //Если заработает то заменить на logger

	return events, nil
}

func (es *JetStreamEventStore) AppendEvents(streamName string, version int, m infrastructure.CommandMetadata, events ...interface{}) error {
	options := jetstream.StreamConfig{
		Name:      streamName,
		Retention: jetstream.LimitsPolicy,
		Storage:   jetstream.FileStorage,
		Subjects:  []string{es.GetFullStreamName(fmt.Sprintf("%s.>", streamName))},
	}

	if events == nil {
		return nil
	}

	if version > 1 {
		events = events[version:]
	}

	es.log.Debugf("Events which will be published: \n%v", events)

	return es.appendEvents(streamName, options, m, events)
}

func (es *JetStreamEventStore) appendEvents(streamName string, o jetstream.StreamConfig, m infrastructure.CommandMetadata, events ...interface{}) error {

	es.js.UpdateStream(context.Background(), o)

	var msgs []*nats.Msg
	for _, i := range events {
		msg, err := es.serde.Serialize(streamName, i, infrastructure.NewEventMetadataFrom(m))
		if err != nil {
			return err
		}

		msgs = append(msgs, msg)
	}

	for _, msg := range msgs {
		es.log.Debugf("Subj publishing msg: %s", msg.Subject)
	}

	return nil
}

func ackMsg(cfg jetstream.ConsumerConfig, msg jetstream.Msg) {
	if cfg.AckPolicy != jetstream.AckNonePolicy {
		msg.Ack()
	}
}
