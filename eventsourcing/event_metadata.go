package eventsourcing

import (
	"time"
)

type EventMetadata struct {
	EventId string
	Time    time.Time
}

func NewEventMetadata(EventId string, Time time.Time) EventMetadata {
	return EventMetadata{
		EventId: EventId,
		Time:    Time,
	}

}

// TODO: Доделать чтобы делал метадату из CommandMetadata
func NewEventMetadataFrom() EventMetadata {
	return EventMetadata{}
}
