package models

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/values"
	"time"
)

func (r *Resume) CreateResume(education values.Education, aboutMe values.AboutMe, skills values.Skills, birthDate values.BirthDate, direction values.Direction, aboutProjects values.AboutProjects, portfolio values.Portfolio) error {
	event := events.NewResumeCreated(r.GetId(), education, aboutMe, skills, birthDate, direction, aboutProjects, portfolio, time.Now())
	r.Raise(event)
	return nil
}

func (r *Resume) ChangeResume(education values.Education, aboutMe values.AboutMe, skills values.Skills, birthDate values.BirthDate, direction values.Direction, aboutProjects values.AboutProjects, portfolio values.Portfolio) error {
	event := events.NewResumeChanged(r.GetId(), education, aboutMe, skills, birthDate, direction, aboutProjects, portfolio, time.Now())
	r.Raise(event)
	return nil
}
