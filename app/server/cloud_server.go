package main

import (
	"context"
	"fmt"
	"log"

	pb "SebStudy/proto/resume"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

func main() {
	c, err := cloudevents.NewClientHTTP(cloudevents.WithPort(8080))
	if err != nil {
		log.Fatalf("failed to create client, %v", err)
	}
	if err = c.StartReceiver(context.Background(), receive); err != nil {
		log.Fatalf("failed to start receiver: %v", err)
	}
}

func receive(event cloudevents.Event) {
	fmt.Printf("%s", event)
	// fmt.Println(event.ID())
	fmt.Printf("%s\n", event.Data())
	res := pb.Resume{}
	proto.Unmarshal(event.Data(), &res)
	fmt.Println("---------------------------------------")
	fmt.Println(res.Photo)
	fmt.Println("---------------------------------------")
	fmt.Println(event.Time())
	fmt.Println("---------------------------------------")
	fmt.Println(event.Source())
	fmt.Println("---------------------------------------")
	fmt.Println(event.Type())
	fmt.Println("---------------------------------------")
	fmt.Println(event.DataContentType())
	fmt.Println("---------------------------------------")
	fmt.Println(event.ID())
	// fmt.Print()
	// fmt.Printf("type: %v", reflect.TypeOf(event))
}

// func SendCloudEvent(ctx context.Context, event cloudevents.Event) error {
// 	c, err := cloudevents.NewClientHTTP()
// 	if err != nil {
// 		log.Fatalf("failed to create client, %v", err)
// 	}
// 	return c.Send(ctx, event)
// }
