package commands

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/logger"
	"context"
	"fmt"
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

	resumeBefore := resume.Copy()
	resume.ChangeResume(command.Education, command.AboutMe, command.Skills, command.BirthDate, command.Direction, command.AboutProjects, command.Portfolio)

	if !resumeBefore.HasChanged(resume) {
		return fmt.Errorf("no changes detected")
	}
	return c.es.Save(resume, md)
}
