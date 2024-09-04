package primary

import (
	"SebStudy/adapters/util"
	"SebStudy/infrastructure"
	"SebStudy/logger"
	"SebStudy/pb"
	"SebStudy/ports"
	"context"
	"fmt"

	"sync"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	// pb "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type CloudEventService struct {
	sync.RWMutex
	pb.UnimplementedCloudEventServiceServer

	log logger.Logger
	// server      *grpc.Server
	subscribers       map[pb.CloudEventService_SubscribeServer]chan *pb.CloudEvent
	eventChan         chan *pb.CloudEvent
	wg                sync.WaitGroup
	commandDispatcher ports.CeCommandDispatcher
	cloudeventMapper  *util.CloudeventMapper
}

func NewCloudEventService(log logger.Logger, cmdDispatcher ports.CeCommandDispatcher, ceMapper *util.CloudeventMapper /* opt ...grpc.ServerOption*/) *CloudEventService {
	return &CloudEventService{
		subscribers:       make(map[pb.CloudEventService_SubscribeServer]chan *pb.CloudEvent),
		eventChan:         make(chan *pb.CloudEvent, 100),
		commandDispatcher: cmdDispatcher,
		cloudeventMapper:  ceMapper,
	}
}

func (s *CloudEventService) Publish(ctx context.Context, req *pb.PublishRequest) (*emptypb.Empty, error) {
	event := req.GetEvent()
	select {
	case s.eventChan <- event:
	default:
		s.log.Printf("eventChan is full, dropping event: %v", event)
		return nil, fmt.Errorf("eventChan is full, event dropped")
	}

	go s.processEvents()

	// err := s.ceReceiver.ReceiveCloudEvent(ctx, event)
	mapper, err := s.cloudeventMapper.GetCloudeventToEvent(event.Type)
	if err != nil {
		s.log.Debugf("unknown event type: %v", err)
		return &empty.Empty{}, status.Errorf(codes.InvalidArgument, "unknown event type")
	}

	mappedEvent, err := mapper(ctx, event)
	if err != nil {
		s.log.Debugf("failed to map cloudevent: %v", err)

		// return status.Errorf(codes.Internal, "failed to map cloudevent: %v", err)
		return &empty.Empty{}, status.Errorf(codes.Internal, "failed to map cloudevent")
	}

	err = s.commandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadata(event.Id))
	if _, ok := status.FromError(err); ok {
		return &empty.Empty{}, err
	}

	s.log.Debugf("Response with error: %v", err)

	return &empty.Empty{}, err
}

func (s *CloudEventService) processEvents() {
	for {
		select {
		case event := <-s.eventChan:
			s.RLock()
			s.wg.Add(1)
			go func(event *pb.CloudEvent) {
				defer s.wg.Done()
				s.processEvent(event)
			}(event)
			s.RUnlock()
		default:
			return
		}
	}
}

func (s *CloudEventService) processEvent(event *pb.CloudEvent) {
	for _, eventChan := range s.subscribers {
		select {
		case eventChan <- event:
		default:
			s.log.Printf("Subscriber queue is full, dropping event")
		}
	}
}

func (s *CloudEventService) addSubscriber(stream pb.CloudEventService_SubscribeServer) {
	s.Lock()
	defer s.Unlock()

	eventChan := make(chan *pb.CloudEvent, 100)
	s.subscribers[stream] = eventChan

	go func() {
		for {
			event, ok := <-eventChan
			if !ok {
				return
			}
			if err := stream.Send(event); err != nil {
				s.log.Printf("Failed to send event to subscriber: %v", err)
				s.removeSubscriber(stream)
				return
			}
		}
	}()
}

func (s *CloudEventService) removeSubscriber(subscriber pb.CloudEventService_SubscribeServer) {
	s.Lock()
	defer s.Unlock()

	if eventChan, ok := s.subscribers[subscriber]; ok {
		subscriber.Context().Done()
		close(eventChan)
		delete(s.subscribers, subscriber)
	}
}

func (s *CloudEventService) Subscribe(req *pb.SubscriptionRequest, stream pb.CloudEventService_SubscribeServer) error {
	s.log.Printf("New subscriber: %v", req)
	s.addSubscriber(stream)

	<-stream.Context().Done()

	err := stream.Context().Err()
	if err != nil {
		s.log.Printf("Subscriber disconnected: %v\n", err)
	}

	s.removeSubscriber(stream)
	return nil
}
