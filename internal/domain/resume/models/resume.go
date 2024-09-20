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

func (r *Resume) HasChanged(newResume *Resume) bool {
	if newResume.education.GetEducation() != "" && r.education.GetEducation() != newResume.education.GetEducation() {
		return true
	}

	if newResume.aboutMe.GetAboutMe() != "" && r.aboutMe.GetAboutMe() != newResume.aboutMe.GetAboutMe() {
		return true
	}

	if len(newResume.skills.GetSkills()) > 0 && !EqualSkills(r.skills.GetSkills(), newResume.skills.GetSkills()) {
		return true
	}

	if !newResume.birthDate.GetBirthDate().IsZero() && r.birthDate.GetBirthDate() != newResume.birthDate.GetBirthDate() {
		return true
	}

	if newResume.direction.GetDirection() != "" && r.direction.GetDirection() != newResume.direction.GetDirection() {
		return true
	}

	if newResume.aboutProjects.GetAboutProjects() != "" && r.aboutProjects.GetAboutProjects() != newResume.aboutProjects.GetAboutProjects() {
		return true
	}

	if newResume.portfolio.GetPortfolio() != "" && r.portfolio.GetPortfolio() != newResume.portfolio.GetPortfolio() {
		return true
	}

	return false
}

func (r *Resume) Copy() *Resume {
	return &Resume{
		AggregateRootBase: r.AggregateRootBase,
		education:         r.education,
		aboutMe:           r.aboutMe,
		skills:            r.skills,
		birthDate:         r.birthDate,
		direction:         r.direction,
		aboutProjects:     r.aboutProjects,
		portfolio:         r.portfolio,
	}
}
