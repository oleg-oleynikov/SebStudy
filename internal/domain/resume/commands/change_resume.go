package commands

import (
	"context"
	"fmt"
	"resume-server/config"
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/infrastructure"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/logger"
)

type ChangeResumeCommandHandler interface {
	Handle(ctx context.Context, cmd *ChangeResume, md infrastructure.CommandMetadata) error
}

type changeResumeCommandHandler struct {
	log logger.Logger
	cfg *config.Config
	es  eventsourcing.AggregateStore
}

func NewChangeResumeCommandHandler(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore) *changeResumeCommandHandler {
	return &changeResumeCommandHandler{
		log: log,
		cfg: cfg,
		es:  es,
	}
}

func (c *changeResumeCommandHandler) Handle(ctx context.Context, command *ChangeResume, md infrastructure.CommandMetadata) error {
	resume := models.NewResume()
	if err := c.es.Load(command.GetAggregateId(), resume); err != nil {
		return err
	}

	resume.Changed = false // TODO: по хорошему чет с этим сделать

	resume.ChangeResume(command.Education, command.AboutMe, command.Skills, command.BirthDate, command.Direction, command.AboutProjects, command.Portfolio)

	if !resume.Changed {
		return fmt.Errorf("no changes detected")
	}

	return c.es.Save(resume, md)
}
