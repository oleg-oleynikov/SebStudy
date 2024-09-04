package eventsourcing

import (
	"SebStudy/infrastructure"
)

type EsAggregateStore struct {
	AggregateStore

	store EventStore
}

func NewEsAggregateStore(eventStore EventStore) *EsAggregateStore {
	return &EsAggregateStore{
		store: eventStore,
	}
}

func (s *EsAggregateStore) Save(a AggregateRoot, m infrastructure.CommandMetadata) error {
	changes := a.GetChanges()
	streamName := GetStreamName(a)
	err := s.store.AppendEvents(streamName, a.GetVersion(), m, changes)
	if err != nil {
		return err
	}
	a.ClearChanges()
	return nil
}

func (s *EsAggregateStore) Load(id string, a AggregateRoot) error {
	streamName := GetStreamName(a)

	events, err := s.store.LoadEvents(streamName)
	if err != nil {
		return err
	}

	a.Load(events)
	a.ClearChanges()
	return nil
}
