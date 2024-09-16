package events

import (
	"SebStudy/internal/domain/resume/values"
	"time"
)

type ResumeCreated struct {
	ResumeId string
	// FirstName     values.FirstName
	// MiddleName    values.MiddleName
	// LastName      values.LastName
	// PhoneNumber   values.PhoneNumber
	Education values.Education
	AboutMe   values.AboutMe
	Skills    values.Skills
	// Photo         values.Photo
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
		Direction:     Direction,
		AboutProjects: AboutProjects,
		Portfolio:     Portfolio,
		CreatedAt:     CreatedAt,
	}
}
