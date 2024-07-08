package util

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
	"google.golang.org/protobuf/types/known/timestamppb"
)

type EventType string

type CeToEvent func(ctx context.Context, ce cloudevents.Event) (interface{}, error)
type EventToCe func(eventType, source string, e interface{}) (cloudevents.Event, error)

const (
	EVENT   EventType = "event"
	COMMAND EventType = "command"
)

type CeMapper struct {
	CeToEvent         map[string]CeToEvent
	EventToCe         map[string]EventToCe
	handlersTypeEvent map[string]EventType
}

func NewCeMapper() *CeMapper {
	c := &CeMapper{}
	c.CeToEvent = make(map[string]CeToEvent, 0)
	c.EventToCe = make(map[string]EventToCe, 0)
	c.handlersTypeEvent = make(map[string]EventType, 0)

	c.RegisterCommand("resume.send", toSendResume)

	c.RegisterEvent("resume.sended", toResumeSended, resumeSendedToCloudEvent)

	return c
}

func (cm *CeMapper) MapToEvent(ctx context.Context, c cloudevents.Event) (interface{}, error) {
	handler, err := cm.Get(c.Type())
	if err != nil {
		return nil, err
	}

	cmd, err := handler(ctx, c)
	if err != nil {
		return nil, err
	}

	return cmd, nil
}

func (cm *CeMapper) MapToCloudEvent(e interface{}, eventType, source string) (cloudevents.Event, error) {
	handler, err := cm.GetToCe(eventType)
	if err != nil {
		return cloudevents.Event{}, err
	}

	cloudEvent, err := handler(eventType, source, e)
	if err != nil {
		return cloudevents.Event{}, err
	}

	return cloudEvent, nil
}

func (cm *CeMapper) Get(t string) (CeToEvent, error) {
	if h, ex := cm.CeToEvent[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for {%v} type doesnt exist", t)
}

func (cm *CeMapper) GetToCe(t string) (EventToCe, error) {
	if h, ex := cm.EventToCe[t]; ex {
		return h, nil
	}

	return nil, fmt.Errorf("handler for {%v} type doesnt exist", t)
}

func (c *CeMapper) register(ceType string, f CeToEvent, typeEvent EventType) {
	c.CeToEvent[ceType] = f
	c.handlersTypeEvent[ceType] = typeEvent
}

func (c *CeMapper) RegisterEvent(ceType string, e CeToEvent, ce EventToCe) {
	c.register(ceType, e, EVENT)
	c.EventToCe[ceType] = ce
}

func (c *CeMapper) RegisterCommand(ceType string, e CeToEvent) {
	c.register(ceType, e, COMMAND)
}

func (c *CeMapper) IsCommand(ceType string) bool {
	typeOfEvent, _ := c.GetEventType(ceType)
	return typeOfEvent == COMMAND
}

func (c *CeMapper) IsEvent(ceType string) bool {
	typeOfEvent, _ := c.GetEventType(ceType)
	return typeOfEvent == EVENT
}

func (c *CeMapper) GetEventType(ceType string) (EventType, error) {
	typeEvent, ok := c.handlersTypeEvent[ceType]
	if ok {
		return typeEvent, nil
	}
	return typeEvent, fmt.Errorf("type %s does not exist", ceType)
}

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

func resumeSendedToCloudEvent(eventType, source string, e interface{}) (cloudevents.Event, error) {
	event, ok := e.(events.ResumeSended)
	if !ok {
		return cloudevents.Event{}, fmt.Errorf("impossible cast")
	}

	pbEducations := []*pb.Education{}
	educations := event.Educations.GetEducations()
	for _, v := range educations {
		pbEducations = append(pbEducations, &pb.Education{
			Education: v.GetEducation(),
		})
	}

	pbSkills := []*pb.Skill{}
	skills := event.Skills.GetSkills()
	for _, v := range skills {
		pbSkills = append(pbSkills, &pb.Skill{
			Skill: v.GetSkill(),
		})
	}

	pbDirections := []*pb.Direction{}
	directions := event.Directions.GetDirections()
	for _, v := range directions {
		pbDirections = append(pbDirections, &pb.Direction{
			Direction: v.GetDirection(),
		})
	}

	timestamp := timestamppb.Timestamp{
		Seconds: event.CreatedAt.Unix(),
		Nanos:   int32(event.CreatedAt.Second()),
	}

	pbEvent := pb.ResumeSended{
		ResumeId:      uint64(event.ResumeId.Value),
		FirstName:     event.FirstName.GetFirstName(),
		MiddleName:    event.MiddleName.GetMiddleName(),
		LastName:      event.LastName.GetLastName(),
		PhoneNumber:   event.PhoneNumber.GetPhoneNumber(),
		Educations:    pbEducations,
		AboutMe:       event.AboutMe.GetAboutMe(),
		Skills:        pbSkills,
		Photo:         event.Photo.GetUrl(),
		Directions:    pbDirections,
		AboutProjects: event.AboutProjects.GetAboutProjects(),
		Portfolio:     event.Portfolio.GetPortfolio(),
		StudentGroup:  event.StudentGroup.GetStudentGroup(),
		CreatedAt:     &timestamp,
	}

	cloudEvent := initCloudEvent(eventType, source, &pbEvent)

	return cloudEvent, nil
}

func initCloudEvent(eventType, source string, mes proto.Message) cloudevents.Event {
	cloudEvent := cloudevents.Event{}
	cloudEvent.SetSpecVersion("1.0")
	cloudEvent.SetType(eventType)
	cloudEvent.SetSource(source)
	b, _ := proto.Marshal(mes)
	var protoBytes []byte = make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(protoBytes, b)
	cloudEvent.SetData("application/protobuf", protoBytes)

	return cloudEvent
}
