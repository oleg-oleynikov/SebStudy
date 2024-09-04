package infrastructure

import (
	// "SebStudy/infrastructure/eventsourcing"
	"encoding/json"
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type EventToMap func(event interface{}, md *CommandMetadata) (*nats.Msg, error)
type MapToEvent func(msg *nats.Msg) (interface{}, *CommandMetadata, error)

type EsEventSerde struct {
	// tm *eventsourcing.TypeMapper

	// serializeHandlers   map[string]EventToMap
	// deserializeHandlers map[string]MapToEvent
}

func NewEsEventSerde() *EsEventSerde {
	return &EsEventSerde{
		// tm: tm,
		// serializeHandlers:   make(map[string]EventToMap),
		// deserializeHandlers: make(map[string]MapToEvent),
	}
}

// func (es *EsEventSerde) GetEventToMap(eventType string) (EventToMap, error) {
// 	// eventToMap, ok := es.serializeHandlers[eventType]
// 	if !ok {
// 		return nil, fmt.Errorf("serialize handler for type %s doesnt exist", eventType)
// 	}

// 	return eventToMap, nil
// }

// func (es *EsEventSerde) GetMapToEvent(tEvent string) (MapToEvent, error) {
// 	mapToEvent, ok := es.deserializeHandlers[tEvent]
// 	if !ok {
// 		return nil, fmt.Errorf("deserialize handler for type %s doesnt exist", tEvent)
// 	}

// 	return mapToEvent, nil
// }

func (e *EsEventSerde) Serialize(event interface{}, m *EventMetadata) (*nats.Msg, error) {
	// typeToData, err := e.tm.GetTypeToData(GetValueType(event))

	// if err != nil {
	// 	logger.Logger.Debugf("Failed to serialize event to nats message: %v", err)
	// 	return nil, fmt.Errorf("failed to serialize event to nats msg %s", err)
	// }

	eventEnvelope := NewEventEnvelope(event, m)
	// eventType, _ := typeToData(event)
	dataBytes, err := json.Marshal(eventEnvelope)

	if err != nil {
		return nil, err
	}

	return &nats.Msg{
		Data: dataBytes,
		Header: nats.Header{
			"event-type": []string{""}, // Доделать
		},
	}, nil
}

func (e *EsEventSerde) Deserialize(data jetstream.Msg) (interface{}, *CommandMetadata, error) {

	eventType := data.Headers().Get("event-type")

	if eventType == "" {
		return nil, nil, fmt.Errorf("doesnt exist eventType")
	}

	// _, err := e.GetMapToEvent(eventType)

	// if err != nil {
	// 	return nil, nil, err
	// }

	return nil, nil, nil
}
