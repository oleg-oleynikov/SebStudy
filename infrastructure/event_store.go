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
		imageUrl, err := imageStore.SaveImage(event.Event.Photo.GetPhoto())
		if err != nil {
			log.Printf("failed to save image, %s\n", err)
			return
		}
		event.Event.Photo.SetUrl(imageUrl)
		data, err := es.eventSerde.Serialize(event.Event, event.Metadata) // БЛЯТЬ ЕБЛО ТУПОЕ ДОДелОАЙ СЕРИАЛИЗАЦИЮ НОРМ
		if err != nil {
			es.imageStore.DeleteImageByPath(imageUrl)
			return
		}
		if err := es.writeRepo.Save(data); err != nil {
			es.imageStore.DeleteImageByPath(imageUrl)
			return
		}
		log.Printf("Ивент сохранился с id: %s\n", event.Metadata.EventId)
	})

	return es
}

func (es *EsEventStore) LoadEvents(aggregateId string) ([]interface{}, error) {
	// return nil, nil
	events, err := es.writeRepo.Get(aggregateId)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (es *EsEventStore) AppendEvents(m CommandMetadata, version int, events ...interface{}) error {
	return nil
}
