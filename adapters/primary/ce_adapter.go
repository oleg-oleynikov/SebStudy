package primary

import (
	"SebStudy/adapters/util"
	"SebStudy/infrastructure"
	"SebStudy/ports"
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

type CloudEventsAdapter struct {
	CommandDispatcher ports.CeCommandDispatcher
	EventDispatcher   ports.CeEventHandler
	CeMapper          *util.CeMapper
}

func NewCloudEventsAdapter(d ports.CeCommandDispatcher, e ports.CeEventHandler, ceMapper *util.CeMapper) *CloudEventsAdapter {
	return &CloudEventsAdapter{
		CommandDispatcher: d,
		EventDispatcher:   e,

		CeMapper: ceMapper,
	}
}

func (c *CloudEventsAdapter) ReceiveCloudEvent(event *v1.CloudEvent) error {
	if _, err := c.CeMapper.GetEventType(event.Type); err != nil {
		log.Printf("unknown event type: %s\n", err)
		return status.Errorf(codes.InvalidArgument, "unknown event type: %s", err)
	}

	mappedEvent, err := c.CeMapper.MapToEvent(context.Background(), event)

	if err != nil {
		log.Printf("failed to map cloudevent: %v", err)
		return status.Errorf(codes.InvalidArgument, "failed to map cloudevent: %s", err)
	}

	if c.CeMapper.IsCommand(event.Type) {
		err = c.CommandDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadataFromCloudEvent(event))
		// if err != nil {
		// log.Println("Бля")
		if _, ok := status.FromError(err); ok {
			return err
		}

		log.Printf("failed to dispatch command: %v", err)
		return status.Errorf(codes.Internal, "failed to dispatch command: %v", err)
		// }

	} else if c.CeMapper.IsEvent(event.Type) {
		err := c.EventDispatcher.Handle(mappedEvent, *infrastructure.NewEventMetadataFromCloudEvent(event))
		if _, ok := status.FromError(err); ok {
			return err
		}

		log.Printf("failed to dispatch event: %v", err)
		return status.Errorf(codes.InvalidArgument, "failed to dispatch event: %v", err)
	}

	return status.Errorf(codes.Canceled, "Fuck you slave")
}
