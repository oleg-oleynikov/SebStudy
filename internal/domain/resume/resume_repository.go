package resume

import (
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
)

type ResumeRepository interface {
	Get(resumeId *values.ResumeId) (*Resume, error)
	Save(r *Resume, m infrastructure.CommandMetadata) error
}
