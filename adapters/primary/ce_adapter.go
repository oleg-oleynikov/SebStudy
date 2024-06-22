package primary

import (
	"SebStudy/infrastructure"
	"SebStudy/ports"
	"context"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudEventsAdapter struct {
	Dispatcher ports.CeEventDispatcher
	Server     cloudevents.Client
	CeMapper   *CeMapper
}

func NewCloudEventsAdapter(d ports.CeEventDispatcher, ceMapper *CeMapper, port int) *CloudEventsAdapter {
	return &CloudEventsAdapter{
		Dispatcher: d,
		Server:     newCloudEventsClient(port),
		CeMapper:   ceMapper,
	}
}

func newCloudEventsClient(port int) cloudevents.Client {
	ce, err := cloudevents.NewClientHTTP(cloudevents.WithPort(port))
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	return ce
}

func (c *CloudEventsAdapter) Run() {
	if err := c.Server.StartReceiver(context.Background(), c.receive); err != nil {
		log.Fatalf("failed to start receiver: %v", err)
	}
}

func (c *CloudEventsAdapter) receive(ctx context.Context, event cloudevents.Event) {

	cmd, err := c.CeMapper.MapToCommand(ctx, event)
	if err != nil {
		log.Printf("failed to map cloudevent: %v", err)
		return
	}

	err = c.Dispatcher.Dispatch(cmd, infrastructure.NewCommandMetadataFromCloudEvent(event))
	if err != nil {
		log.Printf("failed to dispatch command: %v", err)
		return
	}
}
