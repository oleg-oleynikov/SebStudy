package infrastructure

import (
	"fmt"
	"sync"
)

type EventSerde interface {
	Serialize(event interface{}, metadata EventMetadata) (map[string]interface{}, error)
	Deserialize(data map[string]interface{}) (EventEnvelope, error)
}

type EventToMap func(interface{}) (map[string]interface{}, error)
type MapToEvent func(map[string]interface{}) interface{}

type EsEventSerde struct {
	serializeHandlers   map[string]EventToMap
	deserializeHandlers map[string]MapToEvent
}

var (
	once2     sync.Once
	instance2 *EsEventSerde
)

func GetEsEventSerdeInstance() *EsEventSerde {
	once2.Do(func() {
		es := &EsEventSerde{}
		es.serializeHandlers = make(map[string]EventToMap)
		es.deserializeHandlers = make(map[string]MapToEvent)
		instance2 = es
	})

	return instance2
}

func (es *EsEventSerde) GetEventToMap(eventType string) (EventToMap, error) {
	eventToMap, ok := es.serializeHandlers[eventType]
	if !ok {
		return nil, fmt.Errorf("serialize handler for type %s doesnt exist", eventType)
	}

	return eventToMap, nil
}

func (es *EsEventSerde) GetMapToEvent(tEvent string) (MapToEvent, error) {
	mapToEvent, ok := es.deserializeHandlers[tEvent]
	if !ok {
		return nil, fmt.Errorf("deserialize handler for type %s doesnt exist", tEvent)
	}

	return mapToEvent, nil
}

func (e *EsEventSerde) Serialize(event interface{}, metadata EventMetadata) (map[string]interface{}, error) {
	return nil, nil
}

func (e *EsEventSerde) Deserialize(data map[string]interface{}) (interface{}, *EventMetadata, error) {
	eventType, ok := data["eventType"]
	if !ok {
		return nil, nil, fmt.Errorf("doesnt exist eventType")
	}

	value, ok := eventType.(string)
	if !ok {
		return nil, nil, fmt.Errorf("event type not a string")
	}

	_, err := e.GetMapToEvent(value)
	if err != nil {
		return nil, nil, err
	}
	return nil, nil, nil
}
