package resume

import (
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
)

type ResumeRepository interface {
	Get(resumeId *values.ResumeId) (*models.Resume, error)
	Save(r *models.Resume, m infrastructure.CommandMetadata) error
}
