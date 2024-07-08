package infrastructure

import (
	"fmt"
	"log"
)

type eventHandler func(interface{}, EventMetadata) error

type EventHandlerMap struct {
	handlers map[string]eventHandler
}

func NewEventHandlerMap() EventHandlerMap {
	eh := EventHandlerMap{
		handlers: make(map[string]eventHandler),
	}

	eh.register("resume.sended", resumeSended)

	return eh
}

func (eh *EventHandlerMap) register(eventType string, f eventHandler) {
	eh.handlers[eventType] = f
}

func (eh *EventHandlerMap) Get(eventType string) (eventHandler, error) {
	if h, ok := eh.handlers[eventType]; ok {
		return h, nil
	}

	return nil, fmt.Errorf("handler for event %s does not exist", eventType)
}

func resumeSended(e interface{}, m EventMetadata) error {
	log.Printf("Дошло до нужного хендлера событий с типом - {%s} и uuid - {%s}", m.EventType, m.EventId)
	return nil
}
