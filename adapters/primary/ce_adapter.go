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
	CommandDispatcher ports.CeCommandDispatcher
	Server            cloudevents.Client
	CeMapper          *CeMapper
}

func NewCloudEventsAdapter(d ports.CeCommandDispatcher, ceMapper *CeMapper, port int) *CloudEventsAdapter {
	return &CloudEventsAdapter{
		CommandDispatcher: d,
		Server:            newCloudEventsClient(port),
		CeMapper:          ceMapper,
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
	go func() {
		log.Fatalf("failed to start receiver: %s", c.Server.StartReceiver(context.Background(), c.receive))
	}()
}

func (c *CloudEventsAdapter) receive(ctx context.Context, event cloudevents.Event) cloudevents.Result {

	// _, err := cutOffPostfix(event)
	// if err != nil {
	// 	return cloudevents.NewHTTPResult(500, "ce-type must have postfix from: %s", "{event, command}")
	// }

	ceEventType, err := c.CeMapper.GetEventType(event.Type())
	if err != nil {
		return cloudevents.NewHTTPResult(400, "%s", err)
	}

	mappedEvent, err := c.CeMapper.MapFromCloudEvent(ctx, event)

	if err != nil {
		log.Printf("failed to map cloudevent: %v", err)
		return cloudevents.NewHTTPResult(400, "failed to map cloudevent: %v", err)
	}

	if ceEventType == COMMAND {
		err = c.CommandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadataFromCloudEvent(event))
		if err != nil {
			log.Printf("failed to dispatch command: %v", err)
			return cloudevents.NewHTTPResult(500, "failed to dispatch command: %v", err)
		}

		return cloudevents.ResultACK
	} else if ceEventType == EVENT {
		fmt.Println("Отправка в Event Bus (NATS), а там еще чет дальше event store еще что то")
	}

	return cloudevents.ResultACK
}

// func cutOffPostfix(c cloudevents.Event) (string, error) {
// 	typesSplit := strings.Split(c.Type(), ".")
// 	lenTypes := len(typesSplit)
// 	postfix := typesSplit[lenTypes-1]
// 	c.SetType(strings.Join(typesSplit[:lenTypes-1], "."))
// 	return postfix, nil
// }
