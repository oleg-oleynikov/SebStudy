package infrastructure

import (
	"SebStudy/domain/resume/events"
	"SebStudy/ports/db_ports"
	"log"
)

type EventStore struct {
	eventBus   *EventBus
	eventSerde *EventSerde
	writeRepo  db_ports.WriteModel
	imageStore *ImageStore
}

func NewEventStore(eventBus *EventBus, eventSerde *EventSerde, writeRepo db_ports.WriteModel, imageStore *ImageStore) *EventStore {
	es := &EventStore{
		eventBus:   eventBus,
		eventSerde: eventSerde,
		writeRepo:  writeRepo,
		imageStore: imageStore,
	}

	es.eventBus.Subscribe("resume.sended", func(event events.ResumeSended) {
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

		log.Printf("Ну типа сохранилось вот ивент мес: %v\n", event)
	})

	return es
}

// Получение всех ивентов по id агрегата

func GetEvents(aggregateId int) []interface{} {
	return nil
}
