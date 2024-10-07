package queries

import (
	"resume-server/config"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/internal/infrastructure/repository"
	"resume-server/logger"
)

type ResumeQueries struct {
	GetResumeByAccountId    GetResumeByAccountIdQueryHandler
	ResumeExistsByAccountId ResumeExistsByAccountIdHandler
	GetResumeById           GetResumeByIdQueryHandler
}

func NewResumeQueries(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeRepository) *ResumeQueries {
	return &ResumeQueries{
		GetResumeByAccountId:    NewGetResumeByAccountIdQueryHandler(log, cfg, es, mongoRepo),
		ResumeExistsByAccountId: NewResumeExistsByAccountIdHandler(log, cfg, es, mongoRepo),
		GetResumeById:           NewGetResumeByIdQueryHandler(log, cfg, es, mongoRepo),
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

type GetResumeByIdQuery struct {
	ResumeId string
}

func NewGetResumeByIdQuery(resumeId string) *GetResumeByIdQuery {
	return &GetResumeByIdQuery{ResumeId: resumeId}
}
