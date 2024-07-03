package eventsourcing

import (
	"SebStudy/infrastructure"
	"time"
)

type EventMetadata struct {
	EventId   string
	Timestamp time.Time
}

func NewEventMetadata(eventId string, timestamp time.Time) *EventMetadata {
	return &EventMetadata{
		EventId:   eventId,
		Timestamp: timestamp,
	}

}

func NewEventMetadataFrom(c infrastructure.CommandMetadata) *EventMetadata {
	return &EventMetadata{
		EventId:   c.CloudEvent.ID(),
		Timestamp: c.CloudEvent.Time(),
	}
}
