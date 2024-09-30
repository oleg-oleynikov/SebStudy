package queries

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
	"context"
)

type GetResumeByIdQueryHandler interface {
	Handle(ctx context.Context, query *GetResumeByIdQuery) (*models.ResumeProjection, error)
}

type getResumeByIdQueryHandler struct {
	log       logger.Logger
	cfg       *config.Config
	es        eventsourcing.AggregateStore
	mongoRepo repository.ResumeMongoRepository
}

func NewGetResumeByIdQueryHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeMongoRepository) *getResumeByIdQueryHandler {
	return &getResumeByIdQueryHandler{
		log:       log,
		cfg:       cfg,
		es:        es,
		mongoRepo: mongoRepo,
	}
}

func (q *getResumeByIdQueryHandler) Handle(ctx context.Context, query *GetResumeByIdQuery) (*models.ResumeProjection, error) {
	rp, err := q.mongoRepo.GetById(ctx, query.ResumeId)
	if err != nil {
		return nil, err
	}

	return rp, nil
}
