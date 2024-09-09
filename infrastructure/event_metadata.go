package infrastructure

type EventMetadata struct {
	// AggregateId string
	addOptions map[string]string
}

func NewEventMetadata(eventId string) *EventMetadata {
	return &EventMetadata{
		addOptions: make(map[string]string),
	}
}

func (md *EventMetadata) AddOption(key string, value string) {
	if _, ex := md.addOptions[key]; ex {
		return
	}

	md.addOptions[key] = value
}

func NewEventMetadataFrom(m CommandMetadata) *EventMetadata {
	md := &EventMetadata{
		// AggregateId: m.AggregateId,
	}

	if m.UserId != "" {
		md.AddOption("user", m.UserId)
	}

	return md
}
