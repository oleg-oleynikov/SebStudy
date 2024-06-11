package resume

import (
	"SebStudy/infrastructure"
)

type ResumeRepository interface {
	SaveResume(resume *Resume, metadata infrastructure.CommandMetadata)
	GetResume(resumeId string) (*Resume, error)
}
