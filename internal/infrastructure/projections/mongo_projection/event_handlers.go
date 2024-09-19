package mongo_projection

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure"
	"context"
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

	return o.mongoRepo.Insert(ctx, rp)
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

	return o.mongoRepo.Update(ctx, rp)
}
