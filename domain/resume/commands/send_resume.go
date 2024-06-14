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
	LastName      values.LastName
	PhoneNumber   values.PhoneNumber
	Educations    values.Educations
	AboutMe       values.AboutMe
	Skills        values.Skills
	Photo         values.Photo
	Directions    values.Directions
	AboutProjects values.AboutProjects
	Portfolio     values.Portfolio
	StudentGroup  values.StudentGroup
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
		LastName:      lastName,
		PhoneNumber:   phoneNumber,
		Educations:    educations,
		AboutMe:       aboutMe,
		Skills:        skills,
		Photo:         photo,
		Directions:    directions,
		AboutProjects: aboutProjects,
		Portfolio:     portfolio,
		StudentGroup:  studentGroup,
	}
}
