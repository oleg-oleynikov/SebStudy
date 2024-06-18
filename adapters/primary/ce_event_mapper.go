package primary

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/domain/resume/values"
	"fmt"
	"reflect"

	pb "SebStudy/proto/resume"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

type CeMapper struct {
	handlers map[string]func(cloudevents.Event) (interface{}, error)
}

func NewCeMapper() *CeMapper {
	c := &CeMapper{}
	c.handlers = make(map[string]func(cloudevents.Event) (interface{}, error), 0)

	// Регистрация handler а, для преобразования cloudevents.Event в объект понятный приложению
	c.Register("resume.send", toSendResume)

	return c
}

func (cm *CeMapper) MapToCommand(c cloudevents.Event) (interface{}, error) {
	handler, err := cm.Get(c.Type())
	if err != nil {
		return nil, err
	}

	cmd, err := handler(c)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cm *CeMapper) Get(t string) (func(cloudevents.Event) (interface{}, error), error) {
	if h, ex := cm.handlers[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for %s type doesnt exist", t)
}

func (c *CeMapper) Register(ceType string, f func(cloudevents.Event) (interface{}, error)) {
	c.handlers[ceType] = f
}

func toSendResume(c cloudevents.Event) (interface{}, error) {
	var de pb.ResumeSended
	proto.Unmarshal(c.Data(), &de)
	fmt.Printf("Event data: %v\nEvent type: %s\n", &de, reflect.ValueOf(&de))
	//TODO: Заполнить поля cmd из cloudevents.Event (c)
	// proto.Unmarshal(c.Data())

	// if reflect.ValueOf(*de).IsZero() {
	// 	return nil, errors.New("struct is empty")
	// }
	resumeID := values.NewResumeId(int(de.GetResumeId()))

	firstName, err := values.NewFirstName(de.GetFirstName())
	if err != nil {
		return nil, err
	}

	lastName, err := values.NewLastName(de.GetLastName())
	if err != nil {
		return nil, err
	}

	middleName, err := values.NewMiddleName(de.GetMiddleName())
	if err != nil {
		return nil, err
	}

	phoneNumber, err := values.NewPhoneNumber(de.GetPhoneNumber())
	if err != nil {
		return nil, err
	}

	var educations values.Educations
	for i := 0; i < len(de.Educations); i++ {
		data := de.Educations[i]
		education, err := values.NewEducation(data.Education)
		if err != nil {
			return nil, err
		}
		educations.AppendEducation(*education)
	}

	aboutMe, err := values.NewAboutMe(de.GetAboutMe())
	if err != nil {
		return nil, err
	}

	var skills values.Skills
	for i := 0; i < len(de.Skills); i++ {
		data := de.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			return nil, err
		}
		skills.AppendSkill(*skill)
	}

	photo, err := values.NewPhoto(de.GetPhoto())
	if err != nil {
		return nil, err
	}

	var directions values.Directions
	for i := 0; i < len(de.Directions); i++ {
		data := de.Directions[i]
		direction, err := values.NewDirection(data.Direction)
		if err != nil {
			return nil, err
		}
		directions.AppendDirection(*direction)
	}

	aboutProjects, err := values.NewAboutProjects(de.GetAboutProjects())
	if err != nil {
		return nil, err
	}

	portfolio, err := values.NewPortfolio(de.GetPortfolio())
	if err != nil {
		return nil, err
	}

	studentGroup, err := values.NewStudentGroup(de.GetStudentGroup())
	if err != nil {
		return nil, err
	}

	createdResume := commands.NewSendResume(
		*resumeID, *firstName, *middleName, *lastName, *phoneNumber,
		educations, *aboutMe, skills, *photo, directions,
		*aboutProjects, *portfolio, *studentGroup)

	return createdResume, nil
}
