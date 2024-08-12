package infrastructure

import (
	"SebStudy/domain/resume/events"
	"SebStudy/ports/db_ports"
	"log"
)

type EventStore interface {
	LoadEvents(aggregateId string) ([]interface{}, error)
	AppendEvents(CommandMetadata, ...interface{}) error
}

type EsEventStore struct {
	eventBus   EventBus
	eventSerde *EventSerde
	writeRepo  db_ports.WriteModel
	imageStore *ImageStore
}

func NewEsEventStore(eventBus EventBus, eventSerde *EventSerde, writeRepo db_ports.WriteModel, imageStore *ImageStore) *EsEventStore {
	es := &EsEventStore{
		eventBus:   eventBus,
		eventSerde: eventSerde,
		writeRepo:  writeRepo,
		imageStore: imageStore,
	}

	es.eventBus.Subscribe("resume.sended", func(event EventMessage[events.ResumeCreated]) {
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
		log.Printf("Ивент сохранился с id: %s\n", event.Metadata.EventId)
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

func (es *EsEventStore) AppendEvents(m CommandMetadata, events ...interface{}) error {
	// if events == nil || len(events) == 0 {
	// 	return nil
	// }

	var serializedEvents []map[string]interface{}
	for _, i := range events {
		// fmt.Println(i)
		serializedEvent, err := es.eventSerde.Serialize(i, EventMetadata{}) // TODO: Доделать EventMetadata
		if err != nil {
			return err
		}

		serializedEvents = append(serializedEvents, serializedEvent)
	}

	for _, e := range serializedEvents {
		es.writeRepo.Save(e)
	}

	return nil
}
