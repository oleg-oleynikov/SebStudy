package resume

import (
	"SebStudy/internal/domain/resume/events"
	"SebStudy/internal/domain/resume/values"
	"SebStudy/internal/infrastructure"
	"SebStudy/internal/infrastructure/eventsourcing"
	"time"
)

func RegisterResumeMappingTypes(tm *eventsourcing.TypeMapper) {

	tm.MapEvent(infrastructure.GetValueType(events.ResumeCreated{}), "resumeCreated",
		func(d map[string]interface{}) interface{} {
			createdAt, _ := time.Parse("2006-01-02 15:04:05.999999 -0700 MST", d["createdAt"].(string))
			skills := []values.Skill{}
			for _, s := range d["skills"].([]interface{}) {
				skills = append(skills, values.Skill{Skill: s.(string)})
			}

			birthDate, _ := time.Parse("2006-01-02T15:04:05Z", d["birthDate"].(string))

			return events.ResumeCreated{
				ResumeId:      d["resumeId"].(string),
				Education:     values.Education{Education: d["education"].(string)},
				AboutMe:       values.AboutMe{AboutMe: d["aboutMe"].(string)},
				Skills:        values.Skills{Skills: skills},
				BirthDate:     values.BirthDate{BirthDate: birthDate},
				Direction:     values.Direction{Direction: d["direction"].(string)},
				AboutProjects: values.AboutProjects{AboutProjects: d["aboutProjects"].(string)},
				Portfolio:     values.Portfolio{Portfolio: d["portfolio"].(string)},
				CreatedAt:     createdAt,
			}
		},
		func(v interface{}) (string, map[string]interface{}) {
			t := v.(events.ResumeCreated)
			skills := []string{}
			for _, s := range t.Skills.GetSkills() {
				skills = append(skills, s.Skill)
			}
			return "resumeCreated",
				map[string]interface{}{
					"resumeId":      t.ResumeId,
					"education":     t.Education.GetEducation(),
					"aboutMe":       t.AboutMe.GetAboutMe(),
					"skills":        skills,
					"birthDate":     t.BirthDate.GetBirthDate(),
					"direction":     t.Direction.GetDirection(),
					"aboutProjects": t.AboutProjects.GetAboutProjects(),
					"portfolio":     t.Portfolio.GetPortfolio(),
					"createdAt":     t.CreatedAt.Format("2006-01-02 15:04:05.999999 -0700 MST"),
				}
		})

	tm.MapEvent(infrastructure.GetValueType(events.ResumeChanged{}), "resumeChanged",
		func(d map[string]interface{}) interface{} {
			return nil
		},
		func(v interface{}) (string, map[string]interface{}) {
			return "", nil
		})
}
