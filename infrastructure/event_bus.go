package infrastructure

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

type EventBus interface {
	Publish(topic string, data interface{}) error
	Subscribe(topic string, handler nats.Handler) error
}

type EventBusNats struct {
	nc          *nats.EncodedConn
	subscribers map[string]*nats.Subscription
	mu          sync.Mutex
}

func NewEventBusNats(natsConn *nats.Conn) *EventBusNats {
	c, _ := nats.NewEncodedConn(natsConn, nats.JSON_ENCODER)

	return &EventBusNats{
		nc:          c,
		subscribers: make(map[string]*nats.Subscription),
	}
}

func (eb *EventBusNats) Subscribe(topic string, cb nats.Handler) error {
	eb.mu.Lock()
	defer eb.mu.Unlock()

	if _, ok := eb.subscribers[topic]; ok {
		return fmt.Errorf("already subscribed to topic %q", topic)
	}

	sub, err := eb.nc.Subscribe(topic, cb)
	if err != nil {
		return err
	}

	eb.subscribers[topic] = sub
	return nil
}

func (eb *EventBusNats) Publish(topic string, data interface{}) error {
	return eb.nc.Publish(topic, data)
}
