package resume

import (
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
)

type ResumeRepository interface {
	Save(resume *Resume, metadata infrastructure.CommandMetadata) error
	Get(resumeId *values.ResumeId) (*Resume, error)
}
