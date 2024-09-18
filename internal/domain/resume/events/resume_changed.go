package events

import (
	"SebStudy/internal/domain/resume/values"
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
