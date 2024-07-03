package events

import (
	// eventsourcing "SebStudy/event_sourcing"
	"SebStudy/domain/resume/values"
	"time"
)

type ResumeSended struct {
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
	CreatedAt     time.Time
}

func NewResumeSended(
	ResumeId values.ResumeId,
	FirstName values.FirstName,
	MiddleName values.MiddleName,
	LastName values.LastName,
	PhoneNumber values.PhoneNumber,
	Educations values.Educations,
	AboutMe values.AboutMe,
	Skills values.Skills,
	Photo values.Photo,
	Directions values.Directions,
	AboutProjects values.AboutProjects,
	Portfolio values.Portfolio,
	StudentGroup values.StudentGroup,
	CreatedAt time.Time,
) ResumeSended {
	return ResumeSended{
		ResumeId:      ResumeId,
		FirstName:     FirstName,
		MiddleName:    MiddleName,
		LastName:      LastName,
		PhoneNumber:   PhoneNumber,
		Educations:    Educations,
		AboutMe:       AboutMe,
		Skills:        Skills,
		Photo:         Photo,
		Directions:    Directions,
		AboutProjects: AboutProjects,
		Portfolio:     Portfolio,
		StudentGroup:  StudentGroup,
		CreatedAt:     CreatedAt,
	}
}
