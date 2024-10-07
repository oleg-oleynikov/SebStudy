package resume

import (
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/domain/resume/values"
	"resume-server/internal/infrastructure"
)

type ResumeRepository interface {
	Get(resumeId *values.ResumeId) (*models.Resume, error)
	Save(r *models.Resume, m infrastructure.CommandMetadata) error
}
