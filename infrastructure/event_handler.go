package infrastructure

import (
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type ToType func()

type EventHandler struct {
	EventBus *EventBus
	handlers map[string]ToType
}

func NewEventHandler(eventBus *EventBus) *EventHandler {
	eh := &EventHandler{
		EventBus: eventBus,
		handlers: make(map[string]ToType, 0),
	}

	return eh
}

func (eh *EventHandler) Handle(event interface{}, metadata EventMetadata) error { // Пофиксить тему которая касается определения версии
	eventMes := NewEventMessage(event, metadata, 0)
	log.Println(event)
	eh.EventBus.Publish(metadata.EventType, eventMes)

	return cloudevents.ResultACK
}
