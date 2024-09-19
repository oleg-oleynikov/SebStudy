package adapters

// import (
// 	"SebStudy/internal/infrastructure/ports"
// 	"SebStudy/internal/infrastructure/util"
// 	"SebStudy/logger"
// 	"SebStudy/pb"
// 	"context"

// 	"github.com/golang/protobuf/ptypes/empty"
// 	"google.golang.org/protobuf/types/known/emptypb"
// )

// type CloudEventService struct {
// 	pb.UnimplementedCloudEventServiceServer
// 	log               logger.Logger
// 	eventProcessor    *EventProcessor
// 	commandHandler    *CommandHandler
// 	subscriberManager *SubscriberManager
// }

// func NewCloudEventService(
// 	log logger.Logger,
// 	cmdDispatcher ports.CommandDispatcher,
// 	cmdAdapter *util.CloudEventCommandAdapter,
// ) *CloudEventService {

// 	subscriberManager := NewSubcriberManager(log)

// 	eventProcessor := NewEventProcessor(100, log, subscriberManager)

// 	commandHandler := NewCommandHandler(cmdDispatcher, cmdAdapter, log)

// 	return &CloudEventService{
// 		log:               log,
// 		eventProcessor:    eventProcessor,
// 		commandHandler:    commandHandler,
// 		subscriberManager: subscriberManager,
// 	}
// }

// func (s *CloudEventService) Publish(ctx context.Context, req *pb.PublishRequest) (*emptypb.Empty, error) {
// 	event := req.GetEvent()

// 	err := s.eventProcessor.SubmitEvent(event)
// 	if err != nil {
// 		return nil, err
// 	}

// 	go s.eventProcessor.StartProcessing()

// 	err = s.commandHandler.HandleCommand(ctx, event)
// 	if err != nil {
// 		return &empty.Empty{}, err
// 	}

// 	return &empty.Empty{}, nil
// }

// func (s *CloudEventService) Subscribe(req *pb.SubscriptionRequest, stream pb.CloudEventService_SubscribeServer) error {
// 	s.log.Printf("New subscriber: %v", req)
// 	s.subscriberManager.AddSubscriber(stream)

// 	<-stream.Context().Done()

// 	err := stream.Context().Err()
// 	if err != nil {
// 		s.log.Debugf("Subscriber disconnected: %v\n", err)
// 	}

// 	s.subscriberManager.RemoveSubscriber(stream)
// 	return nil
// }
