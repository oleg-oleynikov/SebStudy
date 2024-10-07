package eventsourcing

import (
	"encoding/json"
	"fmt"
	"resume-server/internal/infrastructure"
	"resume-server/logger"

	"github.com/EventStore/EventStore-Client-Go/esdb"
	"github.com/gofrs/uuid"
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

func (e *EsEventSerde) Serialize(event interface{}, m *infrastructure.EventMetadata) (esdb.EventData, error) {
	typeToData, err := e.typeMapper.GetTypeToData(getValueType(event))
	if err != nil {
		e.log.Debugf("Failed to get type to data: %v", err)
		return esdb.EventData{}, err
	}

	id, _ := uuid.NewV7()

	name, jsonData := typeToData(event)
	dataBytes, err := json.Marshal(jsonData)
	if err != nil {
		e.log.Debugf("Failed to marshal event to json: %v", err)
		return esdb.EventData{}, err
	}

	metadataBytes, err := json.Marshal(m)
	if err != nil {
		e.log.Debugf("Failed to marshal event metadata to json: %v", err)
		return esdb.EventData{}, err
	}

	eventData := esdb.EventData{
		EventID:     id,
		EventType:   name,
		ContentType: esdb.JsonContentType,
		Data:        dataBytes,
		Metadata:    metadataBytes,
	}

	return eventData, nil
}

func (e *EsEventSerde) Deserialize(data *esdb.ResolvedEvent) (interface{}, *infrastructure.EventMetadata, error) {
	dataToType, err := e.typeMapper.GetDataToType(data.Event.EventType)
	if err != nil {
		return nil, nil, err
	}

	m := map[string]interface{}{}
	if err := json.Unmarshal(data.Event.Data, &m); err != nil {
		return nil, nil, err
	}

	metadata := infrastructure.EventMetadata{}
	if err := json.Unmarshal(data.Event.UserMetadata, &metadata); err != nil {
		return nil, nil, err
	}

	return dataToType(m), &metadata, nil
}
