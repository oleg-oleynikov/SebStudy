package resume

import (
	"SebStudy/domain/resume/values"
)

type ResumeRepository interface {
	Get(resumeId *values.ResumeId) (*Resume, error)
	// Save()
}
