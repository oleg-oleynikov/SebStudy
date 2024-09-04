package eventsourcing

import "SebStudy/infrastructure"

type AggregateStore interface {
	Save(a AggregateRoot, m infrastructure.CommandMetadata) error
	Load(id string, a AggregateRoot) error
}
