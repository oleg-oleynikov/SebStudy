package mongo_projection

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/models"
	"SebStudy/internal/infrastructure"
	"context"
	"fmt"
)

func (o *mongoProjection) onResumeCreate(ctx context.Context, e interface{}, md *infrastructure.EventMetadata) error {
	event, ok := e.(events.ResumeCreated)
	if !ok {
		o.log.Debugf("mongoProjection.onResumeCreate: Failed to cast")
		return fmt.Errorf("failed to cast")
	}
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

// func (o *mongoProjection) onResumeChanged(ctx context.Context, e interface{}) error {
// 	o.log.Debugf("mongoProjection.ResumeChanged: Not impl")
// 	return fmt.Errorf("not impl")
// }
