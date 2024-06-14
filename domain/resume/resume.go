package resume

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	eventsourcing "SebStudy/eventsourcing"
)

type Resume struct {
	eventsourcing.AggregateRootBase

	resumeId      values.ResumeId
	firstName     values.FirstName
	middleName    values.MiddleName
	lastName      values.LastName
	phoneNumber   values.PhoneNumber
	education     values.Education
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
	education values.Education,
	aboutMe values.AboutMe,
	skills values.Skills,
	photo values.Photo,
	directions values.Directions,
	aboutProjects values.AboutProjects,
	portfolio values.Portfolio,
	studentGroup values.StudentGroup,
) *Resume {
	r := &Resume{
		resumeId:      resumeId,
		firstName:     firstName,
		middleName:    middleName,
		lastName:      lastName,
		phoneNumber:   phoneNumber,
		education:     education,
		aboutMe:       aboutMe,
		skills:        skills,
		photo:         photo,
		directions:    directions,
		aboutProjects: aboutProjects,
		portfolio:     portfolio,
		studentGroup:  studentGroup,
	}

	r.Register(events.ResumeSended{}, func(e interface{}) { r.ResumeSended(e.(events.ResumeSended)) })

	return r
}

func (r *Resume) ToString() []string {
	return []string{
		r.resumeId.ToString(), "\n",
		r.firstName.ToString(), "\n",
		r.middleName.ToString(), "\n",
		r.lastName.ToString(), "\n",
		r.phoneNumber.ToString(), "\n",
		r.education.ToString(), "\n",
		r.aboutMe.ToString(), "\n",
		r.skills.ToString(), "\n",
		r.photo.ToString(), "\n",
		r.directions.ToString(), "\n",
		r.aboutProjects.ToString(), "\n",
		r.portfolio.ToString(), "\n",
		r.studentGroup.ToString(), "\n",
	}
}

func (r *Resume) ResumeSended(e events.ResumeSended) {
	// Проставляем все поля пустому объекту Resume
	return
}

func (r *Resume) SendResume(c commands.SendResume) {
	// Формируем ResumeSended event и делаем r.Raise(ResumeSended{})
	return
}
