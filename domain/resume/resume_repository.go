package resume

import (
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
)

type ResumeRepository interface {
	Get(resumeId *values.ResumeId) (*Resume, error)
	Save(r *Resume, m infrastructure.CommandMetadata) error
}
