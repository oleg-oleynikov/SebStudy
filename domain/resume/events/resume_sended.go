package events

import (
	// eventsourcing "SebStudy/event_sourcing"
	"SebStudy/domain/resume/values"
	"time"
)

type ResumeSended struct {
	// ResumeId      values.ResumeId
	// FirstName     values.FirstName
	// MiddleName    values.MiddleName
	// LastName      values.LastName
	// PhoneNumber   values.PhoneNumber
	// Education     values.Education
	// AboutMe       values.AboutMe
	// Skills        values.Skills
	// Photo         values.Photo
	// Directions    values.Directions
	// AboutProjects values.AboutProjects
	// Portfolio     values.Portfolio
	// StudentGroup  values.StudentGroup
	// CreatedAt     time.Time
	ResumeId      values.ResumeId      `json:"resume_id"`
	FirstName     values.FirstName     `json:"first_name"`
	MiddleName    values.MiddleName    `json:"middle_name"`
	LastName      values.LastName      `json:"last_name"`
	PhoneNumber   values.PhoneNumber   `json:"phone_number"`
	Education     values.Education     `json:"education"`
	AboutMe       values.AboutMe       `json:"about_me"`
	Skills        values.Skills        `json:"skills"`
	Photo         values.Photo         `json:"photo"`
	Directions    values.Directions    `json:"directions"`
	AboutProjects values.AboutProjects `json:"about_projects"`
	Portfolio     values.Portfolio     `json:"portfolio"`
	StudentGroup  values.StudentGroup  `json:"student_group"`
	CreatedAt     time.Time            `json:"created_at"`
}

func NewResumeSended(
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
) ResumeSended {
	return ResumeSended{
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
