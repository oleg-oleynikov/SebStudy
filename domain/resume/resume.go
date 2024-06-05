package resume

import (
	"SebStudy/domain/resume/values"
	eventsourcing "SebStudy/event_sourcing"
)

type Resume struct {
	eventsourcing.AggregateRootBase

	resumeId      values.ResumeId
	firstName     values.FirstName
	middleName    values.MiddleName
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

func NewResume(
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
) *Resume {
	return &Resume{
		resumeId:      resumeId,
		firstName:     firstName,
		middleName:    middleName,
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
