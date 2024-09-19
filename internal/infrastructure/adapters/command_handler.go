package adapters

import (
	"SebStudy/internal/infrastructure/ports"
	"SebStudy/logger"
	"SebStudy/pb"
	"context"
)

type CommandHandler struct {
	cmdDispatcher ports.CommandDispatcher
	// cmdAdapter    *util.CloudEventCommandAdapter
	log logger.Logger
}

func NewCommandHandler(cmdDispatcher ports.CommandDispatcher /*cmdAdapter *util.CloudEventCommandAdapter, */, log logger.Logger) *CommandHandler {
	return &CommandHandler{
		cmdDispatcher: cmdDispatcher,
		// cmdAdapter:    cmdAdapter,
		log: log,
	}
}

func (ch *CommandHandler) HandleCommand(ctx context.Context, event *pb.CloudEvent) error {
	// mapper, err := ch.cmdAdapter.GetCloudeventToEvent(event.GetProtoData().GetTypeUrl())
	// if err != nil {
	// 	ch.log.Debugf("unknown event type: %v", err)
	// 	return status.Errorf(codes.InvalidArgument, "unknown event type")
	// }

	// mappedEvent, err := mapper(ctx, event)
	// if err != nil {
	// 	ch.log.Debugf("failed to map cloudevent: %v", err)
	// 	return status.Errorf(codes.Internal, "failed to map cloudevent")
	// }

	// err = ch.cmdDispatcher.Dispatch(mappedEvent, infrastructure.NewCommandMetadata(event.Id)) // CommandMetadata useless хз че с этим делать
	// if err != nil {
	// 	ch.log.Debugf("Failed to dispatch command: %v", err)
	// }
	//
	// return err
	return nil
}
