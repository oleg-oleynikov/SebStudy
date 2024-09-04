package infrastructure

type EventMetadata struct {
	addOptions map[string]string
}

func NewEventMetadata(eventId string) *EventMetadata {
	return &EventMetadata{
		addOptions: make(map[string]string),
	}
}

func NewEventMetadataFrom(m CommandMetadata) *EventMetadata {
	return nil
}
