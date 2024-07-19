package infrastructure

import (
	"fmt"
	"sync"
)

type EventToMap func(interface{}) (map[string]interface{}, error)
type MapToEvent func(map[string]interface{}) interface{}

type EventSerde struct {
	serializeHandlers   map[string]EventToMap
	deserializeHandlers map[string]MapToEvent
}

var (
	once2     sync.Once
	instance2 *EventSerde
)

func GetEventSerdeInstance() *EventSerde {
	once2.Do(func() {
		es := &EventSerde{}
		es.serializeHandlers = make(map[string]EventToMap)
		es.deserializeHandlers = make(map[string]MapToEvent)
		instance2 = es
	})

	return instance2
}

func (es *EventSerde) GetEventToMap(eventType string) (EventToMap, error) {
	eventToMap, ok := es.serializeHandlers[eventType]
	if !ok {
		return nil, fmt.Errorf("serialize handler for type %s doesnt exist", eventType)
	}

	return eventToMap, nil
}

func (es *EventSerde) GetMapToEvent(tEvent string) (MapToEvent, error) {
	mapToEvent, ok := es.deserializeHandlers[tEvent]
	if !ok {
		return nil, fmt.Errorf("deserialize handler for type %s doesnt exist", tEvent)
	}

	return mapToEvent, nil
}

func (e *EventSerde) Serialize(event interface{}, metadata EventMetadata) (map[string]interface{}, error) {
	return nil, nil
	// eventToMap, err := e.GetEventToMap(eventMessage.Metadata.EventType)
	// if err != nil {
	// 	return nil, err
	// }
	// // eventToMap(eventMessage)
	// return eventToMap(eventMessage)
}

func (e *EventSerde) Deserialize(data map[string]interface{}) (interface{}, *EventMetadata, error) {
	// mapToEvent, err := e.GetMapToEvent()
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
