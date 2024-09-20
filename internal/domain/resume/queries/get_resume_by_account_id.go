package queries

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
	"context"
)

type GetResumeByAccountIdQueryHandler interface {
	Handle(ctx context.Context, query *GetResumeByAccountIdQuery) (*models.ResumeProjection, error)
}

type getResumeByAccountIdQueryHandler struct {
	log       logger.Logger
	cfg       *config.Config
	es        eventsourcing.AggregateStore
	mongoRepo repository.ResumeMongoRepository
}

func NewGetResumeByAccountIdQueryHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeMongoRepository) *getResumeByAccountIdQueryHandler {
	return &getResumeByAccountIdQueryHandler{
		log:       log,
		cfg:       cfg,
		es:        es,
		mongoRepo: mongoRepo,
	}
}

func (q *getResumeByAccountIdQueryHandler) Handle(ctx context.Context, query *GetResumeByAccountIdQuery) (*models.ResumeProjection, error) {
	return nil, nil
}
