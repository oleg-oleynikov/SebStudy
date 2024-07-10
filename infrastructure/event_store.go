package infrastructure

import (
	"SebStudy/ports/db_ports"
	"log"
)

type EventStore struct {
	eventBus   *EventBus
	eventSerde *EventSerde
	writeRepo  db_ports.WriteModel
}

// TODO: Тут нужна обработка полученных из шины событий и отправка их в event store через port

func NewEventStore(eventBus *EventBus, eventSerde *EventSerde, writeRepo db_ports.WriteModel) *EventStore {
	es := &EventStore{
		eventBus:   eventBus,
		eventSerde: eventSerde,
		writeRepo:  writeRepo,
	}

	es.eventBus.Subscribe("resume.*", func(eventMes *EventMessage) {
		data, err := es.eventSerde.Serialize(eventMes.Event, eventMes.Metadata)
		if err != nil {
			log.Printf("Event serde can not serialize event: %v\n", err)
			return
		}
		es.writeRepo.Save(data)
		log.Println("Ну типа сохранилось")
	})

	return es
}

func GetEvents(aggregateId int) []interface{} {
	return nil
}

// Отправка в write модель
