package infrastructure

import (
	"SebStudy/domain/resume/events"

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

func (eh *EventHandler) Handle(event interface{}, metadata EventMetadata) error {
	if metadata.EventType == "resume.sended" { // Сделать нормально надо это))
		eventMes := NewEventMessage(event.(events.ResumeSended), metadata, 0)
		eh.EventBus.Publish(metadata.EventType, eventMes)
	}

	return cloudevents.ResultACK
}
