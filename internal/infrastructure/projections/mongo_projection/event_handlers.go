package mongo_projection

import (
	"context"
	"resume-server/internal/domain/resume/events"
	"resume-server/internal/domain/resume/models"
	"resume-server/internal/infrastructure"
)

func (o *mongoProjection) onResumeCreate(ctx context.Context, event events.ResumeCreated, md *infrastructure.EventMetadata) error {
	skills := []string{}
	for _, s := range event.Skills.GetSkills() {
		skills = append(skills, s.GetSkill())
	}

	rp := &models.ResumeProjection{
		Id:            event.ResumeId,
		Education:     event.Education.GetEducation(),
		AboutMe:       event.AboutMe.GetAboutMe(),
		Skills:        skills,
		BirthDate:     event.BirthDate.GetBirthDate(),
		Direction:     event.Direction.GetDirection(),
		AboutProjects: event.AboutProjects.GetAboutProjects(),
		Portfolio:     event.Portfolio.GetPortfolio(),
		UserId:        md.UserId,
	}

	return o.resumeRepo.Insert(ctx, rp)
}

func (o *mongoProjection) onResumeChanged(ctx context.Context, event events.ResumeChanged, md *infrastructure.EventMetadata) error {
	skills := []string{}
	for _, s := range event.Skills.GetSkills() {
		skills = append(skills, s.GetSkill())
	}

	rp := &models.ResumeProjection{
		Id:            event.ResumeId,
		Education:     event.Education.GetEducation(),
		AboutMe:       event.AboutMe.GetAboutMe(),
		Skills:        skills,
		BirthDate:     event.BirthDate.GetBirthDate(),
		Direction:     event.Direction.GetDirection(),
		AboutProjects: event.AboutProjects.GetAboutProjects(),
		Portfolio:     event.Portfolio.GetPortfolio(),
		UserId:        md.UserId,
	}

	return o.resumeRepo.Update(ctx, rp)
}
