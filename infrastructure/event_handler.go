package infrastructure

import (
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type ToType func()

type EventHandler struct {
	EventBus   *EventBus
	EventStore EventStore
	handlers   map[string]ToType
}

func NewEventHandler(eventBus *EventBus, eventStore EventStore) *EventHandler {
	eh := &EventHandler{
		EventBus:   eventBus,
		EventStore: eventStore,
		handlers:   make(map[string]ToType, 0),
	}

	return eh
}

func (eh *EventHandler) Handle(event interface{}, metadata EventMetadata) error {
	// TODO: Сделать запрос в event store для сбора событий по агрегату и его восстановление
	// tEvent := GetType(event)
	// eventMessage := NewEventMessage(event, metadata, 0) // Публикация но тут я б еще подумал

	// eh.EventBus.Publish(metadata.EventType, eventMessage)
	// log.Println("До публикации: ", event)
	// err := eh.EventBus.Publish(metadata.EventType, eventMessage)
	// if err != nil {
	// 	return err
	// }

	log.Println("Пока блять на переделку нахуй")

	return cloudevents.ResultACK
}
