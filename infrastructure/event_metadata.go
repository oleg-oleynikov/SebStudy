package infrastructure

type EventMetadata struct {
	AggregateId string
	UserId      string
	// AddOptions  map[string]string
}

func NewEventMetadata(aggregateId string) *EventMetadata {
	return &EventMetadata{
		AggregateId: aggregateId,
		// AddOptions:  make(map[string]string),
	}
}

// func (md *EventMetadata) AddOption(key string, value string) {
// if _, ex := md.AddOptions[key]; ex {
// 	return
// }

// md.AddOptions[key] = value
// }

func NewEventMetadataFrom(m CommandMetadata) *EventMetadata {
	md := &EventMetadata{
		// AggregateId: m.AggregateId,
	}

	if m.UserId != "" {
		// md.AddOption("user", m.UserId)
		md.UserId = m.UserId
	}

	return md
}
