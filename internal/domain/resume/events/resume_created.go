package events

import (
	"resume-server/internal/domain/resume/values"
	"time"
)

type ResumeCreated struct {
	ResumeId      string
	AboutMe       values.AboutMe
	Skills        values.Skills
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
	CreatedAt     time.Time
}

func NewResumeCreated(
	ResumeId string,
	AboutMe values.AboutMe,
	Skills values.Skills,
	Direction values.Direction,
	AboutProjects values.AboutProjects,
	Portfolio values.Portfolio,
	CreatedAt time.Time,
) ResumeCreated {
	return ResumeCreated{
		ResumeId:      ResumeId,
		AboutMe:       AboutMe,
		Skills:        Skills,
		Direction:     Direction,
		AboutProjects: AboutProjects,
		Portfolio:     Portfolio,
		CreatedAt:     CreatedAt,
	}
}
