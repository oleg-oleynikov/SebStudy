package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
	"SebStudy/ports"

	"github.com/google/uuid"
)

func NewHandlers(eventSender ports.CeEventSender, repository ResumeRepository) *infrastructure.CommandHandlerBase {
	commandHandlers := infrastructure.NewCommandHandler()

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

		repository.Save(resume, m)
		return nil
	})

	return commandHandlers
}

type ResumeCommandHandlers struct {
	cmdHandlers *infrastructure.CommandHandlerBase
}

func (m *ResumeCommandHandlers) RegisterCommands(cmdHandlerMap *infrastructure.CommandHandlerMap) {
	cmdHandlerMap.AppendHandlers(m.cmdHandlers)
}

func NewResumeCommandHandlers(eventSender ports.CeEventSender, repository ResumeRepository) *ResumeCommandHandlers {
	commandHandlers := infrastructure.NewCommandHandler()

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

		return repository.Save(resume, m)
	})

	return &ResumeCommandHandlers{
		cmdHandlers: commandHandlers,
	}
}
