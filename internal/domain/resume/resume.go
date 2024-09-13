package resume

import (
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure/eventsourcing"
	"time"
)

type Resume struct {
	eventsourcing.AggregateRootBase

	firstName     values.FirstName
	middleName    values.MiddleName
	lastName      values.LastName
	phoneNumber   values.PhoneNumber
	education     values.Education
	aboutMe       values.AboutMe
	skills        values.Skills
	photo         values.Photo
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
	r.firstName = e.FirstName
	r.middleName = e.MiddleName
	r.lastName = e.LastName
	r.phoneNumber = e.PhoneNumber
	r.education = e.Education
	r.aboutMe = e.AboutMe
	r.skills = e.Skills
	r.photo = e.Photo
	r.direction = e.Direction
	r.aboutProjects = e.AboutProjects
	r.portfolio = e.Portfolio
}

func (r *Resume) CreateResume(c *commands.CreateResume) error {
	// id := r.GenerateUuidWithoutDashes()
	r.Raise(events.NewResumeCreated(r.GenerateUuidWithoutDashes(), c.FirstName, c.MiddleName, c.LastName, c.PhoneNumber, c.Education, c.AboutMe, c.Skills, c.Photo, c.Direction, c.AboutProjects, c.Portfolio, time.Now()))
	return nil
}
