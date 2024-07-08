package infrastructure

import (
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventHandler struct {
	EventBus *EventBus

	eventHandlerMap EventHandlerMap
}

func NewEventHandler(eventBus *EventBus, hm EventHandlerMap) *EventHandler {
	eh := &EventHandler{
		EventBus:        eventBus,
		eventHandlerMap: hm,
	}

	return eh
}

func (eh *EventHandler) Handle(event interface{}, metadata EventMetadata) error {
	log.Printf("Дошло до хендлера событий с типом - {%s} и uuid - {%s}", metadata.EventType, metadata.EventId)
	_, err := eh.eventHandlerMap.Get(metadata.EventType)
	if err != nil {
		return err
	}

	// err := handler(event, metadata)

	return cloudevents.ResultACK
}
