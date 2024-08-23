package primary

import (
	"SebStudy/adapters/util"
	"SebStudy/infrastructure"
	"SebStudy/infrastructure/logger"
	"SebStudy/ports"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type CloudEventsAdapter struct {
	CommandDispatcher ports.CeCommandDispatcher
	EventDispatcher   ports.CeEventDispatcher
	CloudeventMapper  *util.CloudeventMapper
}

func NewCloudEventsAdapter(d ports.CeCommandDispatcher, e ports.CeEventDispatcher, cloudeventMapper *util.CloudeventMapper) *CloudEventsAdapter {
	return &CloudEventsAdapter{
		CommandDispatcher: d,
		EventDispatcher:   e,

		CloudeventMapper: cloudeventMapper,
	}
}

func (c *CloudEventsAdapter) ReceiveCloudEvent(ctx context.Context, event *v1.CloudEvent) error {
	// if _, err := c.CloudeventMapper.(event.Type); err != nil {
	// 	log.Printf("unknown event type: %s\n", err)
	// 	return status.Errorf(codes.InvalidArgument, "unknown event type: %s", err)
	// }

	mapper, err := c.CloudeventMapper.GetCloudeventToEvent(event.Type)

	if err != nil {
		logger.Logger.Printf("unknown event type: %v", err)
		return status.Errorf(codes.InvalidArgument, "unknown event type: %s", err)
	}

	mappedEvent, err := mapper(ctx, event)
	if err != nil {
		return status.Errorf(codes.Internal, "failed to map cloudevent: %v", err)
	}

	// if c.CloudeventMapper.IsCommand(event.Type) {
	err = c.CommandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadataFromCloudEvent(event))
	if _, ok := status.FromError(err); ok {
		return err
	}

	logger.Logger.Debugf("failed to dispatch command: %v", err)

	return status.Errorf(codes.Internal, "failed to dispatch command: %v", err)

	// } else if c.CloudeventMapper.IsEvent(event.Type) {
	// 	err := c.EventDispatcher.Dispatch(mappedEvent, *infrastructure.NewEventMetadataFromCloudEvent(event))
	// 	if _, ok := status.FromError(err); ok {
	// 		return err
	// 	}

	// 	logger.Logger.Printf("failed to dispatch event: %v", err)
	// 	return status.Errorf(codes.InvalidArgument, "failed to dispatch event: %v", err)
	// }

	// return status.Errorf(codes.Canceled, "Something wrong")
}
