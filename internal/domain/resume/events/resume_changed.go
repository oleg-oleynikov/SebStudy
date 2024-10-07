package events

import (
	"resume-server/internal/domain/resume/values"
	"time"
)

type ResumeChanged struct {
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

func NewResumeChanged(
	ResumeId string,
	Education values.Education,
	AboutMe values.AboutMe,
	Skills values.Skills,
	BirthDate values.BirthDate,
	Direction values.Direction,
	AboutProjects values.AboutProjects,
	Portfolio values.Portfolio,
	CreatedAt time.Time,
) ResumeChanged {
	return ResumeChanged{
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
