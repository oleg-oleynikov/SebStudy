package service

import (
	"resume-server/config"
	"resume-server/internal/domain/resume/commands"
	"resume-server/internal/domain/resume/queries"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/internal/infrastructure/repository"
	"resume-server/logger"
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
