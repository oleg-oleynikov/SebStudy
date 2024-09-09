package infrastructure

import (
	// "SebStudy/infrastructure/eventsourcing"
	"encoding/json"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

type EsEventSerde struct {
}

func NewEsEventSerde() *EsEventSerde {
	return &EsEventSerde{}
}

func (e *EsEventSerde) Serialize(streamName string, event interface{}, m *EventMetadata) (*nats.Msg, error) {
	eventEnvelope := NewEventEnvelope(event, m) // Нужна норм сериализация
	dataBytes, err := json.Marshal(eventEnvelope)

	if err != nil {
		return nil, err
	}

	header := nats.Header{}

	return &nats.Msg{
		Subject: streamName,
		Data:    dataBytes,
		Header:  header,
	}, nil
}

func (e *EsEventSerde) Deserialize(data jetstream.Msg) (interface{}, *EventMetadata, error) {
	return nil, nil, nil
}
