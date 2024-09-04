package eventsourcing

import (
	"SebStudy/domain/resume/events"
	"SebStudy/infrastructure"
	"SebStudy/logger"
	"SebStudy/ports/db_ports"

	"github.com/nats-io/nats.go"
)

type EventStore interface {
	AppendEvents(streamName string, version int, m infrastructure.CommandMetadata, events ...interface{}) error
	// AppendEventsToAny(streamName string, m infrastructure.CommandMetadata, events ...interface{}) error
	LoadEvents(streamName string) ([]interface{}, error)
	// LoadEventsFromStart(streamName string) ([]interface{}, error)

	// Create(streamName string) error
	// Delete(streamName string) error
}

type EsEventStore struct {
	log        logger.Logger
	eventBus   infrastructure.EventBus
	eventSerde *infrastructure.EsEventSerde
	writeRepo  db_ports.WriteModel
	imageStore *infrastructure.ImageStore
}

func NewEsEventStore(log logger.Logger, eventBus infrastructure.EventBus, eventSerde *infrastructure.EsEventSerde, writeRepo db_ports.WriteModel, imageStore *infrastructure.ImageStore) *EsEventStore {
	es := &EsEventStore{
		log:        log,
		eventBus:   eventBus,
		eventSerde: eventSerde,
		writeRepo:  writeRepo,
		imageStore: imageStore,
	}

	es.eventBus.Subscribe("resume.sended", func(event infrastructure.EventEnvelope) {
		resumeCreated, ok := (event.Event).(events.ResumeCreated)
		if !ok {
			log.Debugf("failed to cast event\n")
			return
		}

		imageUrl, err := imageStore.SaveImage(resumeCreated.Photo.GetPhoto())
		if err != nil {
			log.Debugf("failed to save image, %s\n", err)
			return
		}
		resumeCreated.Photo.SetUrl(imageUrl)
		// eventEnvelope := infrastructure.NewEventEnvelope(event.Event, event.Metadata)
		data, err := eventSerde.Serialize(event, nil)
		if err != nil {
			imageStore.DeleteImageByPath(imageUrl)
			return
		}
		if err := es.writeRepo.Save(data); err != nil {
			es.imageStore.DeleteImageByPath(imageUrl)
			return
		}
	})

	return es
}

func (es *EsEventStore) LoadEvents(aggregateId string) ([]interface{}, error) {
	events, err := es.writeRepo.Get(aggregateId)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (es *EsEventStore) AppendEvents(events []interface{}, m infrastructure.CommandMetadata) error {
	// if events == nil || len(events) == 0 {
	// 	return nil
	// }

	var msgs []*nats.Msg
	for _, i := range events {
		serializedEvent, err := es.eventSerde.Serialize(i, nil) // TODO: Доделать EventMetadata
		if err != nil {
			return err
		}

		msgs = append(msgs, serializedEvent)
	}

	for _, m := range msgs {
		es.writeRepo.Save(m)
	}

	return nil
}
