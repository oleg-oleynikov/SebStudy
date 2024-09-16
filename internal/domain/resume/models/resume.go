package models

import (
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure/eventsourcing"
	"time"
)

type Resume struct {
	eventsourcing.AggregateRootBase

	education     values.Education
	aboutMe       values.AboutMe
	skills        values.Skills
	direction     values.Direction
	aboutProjects values.AboutProjects
	portfolio     values.Portfolio
}

func NewResume() *Resume {
	r := &Resume{
		AggregateRootBase: eventsourcing.NewAggregateRootBase(),
	}

	r.registerHandlers()

	return r
}

func (r *Resume) registerHandlers() {
	r.Register(events.ResumeCreated{}, func(e interface{}) { r.ResumeCreated(e.(events.ResumeCreated)) })
}

func (r *Resume) ResumeCreated(e events.ResumeCreated) {
	r.Id = e.ResumeId
	r.education = e.Education
	r.aboutMe = e.AboutMe
	r.skills = e.Skills
	r.direction = e.Direction
	r.aboutProjects = e.AboutProjects
	r.portfolio = e.Portfolio
}

func (r *Resume) CreateResume(c *commands.CreateResume) error {
	// id := r.GenerateUuidWithoutDashes()
	r.Raise(events.NewResumeCreated(r.GenerateUuidWithoutDashes(), c.Education, c.AboutMe, c.Skills, c.Direction, c.AboutProjects, c.Portfolio, time.Now()))
	return nil
}
