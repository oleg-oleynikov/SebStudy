package queries

import (
	"SebStudy/config"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
)

type ResumeQueries struct {
	GetResumeByAccountId GetResumeByAccountIdQueryHandler
}

func NewResumeQueries(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeMongoRepository) *ResumeQueries {
	return &ResumeQueries{
		GetResumeByAccountId: NewGetResumeByAccountIdQueryHandler(log, cfg, es, mongoRepo),
	}
}

type GetResumeByAccountIdQuery struct {
	AccountId string
}

func NewGetResumeByAccountIdQuery(accountId string) *GetResumeByAccountIdQuery {
	return &GetResumeByAccountIdQuery{AccountId: accountId}
}
