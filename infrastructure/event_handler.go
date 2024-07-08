package infrastructure

import (
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventHandler struct {
	EventBus *EventBus
}

func NewEventHandler(eventBus *EventBus) *EventHandler {
	eh := &EventHandler{
		EventBus: eventBus,
	}
	// eh.EventBus.Subscribe("hello", func(mes *nats.Msg) {
	// 	fmt.Printf("received message: %v\n", mes)
	// })
	// return &EventHandler {
	// 	EventBus: eventBus,
	// }
	return eh
}

func (eh *EventHandler) Handle(event interface{}, metadata EventMetadata) error {
	log.Printf("Дошло до хендлера событий с типом - {%s} и uuid - {%s}", metadata.EventType, metadata.EventId)
	return cloudevents.ResultACK

}
