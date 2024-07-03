package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
	"log"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(repository ResumeRepository) *CommandHandlers {
	commandHandlers := &CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.SendResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(*commands.SendResume)

		// id := values.NewResumeId(cmd.ResumeId.Value)

		// resume, err := repository.Get(id)
		// if err != nil {
		// 	return err
		// }

		resume := NewResume()

		resume.SendResume(cmd)

		// PushEvent

		// repository.Save(resume, m)
		log.Printf("Событие %v успешно обработано, версия агрегата: %d", cmd, resume.GetVersion())
		return nil
	})

	return commandHandlers
}
