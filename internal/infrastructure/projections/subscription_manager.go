package projections

import (
	"SebStudy/internal/infrastructure/eventsourcing"

	"github.com/nats-io/nats.go/jetstream"
)

type SubscriptionsManager struct {
	js    jetstream.Stream
	serde *eventsourcing.EsEventSerde
}
