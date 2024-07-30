package primary

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	proto "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type CloudEventServiceServer struct {
	sync.RWMutex
	proto.UnimplementedCloudEventServiceServer

	server      *grpc.Server
	subscribers map[proto.CloudEventService_SubscribeServer]chan *proto.CloudEvent
	eventChan   chan *proto.CloudEvent
}

func NewCloudEventServiceServer(opt ...grpc.ServerOption) *CloudEventServiceServer {
	server := grpc.NewServer(opt...)
	return &CloudEventServiceServer{
		server:      server,
		subscribers: make(map[proto.CloudEventService_SubscribeServer]chan *proto.CloudEvent),
		eventChan:   make(chan *proto.CloudEvent, 100),
	}
}

func (s *CloudEventServiceServer) Publish(ctx context.Context, req *proto.PublishRequest) (*emptypb.Empty, error) {
	event := req.GetEvent()
	select {
	case s.eventChan <- event:
	default:
		log.Printf("eventChan is full, dropping event: %v", event)
		return nil, fmt.Errorf("eventChan is full, event dropped")
	}

	go s.processEvents()
	return &empty.Empty{}, nil
}

func (s *CloudEventServiceServer) processEvents() {
	for event := range s.eventChan {
		s.RLock()
		go func() {
			for _, eventChan := range s.subscribers {
				select {
				case eventChan <- event:
				default:
					log.Printf("Subscriber queue is full, dropping event")

					// log.Printf("Subscriber queue is full, trying to resend later")
					// go func(sub proto.CloudEventService_SubscribeServer, event *proto.CloudEvent) {
					// 	select {
					// 	case eventChan <- event:
					// 	case <-sub.Context().Done():
					// 		s.removeSubscriber(sub)
					// 	}
					// }(sub, event)
				}
			}
		}()
		s.RUnlock()
	}
}

func (s *CloudEventServiceServer) addSubscriber(stream proto.CloudEventService_SubscribeServer) {
	s.Lock()
	defer s.Unlock()

	eventChan := make(chan *proto.CloudEvent, 100)
	s.subscribers[stream] = eventChan

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}
			if err := stream.Send(event); err != nil {
				log.Printf("Failed to send event to subscriber: %v", err)
				s.removeSubscriber(stream)
				return
			}
		}
	}()
}

func (s *CloudEventServiceServer) removeSubscriber(subscriber proto.CloudEventService_SubscribeServer) {
	s.Lock()
	defer s.Unlock()

	if eventChan, ok := s.subscribers[subscriber]; ok {
		subscriber.Context().Done()
		close(eventChan)
		delete(s.subscribers, subscriber)
	}
}

func (s *CloudEventServiceServer) Subscribe(req *proto.SubscriptionRequest, stream proto.CloudEventService_SubscribeServer) error {
	log.Printf("New subscriber: %v", req)
	s.addSubscriber(stream)

	<-stream.Context().Done()

	err := stream.Context().Err()
	if err != nil {
		log.Printf("Subscriber disconnected: %v\n", err)
	}

	s.removeSubscriber(stream)
	return nil
}

func (s *CloudEventServiceServer) Run(network string, addr string) {
	lis, err := net.Listen(network, addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	proto.RegisterCloudEventServiceServer(s.server, s)

	log.Printf("Try starting listening server on %s\n", addr)
	go func() {
		if err := s.server.Serve(lis); err != nil {
			fmt.Printf("Failed to serve: %v", err)
		}
	}()
}

func (s *CloudEventServiceServer) Shutdown() {
	s.server.GracefulStop()

	s.RLock()
	for sub := range s.subscribers {
		s.removeSubscriber(sub)
	}
	s.RUnlock()
}

// func main() {
// 	quit := make(chan os.Signal, 1)
// 	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

// 	// Вот здесь в opt если что можно по идее будет подключить middleware
// 	c := NewCloudEventServiceServer()

// 	c.Run("tcp", ":50051")

// 	<-quit
// 	c.Shutdown()
// }
