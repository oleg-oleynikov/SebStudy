package mongo_projection

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure"
	"context"
	"fmt"
)

func (o *mongoProjection) onResumeCreate(ctx context.Context, e interface{}, _ *infrastructure.EventMetadata) error {
	event, ok := e.(events.ResumeCreated)
	if !ok {
		o.log.Debugf("mongoProjection.onResumeCreate: Failed to cast")
		return fmt.Errorf("failed to cast")
	}
	skills := []string{}
	for _, s := range event.Skills.GetSkills() {
		skills = append(skills, s.GetSkill())
	}

	o.log.Debugf("Добавить BornDate")
	rp := &models.ResumeProjection{
		// Education:     event.Education.GetEducation(),
		AboutMe:       event.AboutMe.GetAboutMe(),
		BornDate:      1, // Добавить bornDate
		Skills:        skills,
		Direction:     event.Direction.GetDirection(),
		AboutProjects: event.AboutProjects.GetAboutProjects(),
		Portfolio:     event.Portfolio.GetPortfolio(),
	}

	return o.mongoRepo.Insert(ctx, rp)
}

func (o *mongoProjection) onResumeChanged(ctx context.Context, e interface{}) error {
	o.log.Debugf("mongoProjection.ResumeChanged: Not impl")
	return fmt.Errorf("not impl")
}
