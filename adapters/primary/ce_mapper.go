package primary

import (
	"SebStudy/domain/resume/commands"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	"context"
	"encoding/base64"
	"fmt"

	pb "SebStudy/proto/resume"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
)

type EventType string

const (
	EVENT   EventType = "event"
	COMMAND EventType = "command"
)

type CeMapper struct {
	handlers          map[string]func(context.Context, cloudevents.Event) (interface{}, error)
	handlersTypeEvent map[string]EventType
}

func NewCeMapper() *CeMapper {
	c := &CeMapper{}
	c.handlers = make(map[string]func(context.Context, cloudevents.Event) (interface{}, error), 0)
	c.handlersTypeEvent = make(map[string]EventType, 0)

	c.Register("resume.send", toSendResume, COMMAND)

	c.Register("resume.sended", toResumeSended, EVENT)

	return c
}

func (cm *CeMapper) MapFromCloudEvent(ctx context.Context, c cloudevents.Event) (interface{}, error) {
	handler, err := cm.Get(c.Type())
	if err != nil {
		return nil, err
	}

	// if err := decodeBase64(c); err != nil {
	// 	return nil, err
	// }

	cmd, err := handler(ctx, c)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cm *CeMapper) Get(t string) (func(context.Context, cloudevents.Event) (interface{}, error), error) {
	if h, ex := cm.handlers[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for {%v} type doesnt exist", t)
}

func (c *CeMapper) Register(ceType string, f func(context.Context, cloudevents.Event) (interface{}, error), typeEvent EventType) {
	c.handlers[ceType] = f
	c.handlersTypeEvent[ceType] = typeEvent
}

func (c *CeMapper) GetEventType(ceType string) (EventType, error) {
	if typeEvent, ok := c.handlersTypeEvent[ceType]; ok {
		return typeEvent, nil
	}
	return EventType(""), fmt.Errorf("Type doesnot exist")
	// return
}

// func decodeBase64(c cloudevents.Event) error {
// 	protoBytes, err := base64.StdEncoding.DecodeString(string(c.DataEncoded))
// 	if err != nil {
// 		return err
// 	}

// 	return c.SetData("application/protobuf", protoBytes)
// }

func toSendResume(ctx context.Context, c cloudevents.Event) (interface{}, error) {

	var rs pb.ResumeSended
	bytes, err := base64.StdEncoding.DecodeString(string(c.DataEncoded))
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(bytes, &rs); err != nil {
		return nil, err
	}

	resumeID := values.NewResumeId(int(rs.GetResumeId()))

	firstName, err := values.NewFirstName(rs.GetFirstName())
	if err != nil {
		return nil, err
	}

	lastName, err := values.NewLastName(rs.GetLastName())
	if err != nil {
		return nil, err
	}

	middleName, err := values.NewMiddleName(rs.GetMiddleName())
	if err != nil {
		return nil, err
	}

	phoneNumber, err := values.NewPhoneNumber(rs.GetPhoneNumber())
	if err != nil {
		return nil, err
	}

	var educations values.Educations
	for i := 0; i < len(rs.Educations); i++ {
		data := rs.Educations[i]
		education, err := values.NewEducation(data.Education)
		if err != nil {
			return nil, err
		}
		educations.AppendEducations(*education)
	}

	aboutMe, err := values.NewAboutMe(rs.GetAboutMe())
	if err != nil {
		return nil, err
	}

	var skills values.Skills
	for i := 0; i < len(rs.Skills); i++ {
		data := rs.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			return nil, err
		}
		skills.AppendSkills(*skill)
	}

	photo, err := values.NewPhoto(rs.GetPhoto())
	if err != nil {
		return nil, err
	}

	var directions values.Directions
	for i := 0; i < len(rs.Directions); i++ {
		data := rs.Directions[i]
		direction, err := values.NewDirection(data.Direction)
		if err != nil {
			return nil, err
		}
		directions.AppendDirection(*direction)
	}

	aboutProjects, err := values.NewAboutProjects(rs.GetAboutProjects())
	if err != nil {
		return nil, err
	}

	portfolio, err := values.NewPortfolio(rs.GetPortfolio())
	if err != nil {
		return nil, err
	}

	studentGroup, err := values.NewStudentGroup(rs.GetStudentGroup())
	if err != nil {
		return nil, err
	}

	createdResume := commands.NewSendResume(
		*resumeID, *firstName, *middleName, *lastName, *phoneNumber,
		educations, *aboutMe, skills, *photo, directions,
		*aboutProjects, *portfolio, *studentGroup)

	return createdResume, nil
}

func toResumeSended(ctx context.Context, c cloudevents.Event) (interface{}, error) {
	var rs pb.ResumeSended
	bytes, err := base64.StdEncoding.DecodeString(string(c.DataEncoded))
	if err != nil {
		return nil, err
	}

	if err := proto.Unmarshal(bytes, &rs); err != nil {
		return nil, err
	}

	resumeID := values.NewResumeId(int(rs.GetResumeId()))

	firstName, err := values.NewFirstName(rs.GetFirstName())
	if err != nil {
		return nil, err
	}

	lastName, err := values.NewLastName(rs.GetLastName())
	if err != nil {
		return nil, err
	}

	middleName, err := values.NewMiddleName(rs.GetMiddleName())
	if err != nil {
		return nil, err
	}

	phoneNumber, err := values.NewPhoneNumber(rs.GetPhoneNumber())
	if err != nil {
		return nil, err
	}

	var educations values.Educations
	for i := 0; i < len(rs.Educations); i++ {
		data := rs.Educations[i]
		education, err := values.NewEducation(data.Education)
		if err != nil {
			return nil, err
		}
		educations.AppendEducations(*education)
	}

	aboutMe, err := values.NewAboutMe(rs.GetAboutMe())
	if err != nil {
		return nil, err
	}

	var skills values.Skills
	for i := 0; i < len(rs.Skills); i++ {
		data := rs.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			return nil, err
		}
		skills.AppendSkills(*skill)
	}

	photo, err := values.NewPhoto(rs.GetPhoto())
	if err != nil {
		return nil, err
	}

	var directions values.Directions
	for i := 0; i < len(rs.Directions); i++ {
		data := rs.Directions[i]
		direction, err := values.NewDirection(data.Direction)
		if err != nil {
			return nil, err
		}
		directions.AppendDirection(*direction)
	}

	aboutProjects, err := values.NewAboutProjects(rs.GetAboutProjects())
	if err != nil {
		return nil, err
	}

	portfolio, err := values.NewPortfolio(rs.GetPortfolio())
	if err != nil {
		return nil, err
	}

	studentGroup, err := values.NewStudentGroup(rs.GetStudentGroup())
	if err != nil {
		return nil, err
	}

	createdResume := events.NewResumeSended(
		*resumeID, *firstName, *middleName, *lastName, *phoneNumber,
		educations, *aboutMe, skills, *photo, directions,
		*aboutProjects, *portfolio, *studentGroup, c.Time())

	return createdResume, nil
}
