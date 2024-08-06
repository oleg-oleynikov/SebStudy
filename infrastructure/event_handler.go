package infrastructure

import (
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	log.Println("Event Handler")
	log.Println(event)
	if err := eh.EventBus.Publish(metadata.EventType, eventMes); err != nil {
		return err
	}

	return status.Error(codes.OK, "OK")
}
