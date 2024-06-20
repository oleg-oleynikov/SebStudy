package events

import (
	// eventsourcing "SebStudy/event_sourcing"
	"SebStudy/domain/resume/values"
	"time"
)

// eventsourcing "SebStudy/event_sourcing"

type ResumeSended struct {
	// eventsourcing.DomainEvent

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
	// Также здесь дописать
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
