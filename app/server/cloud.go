package server

import (
	"SebStudy/infrastructure"
	"context"
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudServer struct {
	Dispatcher *infrastructure.Dispatcher
	Server     cloudevents.Client
}

func NewCloudEventsClient(port int) cloudevents.Client {
	ce, err := cloudevents.NewClientHTTP(cloudevents.WithPort(port))
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	return ce
}

func NewCloudServer(d *infrastructure.Dispatcher, ce cloudevents.Client) *CloudServer {
	return &CloudServer{
		Dispatcher: d,
		Server:     ce,
	}
}

func (c *CloudServer) Run() {
	if err := c.Server.StartReceiver(context.Background(), c.receive); err != nil {
		log.Fatalf("failed to start receiver: %v", err)
	}
}

func (c *CloudServer) receive(event cloudevents.Event) {
	fmt.Printf("%s", event)
	//c.Dispatcher.Dispatch(event, infrastructure.NewCommandMetadataFromCloudEvent(event))
}
