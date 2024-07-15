package ce_resume_mapper

import (
	"SebStudy/adapters/util"
	"SebStudy/domain/resume/commands"
	"SebStudy/domain/resume/values"
	"context"

	pb "SebStudy/proto/resume"

	cloudevents "github.com/cloudevents/sdk-go/v2"
)

var toSendResumeCommand util.CeToEvent = func(ctx context.Context, c cloudevents.Event) (interface{}, error) {

	var rs pb.ResumeSended
	// bytes, err := base64.StdEncoding.DecodeString(string(c.DataEncoded))
	// if err != nil {
	// 	return nil, err
	// }

	// if err := proto.Unmarshal(c.Data(), &rs); err != nil {
	// 	return nil, err
	// }

	if err := c.DataAs(&rs); err != nil {
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

	// var educations values.Educations
	// for i := 0; i < len(rs.Educations); i++ {
	// 	data := rs.Educations[i]
	// 	education, err := values.NewEducation(data.Education)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	educations.AppendEducations(*education)
	// }

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

	createdResume := commands.NewSendResume(
		resumeID, *firstName, *middleName, *lastName, *phoneNumber,
		education, *aboutMe, skills, *photo, directions,
		*aboutProjects, *portfolio, *studentGroup)

	return createdResume, nil
}

func init() {
	ceMapper := util.GetCeMapperInstance()
	ceMapper.RegisterCommand("resume.send", toSendResumeCommand)
}