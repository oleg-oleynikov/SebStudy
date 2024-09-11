package eventsourcing

import (
	"SebStudy/internal/infrastructure"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type EventSerde interface {
	Serialize(streamName string, event interface{}, md *infrastructure.EventMetadata) (*nats.Msg, error)
	Deserialize(data jetstream.Msg) (interface{}, *infrastructure.EventMetadata, error)
}
