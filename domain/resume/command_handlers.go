package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/infrastructure"
)

type CommandHandlers struct {
	*infrastructure.CommandHandlerBase
}

func NewHandlers(resumeRepos ResumeRepository) CommandHandlers {
	commandHandler := CommandHandlers{infrastructure.NewCommandHandler()}

	commandHandler.Register(commands.SendResume{}, func(c infrastructure.Command, m infrastructure.CommandMetadata) error {
		cmd := c.(commands.SendResume)
		id := cmd.ResumeId
		resume, err := resumeRepos.GetResume(id.ToString())
		if err != nil {
			return err
		}

		resumeRepos.SaveResume(resume, m)
		return nil
	})

	return commandHandler
}
