package projections

import (
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/logger"
	"context"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type SubscriptionManager struct {
	log           logger.Logger
	js            jetstream.JetStream
	serde         eventsourcing.EventSerde
	subscriptions []Subscription
}

func NewSubscriptionManager(log logger.Logger, nc *nats.Conn, serde eventsourcing.EventSerde, subs ...Subscription) *SubscriptionManager {
	js, err := jetstream.New(nc)
	if err != nil {
		log.Fatalf("Failed to get jetstream for subManager: %v", err)
	}

	return &SubscriptionManager{
		log:           log,
		js:            js,
		serde:         serde,
		subscriptions: subs,
	}
}

func (m *SubscriptionManager) Start(ctx context.Context) error {

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
	}

	consumer, err := m.js.CreateOrUpdateConsumer(ctx, "projection_stream", consumerConfig)
	if err != nil {
		m.log.Debugf("Failed to create or update consumer for projection: %v", err)
		return err
	}

	go func(cons jetstream.Consumer) {
		for {

			if cons.CachedInfo().NumPending > 0 {

				batch, err := cons.Fetch(10, jetstream.FetchMaxWait(5*time.Second))
				if err != nil {
					m.log.Fatalf("Failed to batch messages for projection: %v", err)
					break
				}

				if len(batch.Messages()) == 0 {
					continue
				}

				m.log.Debugf("Пришло чет") // DEBUGGGG

				for msg := range batch.Messages() {
					m.log.Debugf("Received message: %s", string(msg.Data())) // DEBUGGGG
					event, _, err := m.serde.Deserialize(msg)
					if err != nil {
						msg.Nak()
						m.log.Fatalf("Failed to deserialize msg: %v", err)
					}

					m.log.Debugf("Event: %v", event) // DEBUGGGG

					// for _, s := range m.subscriptions {
					// 	s.Project(event, *metadata)
					// }

					msg.Ack()
				}
			} else {
				time.Sleep(time.Millisecond * 10)
			}
		}
	}(consumer)

	return nil
}
