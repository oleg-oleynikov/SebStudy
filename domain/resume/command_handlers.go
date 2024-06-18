package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
	"fmt"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(repository ResumeRepository) *CommandHandlers {
	commandHandlers := &CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandlers.Register(commands.SendResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(commands.SendResume)
		// TODO: Сделать репозиторий

		// id := values.NewResumeId(cmd.ResumeId.Value)

		// resume, err := repository.Get(id)
		// if err != nil {
		// 	return err
		// }
		resume := Resume{}
		fmt.Println("-------------------------------------------------")

		resume.SendResume(cmd)

		fmt.Print(resume, "\n", resume.GetVersion())
		// repository.Save(resume, m)
		return nil
	})

	return commandHandlers
}
