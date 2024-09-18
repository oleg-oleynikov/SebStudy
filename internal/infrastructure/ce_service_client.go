package infrastructure

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

func (c *CloudeventsServiceClient) Subscribe(ctx context.Context, source string, fn func(event *pb.CloudEvent)) {
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
				fn(event)
				// c.processEvent(event)
			}()
		}
	}
}

// func (c *CloudeventsServiceClient) processEvent(event *v1.CloudEvent) {
// 	c.Lock()
// 	defer c.Unlock()
// 	log.Printf("Received event: %+v", event)
// }
