package commands

import (
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
)

type CreateResume struct {
	infrastructure.Command

	ResumeId      values.ResumeId
	FirstName     values.FirstName
	MiddleName    values.MiddleName
	LastName      values.LastName
	PhoneNumber   values.PhoneNumber
	Education     values.Education
	AboutMe       values.AboutMe
	Skills        values.Skills
	Photo         values.Photo
	Directions    values.Directions
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
	StudentGroup  values.StudentGroup
}

func NewCreateResume(
	resumeId values.ResumeId,
	firstName values.FirstName,
	middleName values.MiddleName,
	lastName values.LastName,
	phoneNumber values.PhoneNumber,
	education values.Education,
	aboutMe values.AboutMe,
	skills values.Skills,
	photo values.Photo,
	directions values.Directions,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
	studentGroup values.StudentGroup,
) *CreateResume {
	return &CreateResume{
		ResumeId:      resumeId,
		FirstName:     firstName,
		MiddleName:    middleName,
		LastName:      lastName,
		PhoneNumber:   phoneNumber,
		Education:     education,
		AboutMe:       aboutMe,
		Skills:        skills,
		Photo:         photo,
		Directions:    directions,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
		StudentGroup:  studentGroup,
	}
}
