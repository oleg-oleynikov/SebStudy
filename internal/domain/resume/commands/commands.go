package commands

import (
	"SebStudy/config"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
	"SebStudy/internal/infrastructure/eventsourcing"
	"SebStudy/logger"
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

	Education     values.Education
	AboutMe       values.AboutMe
	Skills        values.Skills
	BirthDate     values.BirthDate
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
}

func NewCreateResume(
	resumeId string,
	education values.Education,
	aboutMe values.AboutMe,
	skills values.Skills,
	birthDate values.BirthDate,
	direction values.Direction,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
) *CreateResume {
	return &CreateResume{
		BaseCommand:   infrastructure.NewBaseCommand(resumeId),
		Education:     education,
		AboutMe:       aboutMe,
		Skills:        skills,
		BirthDate:     birthDate,
		Direction:     direction,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
	}
}

type ChangeResume struct {
	infrastructure.BaseCommand
	Education     values.Education
	AboutMe       values.AboutMe
	Skills        values.Skills
	BirthDate     values.BirthDate
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
}

func NewChangeResume(
	resumeId string,
	education values.Education,
	aboutMe values.AboutMe,
	skills values.Skills,
	birthDate values.BirthDate,
	direction values.Direction,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
) *ChangeResume {
	return &ChangeResume{
		BaseCommand:   infrastructure.NewBaseCommand(resumeId),
		Education:     education,
		AboutMe:       aboutMe,
		Skills:        skills,
		BirthDate:     birthDate,
		Direction:     direction,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
	}
}
