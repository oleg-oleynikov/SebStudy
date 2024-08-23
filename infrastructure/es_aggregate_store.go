package infrastructure

type EsAggregateStore struct {
	AggregateStore

	store EventStore
}

func NewEsAggregateStore() *EsAggregateStore {
	return &EsAggregateStore{}
}

func (s *EsAggregateStore) Save(a AggregateRoot, m CommandMetadata) error {
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
	version := -1
	streamName := GetStreamName(a)

	events, err := s.store.LoadEvents(streamName, version)
	if err != nil {
		return err
	}

	a.Load(events)
	a.ClearChanges()
	return nil
}
