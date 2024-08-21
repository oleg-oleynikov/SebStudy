package mapping

import (
	"SebStudy/adapters/util"
	"SebStudy/domain/resume/events"
	"SebStudy/domain/resume/values"
	"SebStudy/infrastructure"
	pb "SebStudy/proto/resume"

	"context"
	"fmt"

	"google.golang.org/protobuf/types/known/timestamppb"
	v1 "open-cluster-management.io/sdk-go/pkg/cloudevents/generic/options/grpc/protobuf/v1"
)

var toResumeCreated util.CloudeventToEvent = func(ctx context.Context, c *v1.CloudEvent) (interface{}, error) {
	rs := &pb.ResumeCreated{}

	if err := infrastructure.DecodeCloudeventData(c, rs); err != nil {
		return nil, err
	}

	resumeID := values.NewResumeId(rs.GetResumeId())

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

	education, err := values.NewEducation(rs.Education)
	if err != nil {
		return nil, err
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

	photo, err := values.NewPhoto(rs.GetPhoto(), "")
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

	createdResume := events.NewResumeCreated(
		resumeID, *firstName, *middleName, *lastName, *phoneNumber,
		education, *aboutMe, skills, *photo, directions,
		*aboutProjects, *portfolio, *studentGroup, rs.CreatedAt.AsTime())

	return createdResume, nil
}

var toCloudeventResumeCreated util.EventToCloudevent = func(cloudeventType, source string, e interface{}) (*v1.CloudEvent, error) {
	event, ok := e.(events.ResumeCreated)
	if !ok {
		return &v1.CloudEvent{}, fmt.Errorf("impossible cast")
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

	pbEvent := pb.ResumeCreated{
		ResumeId:      event.ResumeId.Value,
		FirstName:     event.FirstName.GetFirstName(),
		MiddleName:    event.MiddleName.GetMiddleName(),
		LastName:      event.LastName.GetLastName(),
		PhoneNumber:   event.PhoneNumber.GetPhoneNumber(),
		Education:     event.Education.GetEducation(),
		AboutMe:       event.AboutMe.GetAboutMe(),
		Skills:        pbSkills,
		Photo:         event.Photo.GetPhoto(),
		Directions:    pbDirections,
		AboutProjects: event.AboutProjects.GetAboutProjects(),
		Portfolio:     event.Portfolio.GetPortfolio(),
		StudentGroup:  event.StudentGroup.GetStudentGroup(),
		CreatedAt:     &timestamp,
	}

	cloudEvent, err := util.InitCloudEvent(cloudeventType, source, &pbEvent)
	if err != nil {
		return nil, err
	}

	return cloudEvent, nil
}
