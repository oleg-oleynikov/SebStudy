package infrastructure

import (
	"SebStudy/domain/resume/events"
	"SebStudy/ports/db_ports"
	"log"
)

type EventStore interface {
	LoadEvents(aggregateId string) ([]interface{}, error)
	AppendEvents(CommandMetadata, int, ...interface{}) error
}

type EsEventStore struct {
	eventBus   *EventBus
	eventSerde *EventSerde
	writeRepo  db_ports.WriteModel
	imageStore *ImageStore
}

func NewEsEventStore(eventBus *EventBus, eventSerde *EventSerde, writeRepo db_ports.WriteModel, imageStore *ImageStore) *EsEventStore {
	es := &EsEventStore{
		eventBus:   eventBus,
		eventSerde: eventSerde,
		writeRepo:  writeRepo,
		imageStore: imageStore,
	}

	es.eventBus.Subscribe("resume.sended", func(event *EventMessage[events.ResumeSended]) {
		log.Println(event)
		imageUrl, err := imageStore.SaveImage(event.Event.Photo.GetPhoto())
		if err != nil {
			log.Printf("failed to save image, %s\n", err)
			return
		}
		event.Event.Photo.SetUrl(imageUrl)
		data, err := es.eventSerde.Serialize(event.Event, event.Metadata)
		if err != nil {
			es.imageStore.DeleteImageByPath(imageUrl)
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
	// return nil, nil
	events, err := es.writeRepo.GetByAggregateId(aggregateId)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (es *EsEventStore) AppendEvents(m CommandMetadata, version int, events ...interface{}) error {
	return nil
}
