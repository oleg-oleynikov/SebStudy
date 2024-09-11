package eventsourcing

import "SebStudy/infrastructure"

type AggregateStore interface {
	Save(a AggregateRoot, m infrastructure.CommandMetadata) error
	Load(aggregateId string, a AggregateRoot) error
}
