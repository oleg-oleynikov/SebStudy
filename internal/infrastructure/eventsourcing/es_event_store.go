package eventsourcing

import (
	"context"
	"errors"
	"fmt"
	"io"
	"math"
	"resume-server/internal/infrastructure"
	"resume-server/logger"

	"github.com/EventStore/EventStore-Client-Go/esdb"
)

type EsEventStore struct {
	EventStore

	log    logger.Logger
	client *esdb.Client
	serde  EventSerde
	prefix string
}

func NewEsEventStore(log logger.Logger, client *esdb.Client, serde EventSerde, prefix string) *EsEventStore {
	return &EsEventStore{
		log:    log,
		client: client,
		serde:  serde,
		prefix: prefix,
	}
}

func (es *EsEventStore) GetFullStreamName(streamName string) string {
	return fmt.Sprintf("%s%s", es.prefix, streamName)
}

func (es *EsEventStore) AppendEvents(streamName string, version int, m infrastructure.CommandMetadata, events ...interface{}) error {
	options := esdb.AppendToStreamOptions{}
	if version == -1 {
		options.ExpectedRevision = esdb.NoStream{}
	} else {
		options.ExpectedRevision = esdb.Revision(uint64(version))
	}

	if len(events) == 0 {
		return nil
	}

	var eventData []esdb.EventData
	for _, e := range events {
		ed, err := es.serde.Serialize(e, infrastructure.NewEventMetadataFrom(m))
		if err != nil {
			return err
		}

		eventData = append(eventData, ed)
	}

	_, err := es.client.AppendToStream(context.Background(), es.GetFullStreamName(streamName), options, eventData...)
	return err
}

func (es *EsEventStore) LoadEvents(streamName string) ([]interface{}, error) {
	options := esdb.ReadStreamOptions{
		From:      esdb.Start{},
		Direction: esdb.Forwards,
	}

	events := make([]interface{}, 0)
	stream, err := es.client.ReadStream(context.Background(), es.GetFullStreamName(streamName), options, math.MaxInt64)
	if err != nil {
		return nil, err
	}

	for {
		event, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return nil, err
		}

		e, _, err := es.serde.Deserialize(event)
		if err != nil {
			return nil, err
		}

		events = append(events, e)
	}
	return events, nil
}
