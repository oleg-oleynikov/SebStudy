package infrastructure

import (
	"fmt"
	"reflect"
	"sync"
)

type EventToMap func(interface{}) (map[string]interface{}, error)
type MapToEvent func(map[string]interface{}) interface{}

type EventSerde struct {
	serializeHandlers   map[string]EventToMap
	deserializeHandlers map[reflect.Type]MapToEvent
}

var (
	once2     sync.Once
	instance2 *EventSerde
)

func GetEventSerdeInstance() *EventSerde {
	once2.Do(func() {
		es := &EventSerde{}
		es.serializeHandlers = make(map[string]EventToMap)
		es.deserializeHandlers = make(map[reflect.Type]MapToEvent)
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

func (es *EventSerde) GetMapToEvent(tEvent reflect.Type) (MapToEvent, error) {
	mapToEvent, ok := es.deserializeHandlers[tEvent]
	if !ok {
		return nil, fmt.Errorf("deserialize handler for type %s doesnt exist", tEvent)
	}

	return mapToEvent, nil
}

func (e *EventSerde) Serialize(event interface{}, metadata *EventMetadata) (map[string]interface{}, error) {
	return nil, nil
	// eventToMap, err := e.GetEventToMap(eventMessage.Metadata.EventType)
	// if err != nil {
	// 	return nil, err
	// }
	// // eventToMap(eventMessage)
	// return eventToMap(eventMessage)
}

func (e *EventSerde) Deserialize(data interface{}) (interface{}, *EventMetadata, error) {

	// mapToEvent, err := e.GetMapToEvent()
	return nil, nil, nil
}
