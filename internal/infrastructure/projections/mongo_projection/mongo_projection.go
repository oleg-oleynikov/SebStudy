package mongo_projection

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/infrastructure"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
	"context"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type mongoProjection struct {
	log       logger.Logger
	cfg       config.Config
	js        jetstream.JetStream
	serde     eventsourcing.EventSerde
	mongoRepo repository.ResumeMongoRepository
}

func NewMongoProjection(log logger.Logger, cfg config.Config, nc *nats.Conn, serde eventsourcing.EventSerde, mongoRepo repository.ResumeMongoRepository) *mongoProjection {
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("(mongoProjection) Failed to get jetstream: %v", err)
	}

	return &mongoProjection{
		log:       log,
		cfg:       cfg,
		js:        js,
		serde:     serde,
		mongoRepo: mongoRepo,
	}
}

func (m *mongoProjection) Start(ctx context.Context) error {

	streamConfig := jetstream.StreamConfig{
		Name:      "projection_stream",
		Subjects:  []string{"projection.>"},
		Retention: jetstream.WorkQueuePolicy,
		Storage:   jetstream.MemoryStorage,
	}

	_, err := m.js.CreateOrUpdateStream(ctx, streamConfig)
	if err != nil {
		m.log.Fatalf("Failed to create or update projection stream: %v", err)
	}

	consumerConfig := jetstream.ConsumerConfig{
		Name:          "projection_consumer",
		DeliverPolicy: jetstream.DeliverAllPolicy,
		AckPolicy:     jetstream.AckExplicitPolicy,
		AckWait:       45 * time.Second,
		ReplayPolicy:  jetstream.ReplayInstantPolicy,
		Durable:       "projection_consumer",
	}

	consumer, err := m.js.CreateOrUpdateConsumer(ctx, "projection_stream", consumerConfig)
	if err != nil {
		m.log.Debugf("Failed to create or update consumer for projection: %v", err)
		return err
	}

	go func() {
		for {
			consumerInfo, err := consumer.Info(context.Background())
			if err != nil {
				m.log.Fatalf("Failed to get consumer info: %v", err)
			}

			if consumerInfo.NumPending > 0 {
				batch, err := consumer.Fetch(10, jetstream.FetchMaxWait(5*time.Second))
				if err != nil {
					m.log.Fatalf("Failed to batch messages for projection: %v", err)
					break
				}

				for msg := range batch.Messages() {
					event, md, err := m.serde.Deserialize(msg)
					if err != nil {
						msg.Nak()
						m.log.Fatalf("Failed to deserialize msg: %v", err)
					}

					if err = m.When(ctx, event, md); err != nil {
						m.log.Debugf("%v", err)
					}

					msg.Ack()
				}
			} else {
				time.Sleep(time.Millisecond * 30)
			}
		}
	}()

	return nil
}

func (m *mongoProjection) When(ctx context.Context, event interface{}, md *infrastructure.EventMetadata) error {
	switch event := event.(type) {
	case events.ResumeCreated:
		return m.onResumeCreate(ctx, event, md)
	case events.ResumeChanged:
		return m.onResumeChanged(ctx, event, md)
	default:
		m.log.Debugf("(mongoProjection) [When unknown EventType] eventType: {%s}")
		return fmt.Errorf("invalid event type")
	}
}
