package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
	"SebStudy/ports"
	"time"
)

type Resume struct {
	infrastructure.AggregateRootBase

	resumeId      values.ResumeId
	firstName     values.FirstName
	middleName    values.MiddleName
	lastName      values.LastName
	phoneNumber   values.PhoneNumber
	education     values.Education
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
		AggregateRootBase: infrastructure.NewAggregateRootBase(),
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
		r.education.ToString(), "\n",
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
	r.Register(events.ResumeCreated{}, func(e interface{}) { r.ResumeCreated(e.(events.ResumeCreated)) })
}

func (r *Resume) ResumeCreated(e events.ResumeCreated) {
	r.firstName = e.FirstName
	r.middleName = e.MiddleName
	r.lastName = e.LastName
	r.phoneNumber = e.PhoneNumber
	r.education = e.Education
	r.aboutMe = e.AboutMe
	r.skills = e.Skills
	r.photo = e.Photo
	r.directions = e.Directions
	r.aboutProjects = e.AboutProjects
	r.portfolio = e.Portfolio
	r.studentGroup = e.StudentGroup
}

func (r *Resume) CreateResume(c *commands.CreateResume, sender ports.CeEventSender) error {

	e := events.NewResumeCreated(c.ResumeId, c.FirstName, c.MiddleName, c.LastName, c.PhoneNumber, c.Education, c.AboutMe, c.Skills, c.Photo, c.Directions, c.AboutProjects, c.Portfolio, c.StudentGroup, time.Now())

	if err := sender.SendEvent(e, "resume.sended", "domain/resume"); err != nil {
		return err
	}

	return nil
}
