package resume

import (
	"SebStudy/domain/resume/values"
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestResume(t *testing.T) {
	uuidR, err := uuid.NewV7()
	if err != nil {
		return
	}

	idResume := values.NewResumeId(uuidR.String())
	firstName, _ := values.NewFirstName("Алексей")
	middleName, _ := values.NewMiddleName("Валерьевич")
	lastName, _ := values.NewLastName("Кузнецов")
	phone, _ := values.NewPhoneNumber("79295132116")
	educ, _ := values.NewEducation("РКСИ")
	description, _ := values.NewAboutMe("хочу оливье")
	skill1, _ := values.NewSkill("Golang")
	skill2, _ := values.NewSkill("MySQL")
	photo, _ := values.NewPhoto([]byte{255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255, 0, 255}, "")
	direction, _ := values.NewDirection("back-end")
	projectDescription, _ := values.NewAboutProjects("вставай шан цунг, ты должен вспахать поле в 100000 гектаров")
	portfolioLink, _ := values.NewPortfolio("https://github.com")
	group, _ := values.NewStudentGroup("СА-33")

	// educs := values.Educations{}
	// educs.AppendEducations(*educ)

	skills := values.Skills{}
	skills.AppendSkills(*skill1, *skill2)

	directions := values.Directions{}
	directions.AppendDirection(*direction)

	resume1 := NewResume()
	fmt.Println(resume1)
	resume1.resumeId = idResume
	resume1.firstName = *firstName
	resume1.middleName = *middleName
	resume1.lastName = *lastName
	resume1.phoneNumber = *phone
	resume1.education = educ
	resume1.aboutMe = *description
	resume1.skills = skills
	resume1.photo = *photo
	resume1.directions = directions
	resume1.aboutProjects = *projectDescription
	resume1.portfolio = *portfolioLink
	resume1.studentGroup = *group

	fmt.Println(resume1.ToString())
}
