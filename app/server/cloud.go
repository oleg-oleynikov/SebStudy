package server

import (
	"SebStudy/infrastructure"
	"context"
	"fmt"
	"log"
	"reflect"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudServer struct {
	Dispatcher *infrastructure.Dispatcher
	Server     cloudevents.Client
	CeMapper   *infrastructure.CeMapper
}

func NewCloudEventsClient(port int) cloudevents.Client {
	ce, err := cloudevents.NewClientHTTP(cloudevents.WithPort(port))
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	return ce
}

func NewCloudServer(d *infrastructure.Dispatcher, c cloudevents.Client, ceMapper *infrastructure.CeMapper) *CloudServer {
	return &CloudServer{
		Dispatcher: d,
		Server:     c,
		CeMapper:   ceMapper,
	}
}

func (c *CloudServer) Run() {
	if err := c.Server.StartReceiver(context.Background(), c.receive); err != nil {
		log.Fatalf("failed to start receiver: %v", err)
	}
}

func (c *CloudServer) receive(event cloudevents.Event) {
	// fmt.Printf("%s", event)
	// fmt.Println(event.Type())

	cmd, err := c.CeMapper.MapToCommand(event)
	if err != nil {
		panic(fmt.Errorf("failed to map cloudevent: %v", err)) // убрать панику
	}
	fmt.Printf("Заглушка: %s\n", reflect.TypeOf(cmd))
	// proto.Unmarshal(event.Data, )
	// c.Dispatcher.Dispatch(event, infrastructure.NewCommandMetadataFromCloudEvent(event))
}
