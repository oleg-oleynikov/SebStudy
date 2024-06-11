package commands

import (
	"SebStudy/domain/resume/values"
	"SebStudy/eventsourcing"
)

type SendResume struct {
	eventsourcing.AggregateRootBase

	ResumeId      values.ResumeId
	FirstName     values.FirstName
	MiddleName    values.MiddleName
	lastName      values.LastName
	phoneNumber   values.PhoneNumber
	educations    values.Educations
	aboutMe       values.AboutMe
	skills        values.Skills
	photo         values.Photo
	directions    values.Directions
	aboutProjects values.AboutProjects
	portfolio     values.Portfolio
	studentGroup  values.StudentGroup
}

func NewSendResume(
	resumeId values.ResumeId,
	firstName values.FirstName,
	middleName values.MiddleName,
	lastName values.LastName,
	phoneNumber values.PhoneNumber,
	educations values.Educations,
	aboutMe values.AboutMe,
	skills values.Skills,
	photo values.Photo,
	directions values.Directions,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
	studentGroup values.StudentGroup,
) *SendResume {
	return &SendResume{
		ResumeId:      resumeId,
		FirstName:     firstName,
		MiddleName:    middleName,
		lastName:      lastName,
		phoneNumber:   phoneNumber,
		educations:    educations,
		aboutMe:       aboutMe,
		skills:        skills,
		photo:         photo,
		directions:    directions,
		aboutProjects: aboutProjects,
		portfolio:     portfolio,
		studentGroup:  studentGroup,
	}
}
