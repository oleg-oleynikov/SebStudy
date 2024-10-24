package commands

import (
	"resume-server/config"
	"resume-server/internal/domain/resume/values"
	"resume-server/internal/infrastructure"
	"resume-server/internal/infrastructure/eventsourcing"
	"resume-server/logger"
)

type ResumeCommands struct {
	CreateResume CreateResumeCommandHandler
	ChangeResume ChangeResumeCommandHandler
}

func NewResumeCommands(log logger.Logger, cfg *config.Config, es eventsourcing.AggregateStore) *ResumeCommands {
	createResume := NewCreateResumeCommandHandler(log, cfg, es)
	changeResume := NewChangeResumeCommandHandler(log, cfg, es)
	return &ResumeCommands{
		CreateResume: createResume,
		ChangeResume: changeResume,
	}
}

type CreateResume struct {
	infrastructure.BaseCommand

	AboutMe       values.AboutMe
	Skills        values.Skills
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
}

func NewCreateResume(
	resumeId string,
	aboutMe values.AboutMe,
	skills values.Skills,
	direction values.Direction,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
) *CreateResume {
	return &CreateResume{
		BaseCommand:   infrastructure.NewBaseCommand(resumeId),
		AboutMe:       aboutMe,
		Skills:        skills,
		Direction:     direction,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
	}
}

type ChangeResume struct {
	infrastructure.BaseCommand
	AboutMe       values.AboutMe
	Skills        values.Skills
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
}

func NewChangeResume(
	resumeId string,
	aboutMe values.AboutMe,
	skills values.Skills,
	direction values.Direction,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
) *ChangeResume {
	return &ChangeResume{
		BaseCommand:   infrastructure.NewBaseCommand(resumeId),
		AboutMe:       aboutMe,
		Skills:        skills,
		Direction:     direction,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
	}
}
