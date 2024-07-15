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
	// imageStore *ImageStore
}

func NewEsEventStore(eventBus *EventBus, eventSerde *EventSerde, writeRepo db_ports.WriteModel, imageStore *ImageStore) *EsEventStore {
	es := &EsEventStore{
		eventBus:   eventBus,
		eventSerde: eventSerde,
		writeRepo:  writeRepo,
		// imageStore: imageStore,
	}

	es.eventBus.Subscribe("resume.sended", func(event *EventMessage[events.ResumeSended]) {
		// log.Println(reflect.TypeOf(eventMes.Event))
		// ev := eventMes.Event
		// log.Println(eventMes)

		// log.Println(GetType(ev))
		// ev1, ok := (eventMes.Event).(events.ResumeSended)
		// log.Println(ev1)
		// log.Println(ok)

		// resumeSended, ok := eventMes.Event.(events.ResumeSended)
		// if !ok {
		// 	log.Println("Что то не так")
		// 	return
		// }

		// imagePath, err := es.imageStore.SaveImage(resumeSended.Photo.GetPhoto())
		// if err != nil {
		// 	log.Printf("Image doesnt save: %v", err)
		// 	return
		// }

		// resumeSended.Photo.SetUrl(imagePath)
		// data, err := es.eventSerde.Serialize(eventMes)
		// if err != nil {
		// 	log.Printf("Event serde can not serialize event{%s}: %v\n", eventMes.Metadata.EventId, err)
		// 	return
		// }

		// if err := es.writeRepo.Save(data); err != nil {
		// 	log.Println()
		// }
		log.Println(event)
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
