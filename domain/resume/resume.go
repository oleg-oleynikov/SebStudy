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

func NewResume(resumeId1 values.ResumeId, firstName1 values.FirstName, middleName1 values.MiddleName,
	lastName1 values.LastName, phoneNumber1 values.PhoneNumber, educations1 values.Educations,
	aboutMe1 values.AboutMe, skills1 values.Skills, photo1 values.Photo, directions1 values.Directions,
	aboutProjects1 values.AboutProjects, portfolio1 values.Portfolio, studentGroup1 values.StudentGroup) Resume {
	return Resume{
		resumeId:      resumeId1,
		firstName:     firstName1,
		middleName:    middleName1,
		lastName:      lastName1,
		phoneNumber:   phoneNumber1,
		educations:    educations1,
		aboutMe:       aboutMe1,
		skills:        skills1,
		photo:         photo1,
		directions:    directions1,
		aboutProjects: aboutProjects1,
		portfolio:     portfolio1,
		studentGroup:  studentGroup1,
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
