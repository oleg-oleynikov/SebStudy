package eventsourcing

import "resume-server/internal/infrastructure"

type AggregateStore interface {
	Save(a AggregateRoot, m infrastructure.CommandMetadata) error
	Load(aggregateId string, a AggregateRoot) error
}
