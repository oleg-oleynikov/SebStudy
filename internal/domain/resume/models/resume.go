package models

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure/eventsourcing"
)

type Resume struct {
	eventsourcing.AggregateRootBase

	education     values.Education
	aboutMe       values.AboutMe
	skills        values.Skills
	birthDate     values.BirthDate
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
	r.Register(events.ResumeChanged{}, func(e interface{}) { r.ResumeChanged(e.(events.ResumeChanged)) })
}
