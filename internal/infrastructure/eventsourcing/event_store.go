package eventsourcing

import (
	"SebStudy/internal/infrastructure"
)

type EventStore interface {
	AppendEvents(streamName string, version int, m infrastructure.CommandMetadata, events ...interface{}) error
	LoadEvents(streamName string) ([]interface{}, error)
}