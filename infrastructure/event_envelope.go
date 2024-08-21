package infrastructure

type EventEnvelope[T any] struct {
	Event    T
	Metadata EventMetadata
	// Version  int
}

func NewEventEnvelope[T any](event T, metadata EventMetadata, version int) *EventEnvelope[T] {
	return &EventEnvelope[T]{
		Event:    event,
		Metadata: metadata,
		// Version:  version,
	}
}
