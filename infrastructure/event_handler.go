package infrastructure

import (
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventHandler struct {
	EventBus   *EventBus
	EventStore *EventStore
}

func NewEventHandler(eventBus *EventBus, eventStore *EventStore) *EventHandler {
	eh := &EventHandler{
		EventBus:   eventBus,
		EventStore: eventStore,
	}

	return eh
}

func (eh *EventHandler) Handle(event interface{}, metadata EventMetadata) error {
	// TODO: Сделать запрос в event store для сбора событий по агрегату и его восстановление

	// eventMessage := NewEventMessage(event, metadata, 0) // Публикация но тут я б еще подумал

	// eh.EventBus.Publish(metadata.EventType, eventMessage)
	err := eh.EventBus.Publish(metadata.EventType, event)
	if err != nil {
		return err
	}

	return cloudevents.ResultACK
}
