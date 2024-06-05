package events

import (
	// eventsourcing "SebStudy/event_sourcing"
	"time"
)

type ResumeSended struct {
	// eventsourcing.DomainEvent

	AggregateRootId string
	AggregateType   string
	Date            time.Time
}
