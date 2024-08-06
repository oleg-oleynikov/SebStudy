package infrastructure

import v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"

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

func NewEventMetadataFromCloudEvent(c *v1.CloudEvent) *EventMetadata {
	return &EventMetadata{
		EventId:   c.Id,
		EventType: c.Type,
	}
}

func NewEventMetadataFrom(m CommandMetadata) *EventMetadata {
	return nil
}
