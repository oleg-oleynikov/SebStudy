package commands

import (
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
)

type CreateResume struct {
	infrastructure.Command

	Education     values.Education
	AboutMe       values.AboutMe
	Skills        values.Skills
	BirthDate     values.BirthDate
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
}

func NewCreateResume(
	education values.Education,
	aboutMe values.AboutMe,
	skills values.Skills,
	birthDate values.BirthDate,
	direction values.Direction,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
) *CreateResume {
	return &CreateResume{
		Education:     education,
		AboutMe:       aboutMe,
		Skills:        skills,
		BirthDate:     birthDate,
		Direction:     direction,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
	}
}
