package ce_mapper_func

import (
	"SebStudy/adapters/util"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	pb "SebStudy/proto/resume"
	"context"
	"encoding/base64"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var toResumeSendedEvent util.CeToEvent = func(ctx context.Context, c cloudevents.Event) (interface{}, error) {
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

var toCeResumeSended util.EventToCe = func(eventType, source string, e interface{}) (cloudevents.Event, error) {
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

func init() {
	ceMapper := util.GetCeMapperInstance()
	ceMapper.RegisterEvent("resume.sended", toResumeSendedEvent, toCeResumeSended)
}
