package events

import (
	"resume-server/internal/domain/resume/values"
	"time"
)

type ResumeChanged struct {
	ResumeId      string
	AboutMe       values.AboutMe
	Skills        values.Skills
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
	CreatedAt     time.Time
}

func NewResumeChanged(
	ResumeId string,
	AboutMe values.AboutMe,
	Skills values.Skills,
	Direction values.Direction,
	AboutProjects values.AboutProjects,
	Portfolio values.Portfolio,
	CreatedAt time.Time,
) ResumeChanged {
	return ResumeChanged{
		ResumeId:      ResumeId,
		AboutMe:       AboutMe,
		Skills:        Skills,
		Direction:     Direction,
		AboutProjects: AboutProjects,
		Portfolio:     Portfolio,
		CreatedAt:     CreatedAt,
	}
}
