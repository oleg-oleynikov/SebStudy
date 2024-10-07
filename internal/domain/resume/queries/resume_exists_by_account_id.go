package queries

import (
	"context"
	"resume-server/config"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/internal/infrastructure/repository"
	"resume-server/logger"
)

type ResumeExistsByAccountIdHandler interface {
	Handle(ctx context.Context, query *ResumeExistsByAccountIdQuery) (bool, error)
}

type getResumeExistsByAccountIdHandler struct {
	log       logger.Logger
	cfg       *config.Config
	es        eventsourcing.AggregateStore
	mongoRepo repository.ResumeRepository
}

func NewResumeExistsByAccountIdHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeRepository) *getResumeExistsByAccountIdHandler {
	return &getResumeExistsByAccountIdHandler{
		log:       log,
		cfg:       cfg,
		es:        es,
		mongoRepo: mongoRepo,
	}
}

func (q *getResumeExistsByAccountIdHandler) Handle(ctx context.Context, query *ResumeExistsByAccountIdQuery) (bool, error) {
	exists, err := q.mongoRepo.ResumeExistsByAccountId(ctx, query.AccountId)
	if err != nil {
		return false, err
	}

	return exists, nil
}
