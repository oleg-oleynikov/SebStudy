package infrastructure

type EventEnvelope struct {
	Event    interface{}
	Metadata EventMetadata
	// Version  int
}

func NewEventEnvelope(event interface{}, metadata EventMetadata, version int) *EventEnvelope {
	return &EventEnvelope{
		Event:    event,
		Metadata: metadata,
		// Version:  version,
	}
}
