package infrastructure

type EventMetadata struct {
	AggregateId string
	UserId      string
}

func NewEventMetadata(aggregateId string) *EventMetadata {
	return &EventMetadata{
		AggregateId: aggregateId,
	}
}

func NewEventMetadataFrom(m CommandMetadata) *EventMetadata {
	md := &EventMetadata{
		AggregateId: m.AggregateId,
		UserId:      m.UserId,
	}

	return md
}
