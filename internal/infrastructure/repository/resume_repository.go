package repository

import (
	"context"
	"resume-server/internal/domain/resume/models"
)

type ResumeRepository interface {
	Insert(ctx context.Context, resume *models.ResumeProjection) error
	Update(ctx context.Context, resume *models.ResumeProjection) error

	GetByAccountId(ctx context.Context, accountId string) (*models.ResumeProjection, error)
	GetById(ctx context.Context, resumeId string) (*models.ResumeProjection, error)
	ResumeExistsByAccountId(ctx context.Context, accountId string) (bool, error)
}
