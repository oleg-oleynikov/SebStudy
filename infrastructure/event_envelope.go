package infrastructure

type EventEnvelope struct {
	Event    interface{}
	Metadata *EventMetadata
	// Version  int
}

func NewEventEnvelope(event interface{}, metadata *EventMetadata) *EventEnvelope {
	return &EventEnvelope{
		Event:    event,
		Metadata: metadata,
		// Version:  version,
	}
}

func (e *EventEnvelope) Unwrap() (interface{}, *EventMetadata) {
	return e.Event, e.Metadata
}
