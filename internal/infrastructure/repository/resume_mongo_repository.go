package repository

import (
	"SebStudy/internal/domain/resume/models"
	"context"
)

type ResumeMongoRepository interface {
	Insert(ctx context.Context, resume *models.ResumeProjection) error
	Update(ctx context.Context, resume *models.ResumeProjection) error
}
