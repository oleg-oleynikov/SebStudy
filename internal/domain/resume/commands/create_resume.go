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
	Direction     values.Direction
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
}

func NewCreateResume(
	education values.Education,
	aboutMe values.AboutMe,
	skills values.Skills,
	direction values.Direction,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
) *CreateResume {
	return &CreateResume{
		Education:     education,
		AboutMe:       aboutMe,
		Skills:        skills,
		Direction:     direction,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
	}
}
