package eventsourcing

import (
	"resume-server/internal/infrastructure"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

type EventSerde interface {
	Serialize(event interface{}, md *infrastructure.EventMetadata) (esdb.EventData, error)
	Deserialize(data *esdb.ResolvedEvent) (interface{}, *infrastructure.EventMetadata, error)
}
