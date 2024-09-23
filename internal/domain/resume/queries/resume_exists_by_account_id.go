package queries

import (
	"SebStudy/config"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
	"context"
)

type ResumeExistsByAccountIdHandler interface {
	Handle(ctx context.Context, query *ResumeExistsByAccountIdQuery) (bool, error)
}

type getResumeExistsByAccountIdHandler struct {
	log       logger.Logger
	cfg       *config.Config
	es        eventsourcing.AggregateStore
	mongoRepo repository.ResumeMongoRepository
}

func NewResumeExistsByAccountIdHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeMongoRepository) *getResumeExistsByAccountIdHandler {
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
