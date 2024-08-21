package events

import (
	"SebStudy/domain/resume/values"
	"time"
)

type ResumeCreated struct {
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
	CreatedAt     time.Time
}

func NewResumeCreated(
	ResumeId values.ResumeId,
	FirstName values.FirstName,
	MiddleName values.MiddleName,
	LastName values.LastName,
	PhoneNumber values.PhoneNumber,
	Education values.Education,
	AboutMe values.AboutMe,
	Skills values.Skills,
	Photo values.Photo,
	Directions values.Directions,
	AboutProjects values.AboutProjects,
	Portfolio values.Portfolio,
	StudentGroup values.StudentGroup,
	CreatedAt time.Time,
) ResumeCreated {
	return ResumeCreated{
		ResumeId:      ResumeId,
		FirstName:     FirstName,
		MiddleName:    MiddleName,
		LastName:      LastName,
		PhoneNumber:   PhoneNumber,
		Education:     Education,
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
