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
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func main() {
	client := NewCloudeventsServiceClient("localhost:50051")

	resumeEvent := &pb.ResumeCreated{
		ResumeId:      "12121",
		FirstName:     "олег",
		MiddleName:    "олейников",
		LastName:      "игоревич",
		PhoneNumber:   "79885352919",
		Education:     "ш",
		AboutMe:       "м",
		AboutProjects: "а",
		Portfolio:     "л",
		StudentGroup:  "ь",
		CreatedAt:     timestamppb.Now(),
	}

	protoEvent, err := anypb.New(resumeEvent)
	if err != nil {
		log.Println("Failed to cast to anypb")
		return
	}

	cloudevent := &pb.CloudEvent{
		Id:          "1231312345",
		Source:      "example.com",
		SpecVersion: "1.0",
		Type:        "resume.create",
		Data: &pb.CloudEvent_ProtoData{
			ProtoData: protoEvent,
		},
	}

	client.Publish(cloudevent)
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
	log.Println("Фрико фркик")

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
	// counterEvents++
	// log.Println("Отловлено: ", counterEvents)
}
