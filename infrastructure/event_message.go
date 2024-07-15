package infrastructure

type EventMessage[T any] struct {
	Event    T
	Metadata EventMetadata
	Version  int
}

func NewEventMessage[T any](event T, metadata EventMetadata, version int) *EventMessage[T] {
	return &EventMessage[T]{
		Event:    event,
		Metadata: metadata,
		Version:  version,
	}
}
