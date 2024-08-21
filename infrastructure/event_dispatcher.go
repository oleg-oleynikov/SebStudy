package infrastructure

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EventDispatcher struct {
	EventBus EventBus
}

func NewEventDispatcher(eventBus EventBus) *EventDispatcher {
	eh := &EventDispatcher{
		EventBus: eventBus,
	}

	return eh
}

func (eh *EventDispatcher) Dispatch(event interface{}, metadata EventMetadata) error {
	eventMes := NewEventEnvelope(event, metadata, 0)
	if err := eh.EventBus.Publish(metadata.EventType, eventMes); err != nil {
		return err
	}

	return status.Error(codes.OK, "OK")
}
