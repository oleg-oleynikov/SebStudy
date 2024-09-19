package commands

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/logger"
	"context"
)

type CreateResumeCommandHandler interface {
	Handle(ctx context.Context, command *CreateResume, md infrastructure.CommandMetadata) error
}

type createResumeCommandHandler struct {
	log logger.Logger
	cfg *config.Config
	es  eventsourcing.AggregateStore
}

func NewCreateResumeCommandHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore) *createResumeCommandHandler {
	return &createResumeCommandHandler{
		log: log,
		cfg: cfg,
		es:  es,
	}
}

func (c *createResumeCommandHandler) Handle(ctx context.Context, command *CreateResume, md infrastructure.CommandMetadata) error {
	resume := models.NewResume()
	resume.Id = command.GetAggregateId()

	if err := resume.CreateResume(command.Education, command.AboutMe, command.Skills, command.BirthDate, command.Direction, command.AboutProjects, command.Portfolio); err != nil {
		return err
	}

	return c.es.Save(resume, md)
}
