package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
	"SebStudy/ports"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(eventSender ports.CeEventSender, repository ResumeRepository) *CommandHandlers {
	commandHandlers := &CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.SendResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(*commands.SendResume)
		resume := NewResume()

		if err := resume.SendResume(cmd, eventSender); err != nil {
			return err
		}

		return nil
	})

	return commandHandlers
}
