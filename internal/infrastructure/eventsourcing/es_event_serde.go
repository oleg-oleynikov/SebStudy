package eventsourcing

import (
	"SebStudy/internal/infrastructure"
	"SebStudy/logger"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type EsEventSerde struct {
	log        logger.Logger
	typeMapper *TypeMapper
}

func NewEsEventSerde(log logger.Logger, typeMapper *TypeMapper) *EsEventSerde {
	return &EsEventSerde{
		log:        log,
		typeMapper: typeMapper,
	}
}

func GenerateUuidWithoutDashes() string {
	u, _ := uuid.NewV7()
	bytes, _ := u.MarshalBinary()

	uuidString := fmt.Sprintf("%x", bytes)

	return uuidString
}

func (e *EsEventSerde) Serialize(streamName string, event interface{}, m *infrastructure.EventMetadata) (*nats.Msg, error) {
	typeToData, err := e.typeMapper.GetTypeToData(infrastructure.GetValueType(event))
	if err != nil {
		e.log.Debugf("Failed to get type to data: %v", err)
		return nil, err
	}

	id := GenerateUuidWithoutDashes()

	name, jsonData := typeToData(event)
	dataBytes, err := json.Marshal(jsonData)
	if err != nil {
		e.log.Debugf("Failed to marshal event to json: %v", err)
		return nil, err
	}

	metadataBytes, err := json.Marshal(m)
	if err != nil {
		e.log.Debugf("Failed to marshal event metadata to json: %v", err)
		return nil, err
	}

	eventData := NewEventData(dataBytes, metadataBytes)
	bytes, err := json.Marshal(eventData)
	if err != nil {
		e.log.Debugf("Failed serialize to event data: %v", err)
		return nil, err
	}

	header := nats.Header{
		"eventType":   {name},
		"aggregateId": {m.AggregateId},
		"userId":      {m.UserId},
	}

	return &nats.Msg{
		Subject: fmt.Sprintf("%s.%s_%s", streamName, name, id),
		Data:    bytes,
		Header:  header,
	}, nil
}

func (e *EsEventSerde) Deserialize(data jetstream.Msg) (interface{}, *infrastructure.EventMetadata, error) {
	dataToType, err := e.typeMapper.GetDataToType(data.Headers().Get("eventType"))
	if err != nil {
		return nil, nil, err
	}

	eventData := EventData{}
	if err := json.Unmarshal(data.Data(), &eventData); err != nil {
		return nil, nil, err
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(eventData.DataBytes, &m); err != nil {
		return nil, nil, err
	}

	metadata := infrastructure.EventMetadata{}
	if err := json.Unmarshal(eventData.MetadataBytes, &metadata); err != nil {
		return nil, nil, err
	}

	return dataToType(m), &metadata, nil
}
