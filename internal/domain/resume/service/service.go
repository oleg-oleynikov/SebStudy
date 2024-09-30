package service

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/queries"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/internal/infrastructure/repository"
	"SebStudy/logger"
)

type ResumeService struct {
	Commands *commands.ResumeCommands
	Queries  *queries.ResumeQueries
}

func NewResumeService(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore, mongoRepo repository.ResumeRepository) *ResumeService {
	resumeCommands := commands.NewResumeCommands(log, cfg, es)
	resumeQueries := queries.NewResumeQueries(log, cfg, es, mongoRepo)
	return &ResumeService{
		Commands: resumeCommands,
		Queries:  resumeQueries,
	}
}
