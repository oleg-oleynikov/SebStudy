package models

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/values"
)

func (r *Resume) ResumeCreated(e events.ResumeCreated) {
	r.Id = e.ResumeId
	r.education = e.Education
	r.aboutMe = e.AboutMe
	r.skills = e.Skills
	r.birthDate = e.BirthDate
	r.direction = e.Direction
	r.aboutProjects = e.AboutProjects
	r.portfolio = e.Portfolio
}

func (r *Resume) ResumeChanged(e events.ResumeChanged) {
	education := e.Education.GetEducation()
	if education != "" && r.education.GetEducation() != education {
		r.education = e.Education
		r.Changed = true
	}

	aboutMe := e.AboutMe.GetAboutMe()
	if aboutMe != "" && r.aboutMe.GetAboutMe() != aboutMe {
		r.aboutMe = e.AboutMe
		r.Changed = true
	}

	newSkills := e.Skills.GetSkills()
	if len(newSkills) != 0 || !EqualSkills(newSkills, r.skills.GetSkills()) {
		r.skills = e.Skills
		r.Changed = true
	}

	birthDate := e.BirthDate.GetBirthDate()
	if birthDate.Compare(r.birthDate.GetBirthDate()) != 0 {
		r.birthDate = e.BirthDate
		r.Changed = true
	}

	direction := e.Direction.GetDirection()
	if direction != "" && r.direction.GetDirection() != direction {
		r.direction = e.Direction
		r.Changed = true
	}

	aboutProjects := e.AboutProjects.GetAboutProjects()
	if aboutProjects != "" && r.aboutProjects.GetAboutProjects() != aboutProjects {
		r.aboutProjects = e.AboutProjects
		r.Changed = true
	}

	portfolio := e.Portfolio.GetPortfolio()
	if portfolio != "" && r.portfolio.GetPortfolio() != portfolio {
		r.portfolio = e.Portfolio
		r.Changed = true
	}
}

func EqualSkills(newSkills []values.Skill, skills []values.Skill) bool {
	if len(newSkills) != len(skills) {
		return false
	}

	for i, skill := range newSkills {
		if skill.GetSkill() != skills[i].GetSkill() {
			return false
		}
	}

	return true
}
