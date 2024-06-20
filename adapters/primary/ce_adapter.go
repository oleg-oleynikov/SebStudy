package primary

import (
	"SebStudy/infrastructure"
	"SebStudy/ports"
	"context"
	"fmt"
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

func (c *CloudEventsAdapter) receive(event cloudevents.Event) {
	// fmt.Printf("%s", event)
	// fmt.Println(event.Type())

	cmd, err := c.CeMapper.MapToCommand(event)
	if err != nil {
		panic(fmt.Errorf("failed to map cloudevent: %v", err)) // убрать панику
	}

	// fmt.Println("Тип команды: ", reflect.TypeOf(cmd))
	c.Dispatcher.Dispatch(cmd, infrastructure.NewCommandMetadataFromCloudEvent(event))
}
