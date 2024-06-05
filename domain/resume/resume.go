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

func (r *Resume) ToString() []string {
	return []string{
		r.resumeId.ToString(), "\n",
		r.firstName.ToString(), "\n",
		r.middleName.ToString(), "\n",
		r.lastName.ToString(), "\n",
		r.phoneNumber.ToString(), "\n",
		r.educations.ToString(), "\n",
		r.aboutMe.ToString(), "\n",
		r.skills.ToString(), "\n",
		r.photo.ToString(), "\n",
		r.directions.ToString(), "\n",
		r.aboutProjects.ToString(), "\n",
		r.portfolio.ToString(), "\n",
    r.studentGroup.ToString(), "\n",
  }
}
