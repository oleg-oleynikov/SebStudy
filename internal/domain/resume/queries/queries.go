package queries

import (
	"SebStudy/config"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
)

type ResumeQueries struct {
}

func NewResumeQueries(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeMongoRepository) *ResumeQueries {
	return &ResumeQueries{}
}
