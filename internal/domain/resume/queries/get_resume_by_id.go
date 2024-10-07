package queries

import (
	"context"
	"resume-server/config"
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/internal/infrastructure/repository"
	"resume-server/logger"
)

type GetResumeByIdQueryHandler interface {
	Handle(ctx context.Context, query *GetResumeByIdQuery) (*models.ResumeProjection, error)
}

type getResumeByIdQueryHandler struct {
	log       logger.Logger
	cfg       *config.Config
	es        eventsourcing.AggregateStore
	mongoRepo repository.ResumeRepository
}

func NewGetResumeByIdQueryHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeRepository) *getResumeByIdQueryHandler {
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
