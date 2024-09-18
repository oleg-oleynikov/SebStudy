package events

import (
	"SebStudy/internal/domain/resume/values"
	"time"
)

type ResumeCreated struct {
	ResumeId      string
	Education     values.Education
	AboutMe       values.AboutMe
	Skills        values.Skills
	BirthDate     values.BirthDate
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
	CreatedAt     time.Time
}

func NewResumeCreated(
	ResumeId string,
	Education values.Education,
	AboutMe values.AboutMe,
	Skills values.Skills,
	BirthDate values.BirthDate,
	Direction values.Direction,
	AboutProjects values.AboutProjects,
	Portfolio values.Portfolio,
	CreatedAt time.Time,
) ResumeCreated {
	return ResumeCreated{
		ResumeId:      ResumeId,
		Education:     Education,
		AboutMe:       AboutMe,
		Skills:        Skills,
		BirthDate:     BirthDate,
		Direction:     Direction,
		AboutProjects: AboutProjects,
		Portfolio:     Portfolio,
		CreatedAt:     CreatedAt,
	}
}
