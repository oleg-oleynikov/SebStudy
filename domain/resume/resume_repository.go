package resume

import (
	"SebStudy/domain/resume/values"
)

type ResumeRepository interface {
	// Save(resume *Resume, metadata infrastructure.CommandMetadata) error
	Get(resumeId *values.ResumeId) (*Resume, error)
}
