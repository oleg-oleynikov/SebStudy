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

		// repository.Save(resume, m)
		log.Printf("Событие %s успешно обработано, версия агрегата: %s", cmd, resume.GetVersion())
		return nil
	})

	return commandHandlers
}
