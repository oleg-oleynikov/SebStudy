package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
	"SebStudy/ports"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(repository ResumeRepository, eventSender ports.CeEventSender) *CommandHandlers {
	commandHandlers := &CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.SendResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(*commands.SendResume)

		// id := values.NewResumeId(cmd.ResumeId.Value)

		// resume, err := repository.Get(id)
		// if err != nil {
		// 	return err
		// }

		resume := NewResume()

		if err := resume.SendResume(cmd, eventSender); err != nil {
			return err
		}

		// resume.SendResume(cmd, eventSender)

		// log.Printf("Событие %v успешно обработано, версия агрегата: %d", cmd, resume.GetVersion())
		return nil
	})

	return commandHandlers
}
