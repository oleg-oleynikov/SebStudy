package resume

import (
	"SebStudy/domain/resume/values"
	"fmt"
	"testing"
)

func TestResume(t *testing.T) {
	idResume := values.NewResumeId(101)
	firstName, _ := values.NewFirstName("Алексей")
	middleName, _ := values.NewMiddleName("Валерьевич")
	lastName, _ := values.NewLastName("Кузнецов")
	phone, _ := values.NewPhoneNumber("79295132116")
	educ, _ := values.NewEducation("РКСИ")
	description, _ := values.NewAboutMe("хочу оливье")
	skill1, _ := values.NewSkill("Golang")
	skill2, _ := values.NewSkill("MySQL")
	photo, _ := values.NewPhoto("myphoto.png")
	direction, _ := values.NewDirection("back-end")
	projectDescription, _ := values.NewAboutProjects("вставай шан цунг, ты должен вспахать поле в 100000 гектаров")
	portfolioLink, _ := values.NewPortfolio("https://github.com")
	group, _ := values.NewStudentGroup("СА-33")

	educs := values.Educations{}
	educs.AppendEducation(*educ)

	skills := values.Skills{}
	skills.AppendSkill(*skill1)
	skills.AppendSkill(*skill2)

	directions := values.Directions{}
	directions.AppendDirection(*direction)

	resume1 := NewResume(*idResume, *firstName, *middleName, *lastName, *phone,
		educs, *description, skills, *photo, directions, *projectDescription, *portfolioLink, *group)

	fmt.Println(resume1.ToString())
}
