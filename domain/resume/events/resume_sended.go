package events

import (
	"SebStudy/domain/resume/values"
	"SebStudy/eventsourcing"
)

// eventsourcing "SebStudy/event_sourcing"

type ResumeSended struct {
	eventsourcing.AggregateRootBase

	ResumeId         values.ResumeId
	FirstName        values.FirstName
	MiddleName       values.MiddleName
	LastName         values.LastName
	PhoneNumber      values.PhoneNumber
	Educations       values.Educations
	AboutMe          values.AboutMe
	Skills           values.Skills
	Photo            values.Photo
	Directions       values.Directions
	AboutProjects    values.AboutProjects
	Portfolio        values.Portfolio
	StudentGroup     values.StudentGroup
	aggregateVersion uint
}

// func (rs *ResumeSended) GetResumeID() *values.ResumeId { return &rs.ResumeId }

// func (rs *ResumeSended) GetFirstName() *values.FirstName { return &rs.FirstName }

// func (rs *ResumeSended) GetMiddleName() *values.MiddleName { return &rs.MiddleName }

// func (rs *ResumeSended) GetLastName() *values.LastName { return &rs.LastName }

// func (rs *ResumeSended) GetPhoneNumber() *values.PhoneNumber { return &rs.PhoneNumber }

// func (rs *ResumeSended) GetEducations() *values.Educations { return &rs.Educations }

// func (rs *ResumeSended) GetAboutMe() *values.AboutMe { return &rs.AboutMe }

// func (rs *ResumeSended) GetSkills() *values.Skills { return &rs.Skills }

// func (rs *ResumeSended) GetPhoto() *values.Photo { return &rs.Photo }

// func (rs *ResumeSended) GetDirections() *values.Directions { return &rs.Directions }

// func (rs *ResumeSended) GetAboutProjects() *values.AboutProjects { return &rs.AboutProjects }

// func (rs *ResumeSended) GetPortfolio() *values.Portfolio { return &rs.Portfolio }

// func (rs *ResumeSended) GetStudentGroup() *values.StudentGroup { return &rs.StudentGroup }

func NewResumeSended(
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
) *ResumeSended {
	return &ResumeSended{
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

func (rs *ResumeSended) EventType() string { return "ResumeSended" }

func (rs *ResumeSended) EventVersion() uint8 { return 1 }

func (rs *ResumeSended) AggregateVersion() uint { return rs.aggregateVersion }
