package infrastructure

type EventMessage struct {
	Event    interface{}
	Metadata EventMetadata
	Version  int
}

func NewEventMessage(event interface{}, metadata EventMetadata, version int) *EventMessage {
	return &EventMessage{
		Event:    event,
		Metadata: metadata,
		Version:  version,
	}
}
