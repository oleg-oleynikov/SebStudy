package util

import (
	"SebStudy/infrastructure"
	"SebStudy/logger"
	"SebStudy/pb"
	"SebStudy/ports"
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommandHandler struct {
	cmdDispatcher ports.CommandDispatcher
	cmdAdapter    *CloudEventCommandAdapter
	log           logger.Logger
}

func NewCommandHandler(cmdDispatcher ports.CommandDispatcher, cmdAdapter *CloudEventCommandAdapter, log logger.Logger) *CommandHandler {
	return &CommandHandler{
		cmdDispatcher: cmdDispatcher,
		cmdAdapter:    cmdAdapter,
		log:           log,
	}
}

func (ch *CommandHandler) HandleCommand(ctx context.Context, event *pb.CloudEvent) error {
	mapper, err := ch.cmdAdapter.GetCloudeventToEvent(event.Type)
	if err != nil {
		ch.log.Debugf("unknown event type: %v", err)
		return status.Errorf(codes.InvalidArgument, "unknown event type")
	}

	mappedEvent, err := mapper(ctx, event)
	if err != nil {
		ch.log.Debugf("failed to map cloudevent: %v", err)
		return status.Errorf(codes.Internal, "failed to map cloudevent")
	}

	err = ch.cmdDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadata(event.Id))
	if err != nil {
		ch.log.Debugf("Failed to dispatch command: %v", err)
	}

	return err
}
