package resume

import (
	"resume-server/internal/domain/resume/events"
	"resume-server/internal/domain/resume/values"
	"resume-server/internal/infrastructure"
	"resume-server/internal/infrastructure/eventsourcing"
	"time"
)

func RegisterResumeMappingTypes(tm *eventsourcing.TypeMapper) {

	tm.MapEvent(infrastructure.GetValueType(events.ResumeCreated{}), "ResumeCreated",
		func(d map[string]interface{}) interface{} {
			createdAt, _ := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", d["createdAt"].(string))
			skills := []values.Skill{}
			for _, s := range d["skills"].([]interface{}) {
				skills = append(skills, values.Skill{Skill: s.(string)})
			}

			return events.ResumeCreated{
				ResumeId:      d["resumeId"].(string),
				AboutMe:       values.AboutMe{AboutMe: d["aboutMe"].(string)},
				Skills:        values.Skills{Skills: skills},
				Direction:     values.Direction{Direction: d["direction"].(string)},
				AboutProjects: values.AboutProjects{AboutProjects: d["aboutProjects"].(string)},
				Portfolio:     values.Portfolio{Portfolio: d["portfolio"].(string)},
				CreatedAt:     createdAt,
			}
		},
		func(v interface{}) (string, map[string]interface{}) {
			event := v.(events.ResumeCreated)
			skills := []string{}
			for _, s := range event.Skills.GetSkills() {
				skills = append(skills, s.Skill)
			}

			return "ResumeCreated",
				map[string]interface{}{
					"resumeId":      event.ResumeId,
					"aboutMe":       event.AboutMe.GetAboutMe(),
					"skills":        skills,
					"direction":     event.Direction.GetDirection(),
					"aboutProjects": event.AboutProjects.GetAboutProjects(),
					"portfolio":     event.Portfolio.GetPortfolio(),
					"createdAt":     event.CreatedAt.Format("2006-01-02 15:04:05.999999 -0700 MST"),
				}
		})

	tm.MapEvent(infrastructure.GetValueType(events.ResumeChanged{}), "ResumeChanged",
		func(d map[string]interface{}) interface{} {
			createdAt, _ := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", d["createdAt"].(string))

			skills := []values.Skill{}
			for _, s := range d["skills"].([]interface{}) {
				skills = append(skills, values.Skill{Skill: s.(string)})
			}

			return events.ResumeChanged{
				ResumeId:      d["resumeId"].(string),
				AboutMe:       values.AboutMe{AboutMe: d["aboutMe"].(string)},
				Skills:        values.Skills{Skills: skills},
				Direction:     values.Direction{Direction: d["direction"].(string)},
				AboutProjects: values.AboutProjects{AboutProjects: d["aboutProjects"].(string)},
				Portfolio:     values.Portfolio{Portfolio: d["portfolio"].(string)},
				CreatedAt:     createdAt,
			}
		},
		func(v interface{}) (string, map[string]interface{}) {
			event := v.(events.ResumeChanged)
			skills := []string{}
			for _, s := range event.Skills.GetSkills() {
				skills = append(skills, s.Skill)
			}

			return "ResumeChanged",
				map[string]interface{}{
					"resumeId":      event.ResumeId,
					"aboutMe":       event.AboutMe.GetAboutMe(),
					"skills":        skills,
					"direction":     event.Direction.GetDirection(),
					"aboutProjects": event.AboutProjects.GetAboutProjects(),
					"portfolio":     event.Portfolio.GetPortfolio(),
					"createdAt":     event.CreatedAt.Format("2006-01-02 15:04:05.999999 -0700 MST"),
				}
		})
}
