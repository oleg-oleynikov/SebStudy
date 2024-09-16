package mapping

import (
	"SebStudy/internal/domain/resume/commands"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
	"SebStudy/pb"
	"context"
)

func toCreateResume(ctx context.Context, c *pb.CloudEvent) (interface{}, error) {

	rs := pb.ResumeCreated{}

	if err := infrastructure.DecodeCloudeventData(c, &rs); err != nil {
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

	direction, err := values.NewDirection(rs.GetDirection())
	if err != nil {
		return nil, err
	}

	aboutProjects, err := values.NewAboutProjects(rs.GetAboutProjects())
	if err != nil {
		return nil, err
	}

	portfolio, err := values.NewPortfolio(rs.GetPortfolio())
	if err != nil {
		return nil, err
	}

	createdResume := commands.NewCreateResume(
		education, *aboutMe, skills, *direction,
		*aboutProjects, *portfolio)

	return createdResume, nil
}
