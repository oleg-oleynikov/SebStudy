package primary

import (
	"SebStudy/adapters/util"
	"SebStudy/infrastructure"
	"SebStudy/ports"
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudEventsAdapter struct {
	CommandDispatcher ports.CeCommandDispatcher
	EventDispatcher   ports.CeEventHandler
	Client            cloudevents.Client
	CeMapper          *util.CeMapper
}

func NewCloudEventsAdapter(d ports.CeCommandDispatcher, e ports.CeEventHandler, ceMapper *util.CeMapper) *CloudEventsAdapter {
	return &CloudEventsAdapter{
		CommandDispatcher: d,
		EventDispatcher:   e,

		Client:   newCloudEventsClient(),
		CeMapper: ceMapper,
	}
}

func newCloudEventsClient() cloudevents.Client {
	p, err := cloudevents.NewHTTP()
	if err != nil {
		log.Fatalf("failed to create protocol: %s", err.Error())
	}

	ce, err := cloudevents.NewClient(p)
	if err != nil {
		log.Fatalf("failed to create http client, %v", err)
	}

	return ce
}

func (c *CloudEventsAdapter) Run() {
	go func() {
		log.Fatalf("failed to start receiver: %s", c.Client.StartReceiver(context.Background(), c.receive))
	}()
}

func (c *CloudEventsAdapter) receive(ctx context.Context, event cloudevents.Event) cloudevents.Result {
	// event.Data()
	log.Println("Пришло что то нахуй")

	if _, err := c.CeMapper.GetEventType(event.Type()); err != nil {
		return cloudevents.NewHTTPResult(400, "Unknown event type: %s", err)
	}

	mappedEvent, err := c.CeMapper.MapToEvent(ctx, event)

	if err != nil {
		log.Printf("failed to map cloudevent: %v", err)
		return cloudevents.NewHTTPResult(400, "failed to map cloudevent: %v", err)
	}

	if c.CeMapper.IsCommand(event.Type()) {
		err = c.CommandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadataFromCloudEvent(event))
		if err != nil {
			if _, ok := err.(cloudevents.Result); ok {
				return err
			}

			log.Printf("failed to dispatch command: %v", err)
			return cloudevents.NewHTTPResult(500, "failed to dispatch command: %v", err)
		}
	} else if c.CeMapper.IsEvent(event.Type()) {
		err := c.EventDispatcher.Handle(mappedEvent, *infrastructure.NewEventMetadataFromCloudEvent(event))
		if err != nil {
			if _, ok := err.(cloudevents.Result); ok {
				return err
			}

			log.Printf("failed to handle event: %v", err)
			return cloudevents.NewHTTPResult(500, "failed to handle event: %v", err)
		}
	}

	return cloudevents.ResultNACK
}
