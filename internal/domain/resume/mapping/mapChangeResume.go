package mapping

import (
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
	"SebStudy/pb"
	"context"
	"time"
)

func toChangeResume(ctx context.Context, c *pb.CloudEvent) (interface{}, error) {
	changeResumeProto := pb.ResumeChanged{}

	if err := infrastructure.DecodeCloudeventData(c, &changeResumeProto); err != nil {
		return nil, err
	}

	resumeId := changeResumeProto.GetResumeId()

	education, err := values.NewEducation(changeResumeProto.Education)
	if err != nil {
		return nil, err
	}

	aboutMe, err := values.NewAboutMe(changeResumeProto.GetAboutMe())
	if err != nil {
		return nil, err
	}

	var skills values.Skills
	for i := 0; i < len(changeResumeProto.Skills); i++ {
		data := changeResumeProto.Skills[i]
		skill, err := values.NewSkill(data.Skill)
		if err != nil {
			return nil, err
		}
		skills.AppendSkills(*skill)
	}

	var timeBirth time.Time
	birthDateProto := changeResumeProto.GetBirthDate()
	if birthDateProto != "" {
		timeBirth, err = time.Parse("2006-01-02", changeResumeProto.GetBirthDate())
		if err != nil {
			return nil, err
		}
	}

	birthDate, err := values.NewBirthDate(timeBirth)
	if err != nil {
		return nil, err
	}

	direction, err := values.NewDirection(changeResumeProto.GetDirection())
	if err != nil {
		return nil, err
	}

	aboutProjects, err := values.NewAboutProjects(changeResumeProto.GetAboutProjects())
	if err != nil {
		return nil, err
	}

	portfolio, err := values.NewPortfolio(changeResumeProto.GetPortfolio())
	if err != nil {
		return nil, err
	}

	changeResume := commands.NewChangeResume(
		resumeId, education, *aboutMe, skills, *birthDate,
		*direction, *aboutProjects, *portfolio)

	return changeResume, nil
}
