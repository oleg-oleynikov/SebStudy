package infrastructure

import (
	"fmt"
	"sync"

	"github.com/nats-io/nats.go"
)

type EventBus struct {
	nc          *nats.Conn
	subscribers map[string]*nats.Subscription
	mu          sync.Mutex
}

func NewEventBus(natsURL string) (*EventBus, error) {
	nc, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	return &EventBus{
		nc:          nc,
		subscribers: make(map[string]*nats.Subscription),
	}, nil
}

func (eb *EventBus) Subscribe(topic string, cb nats.MsgHandler) error {
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

func (eb *EventBus) Publish(topic string, data []byte) error {
	return eb.nc.Publish(topic, data)
}
