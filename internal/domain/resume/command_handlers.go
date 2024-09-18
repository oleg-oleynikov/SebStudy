package resume

import (
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure"
	"fmt"
)

type ResumeCommandHandlers struct {
	cmdHandlers *infrastructure.CommandHandlerBase
}

func (m *ResumeCommandHandlers) RegisterCommands(cmdHandlerMap *infrastructure.CommandHandlerMap) {
	cmdHandlerMap.AppendHandlers(m.cmdHandlers)
}

func NewResumeCommandHandlers(repository ResumeRepository) *ResumeCommandHandlers {
	commandHandlers := infrastructure.NewCommandHandler()

	commandHandlers.Register(commands.CreateResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(*commands.CreateResume)
		resume := models.NewResume()

		if err := resume.CreateResume(cmd); err != nil {
			return err
		}

		return repository.Save(resume, m)
	})

	commandHandlers.Register(commands.ChangeResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		_ = c.(*commands.ChangeResume)
		return fmt.Errorf("does not implement")
	})

	return &ResumeCommandHandlers{
		cmdHandlers: commandHandlers,
	}
}
