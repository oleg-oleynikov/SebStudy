package infrastructure

import "fmt"

type JetStreamEventStore struct {
}

func NewJetStreamEventStore() *JetStreamEventStore {
	return &JetStreamEventStore{}
}

func (es *JetStreamEventStore) LoadEvents(streamName string, version int) ([]interface{}, error) {
	return nil, fmt.Errorf("not impl")
}

func (es *JetStreamEventStore) LoadEventsFromStart(streamName string) ([]interface{}, error) {
	return nil, fmt.Errorf("not impl")
}

func (es *JetStreamEventStore) AppendEvents(streamName string, version int, m CommandMetadata, events ...interface{}) error {
	return fmt.Errorf("not impl")
}

func (es *JetStreamEventStore) AppendEventsToAny(streamName string, m CommandMetadata, events ...interface{}) error {
	return fmt.Errorf("not impl")
}
