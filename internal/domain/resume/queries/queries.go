package queries

import (
	"SebStudy/config"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
)

type ResumeQueries struct {
	GetResumeByAccountId    GetResumeByAccountIdQueryHandler
	ResumeExistsByAccountId ResumeExistsByAccountIdHandler
}

func NewResumeQueries(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeMongoRepository) *ResumeQueries {
	return &ResumeQueries{
		GetResumeByAccountId:    NewGetResumeByAccountIdQueryHandler(log, cfg, es, mongoRepo),
		ResumeExistsByAccountId: NewResumeExistsByAccountIdHandler(log, cfg, es, mongoRepo),
	}
}

type GetResumeByAccountIdQuery struct {
	AccountId string
}

func NewGetResumeByAccountIdQuery(accountId string) *GetResumeByAccountIdQuery {
	return &GetResumeByAccountIdQuery{AccountId: accountId}
}

type ResumeExistsByAccountIdQuery struct {
	AccountId string
}

func NewResumeExistsByAccountId(accountId string) *ResumeExistsByAccountIdQuery {
	return &ResumeExistsByAccountIdQuery{AccountId: accountId}
}
