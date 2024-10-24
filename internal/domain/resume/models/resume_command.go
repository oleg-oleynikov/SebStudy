package models

import (
	"resume-server/internal/domain/resume/events"
	"resume-server/internal/domain/resume/values"
	"time"
)

func (r *Resume) CreateResume(aboutMe values.AboutMe, skills values.Skills, direction values.Direction, aboutProjects values.AboutProjects, portfolio values.Portfolio) error {
	event := events.NewResumeCreated(r.GetId(), aboutMe, skills, direction, aboutProjects, portfolio, time.Now())
	r.Raise(event)
	return nil
}

func (r *Resume) ChangeResume(aboutMe values.AboutMe, skills values.Skills, direction values.Direction, aboutProjects values.AboutProjects, portfolio values.Portfolio) error {
	event := events.NewResumeChanged(r.GetId(), aboutMe, skills, direction, aboutProjects, portfolio, time.Now())
	r.Raise(event)
	return nil
}
