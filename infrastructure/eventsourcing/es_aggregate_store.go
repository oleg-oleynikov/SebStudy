package eventsourcing

import (
	"SebStudy/infrastructure"
	"SebStudy/logger"
)

type EsAggregateStore struct {
	AggregateStore

	log   logger.Logger
	store EventStore
}

func NewEsAggregateStore(log logger.Logger, eventStore EventStore) *EsAggregateStore {
	return &EsAggregateStore{
		log:   log,
		store: eventStore,
	}
}

func (s *EsAggregateStore) Save(a AggregateRoot, m infrastructure.CommandMetadata) error {
	changes := a.GetChanges()
	streamName := GetStreamName(a)

	err := s.store.AppendEvents(streamName, a.GetVersion(), m, changes...)
	if err != nil {
		return err
	}

	a.ClearChanges()
	return nil
}

func (s *EsAggregateStore) Load(aggregateId string, a AggregateRoot) error {
	streamName := GetStreamNameWithId(a, aggregateId)

	events, err := s.store.LoadEvents(streamName)
	if err != nil {
		return err
	}

	a.Load(events)
	a.ClearChanges()
	return nil
}
