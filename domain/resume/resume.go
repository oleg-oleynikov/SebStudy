package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	eventsourcing "SebStudy/eventsourcing"
	"time"
)

type Resume struct {
	eventsourcing.AggregateRootBase

	resumeId      values.ResumeId
	firstName     values.FirstName
	middleName    values.MiddleName
	lastName      values.LastName
	phoneNumber   values.PhoneNumber
	educations    values.Educations
	aboutMe       values.AboutMe
	skills        values.Skills
	photo         values.Photo
	directions    values.Directions
	aboutProjects values.AboutProjects
	portfolio     values.Portfolio
	studentGroup  values.StudentGroup
}

func NewResume() *Resume {
	r := &Resume{
		AggregateRootBase: eventsourcing.NewAggregateRootBase(),
	}

	r.registerHandlers()

	return r
}

func (r *Resume) ToString() []string {
	return []string{
		r.resumeId.ToString(), "\n",
		r.firstName.ToString(), "\n",
		r.middleName.ToString(), "\n",
		r.lastName.ToString(), "\n",
		r.phoneNumber.ToString(), "\n",
		r.educations.ToString(), "\n",
		r.aboutMe.ToString(), "\n",
		r.skills.ToString(), "\n",
		r.photo.ToString(), "\n",
		r.directions.ToString(), "\n",
		r.aboutProjects.ToString(), "\n",
		r.portfolio.ToString(), "\n",
		r.studentGroup.ToString(), "\n",
	}
}

func (r *Resume) registerHandlers() {
	r.Register(events.ResumeSended{}, func(e interface{}) { r.ResumeSended(e.(events.ResumeSended)) })
}

func (r *Resume) ResumeSended(e events.ResumeSended) {
	r.resumeId = e.ResumeId
	r.firstName = e.FirstName
	r.middleName = e.MiddleName
	r.lastName = e.LastName
	r.phoneNumber = e.PhoneNumber
	r.educations = e.Educations
	r.aboutMe = e.AboutMe
	r.skills = e.Skills
	r.photo = e.Photo
	r.directions = e.Directions
	r.aboutProjects = e.AboutProjects
	r.portfolio = e.Portfolio
	r.studentGroup = e.StudentGroup
}

func (r *Resume) SendResume(c *commands.SendResume) {
	// TODO: Исправить дело с time.Now();
	e := events.NewResumeSended(c.ResumeId, c.FirstName, c.MiddleName, c.LastName, c.PhoneNumber, c.Educations, c.AboutMe, c.Skills, c.Photo, c.Directions, c.AboutProjects, c.Portfolio, c.StudentGroup, time.Now())
	r.Raise(e)
}
