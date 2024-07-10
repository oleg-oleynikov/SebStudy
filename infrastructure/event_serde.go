package infrastructure

import "sync"

type EventToMap func() error

type EventSerde struct {
}

var (
	once2     sync.Once
	instance2 *EventSerde
)

func GetEventSerdeInstance() *EventSerde {
	once2.Do(func() {
		es := &EventSerde{}
		instance2 = es
	})

	return instance2
}

// Будет переводить в формат для записи event log
func (e *EventSerde) Serialize(event interface{}, metadata EventMetadata) (interface{}, error) {

	return nil, nil
}

// Ну логично при получении из бд события будут проходить через это
func (e *EventSerde) Deserialize(data interface{}) (interface{}, *EventMetadata, error) {
	return nil, nil, nil
}
