package infrastructure

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type DomainEvent interface {
}

type EventMetadata struct {
	EventId   string
	EventType string
	// Version   int
}

func NewEventMetadata(eventId, eventType string) *EventMetadata {
	return &EventMetadata{
		EventId:   eventId,
		EventType: eventType,
	}
}

// func (em *EventMetadata) SetVersion(version int) {
// 	em.Version = version
// }

func NewEventMetadataFromCloudEvent(c cloudevents.Event) *EventMetadata {
	return &EventMetadata{
		EventId:   c.ID(),
		EventType: c.Type(),
	}
}

func NewEventMetadataFrom(m CommandMetadata) *EventMetadata {
	return nil
}
