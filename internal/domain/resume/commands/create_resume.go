package commands

import (
	"context"
	"resume-server/config"
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/infrastructure"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/logger"
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

	if err := resume.CreateResume(command.AboutMe, command.Skills, command.Direction, command.AboutProjects, command.Portfolio); err != nil {
		return err
	}

	return c.es.Save(resume, md)
}
