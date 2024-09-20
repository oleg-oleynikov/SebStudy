package main

import (
	"SebStudy/pb"
	"context"
	"fmt"
	"io"
	"log"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// client := NewCloudeventsServiceClient("localhost:50051")

	// resumeEvent := &pb.ResumeCreated{
	// 	Education: "PTY",
	// 	AboutMe:   "м",
	// 	Skills: []*pb.Skill{
	// 		{Skill: "fffff"},
	// 	},
	// 	BirthDate:     time.Date(2005, 1, 1, 0, 0, 0, 0, time.Local).Format("2006-01-02"),
	// 	AboutProjects: "а",
	// 	Portfolio:     "л",
	// 	Direction:     "ь",
	// }

	// protoEvent, err := anypb.New(resumeEvent)
	// if err != nil {
	// 	log.Println("Failed to cast to anypb")
	// 	return
	// }
	// log.Println(base64.StdEncoding.EncodeToString(protoEvent.Value))

	// cloudevent := &pb.CloudEvent{
	// 	Id:          "1231312345",
	// 	Source:      "example.com",
	// 	SpecVersion: "1.0",
	// 	Type:        "resume.create",
	// 	Data: &pb.CloudEvent_ProtoData{
	// 		ProtoData: protoEvent,
	// 	},
	// }

	// log.Println(base64.StdEncoding.EncodeToString(protoEvent.Value))

	// client.Publish(cloudevent)

	// cfg := config.InitConfig()
	// log := logger.NewAppLogger(cfg.Logger)
	// log.InitLogger()

	// nc, _ := nats.Connect(nats.DefaultURL)
	// typeMapper := eventsourcing.NewTypeMapper()
	// resume.RegisterResumeMappingTypes(typeMapper)
	// serde := eventsourcing.NewEsEventSerde(log, typeMapper)

	// eventStore := eventsourcing.NewJetStreamEventStore(log, nc, serde, "sebstudy")
	// aggregateStore := eventsourcing.NewEsAggregateStore(log, eventStore)

	// resume := models.NewResume()
	// aggregateStore.Load("0191e67595a17633960283162bffe3c6", resume)

	// js, _ := jetstream.New(nc)
	// stream, err := js.Stream(context.Background(), "0191e67647dd7ccc9cd8f21c423a9615")
	// log.Println(stream)
	// // resume.
	// log.Debugf("БЛЯТТЬ: %v", resume)
}

var (
	counterEvents = 0
)

type CloudeventsServiceClient struct {
	sync.RWMutex
	grpcClient pb.CloudEventServiceClient
	wg         sync.WaitGroup
}

func NewCloudeventsServiceClient(target string) *CloudeventsServiceClient {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("Failed to connect: %v", err)
		return nil
	}

	grpcClient := pb.NewCloudEventServiceClient(conn)

	c := &CloudeventsServiceClient{
		grpcClient: grpcClient,
	}
	return c
}

func (c *CloudeventsServiceClient) Publish(cloudEvent *pb.CloudEvent) error {
	_, err := c.grpcClient.Publish(context.TODO(), &pb.PublishRequest{
		Event: cloudEvent,
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *CloudeventsServiceClient) Subscribe(ctx context.Context, source string) {
	subReq := &pb.SubscriptionRequest{
		Source: source,
	}

	stream, err := c.grpcClient.Subscribe(ctx, subReq)
	if err != nil {
		fmt.Printf("Failed to subscribe: %v", err)
		return
	}

	for {
		event, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Println("Stream closed")
				return
			} else {
				log.Fatalf("Error receiving event: %v", err)
			}
		} else {
			c.wg.Add(1)
			go func() {
				defer c.wg.Done()
				c.processEvent(event)
			}()
		}
	}
}

func (c *CloudeventsServiceClient) processEvent(event *pb.CloudEvent) {
	c.Lock()
	defer c.Unlock()
	log.Printf("Received event: %+v", event)
}
