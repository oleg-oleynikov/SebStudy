package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
	"SebStudy/ports"

	"github.com/google/uuid"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(eventSender ports.CeEventSender, repository ResumeRepository) *CommandHandlers {
	commandHandlers := &CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.CreateResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(*commands.CreateResume)
		resume := NewResume()

		aggregateId, err := uuid.NewV7()
		if err != nil {
			return err
		}
		cmd.ResumeId.Value = aggregateId.String()

		if err := resume.CreateResume(cmd, eventSender); err != nil {
			return err
		}

		return nil
	})

	return commandHandlers
}
