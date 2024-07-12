package infrastructure

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type DomainEvent interface {
}

type EventMetadata struct {
	EventId   string
	EventType string
}

func NewEventMetadata(eventId, eventType string) *EventMetadata {
	return &EventMetadata{
		EventId:   eventId,
		EventType: eventType,
	}

}

func NewEventMetadataFromCloudEvent(c cloudevents.Event) *EventMetadata {
	return &EventMetadata{
		EventId:   c.ID(),
		EventType: c.Type(),
	}
}
