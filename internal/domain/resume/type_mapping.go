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
			createdAt, _ := time.Parse("2006-01-02 15:04:05", d["createdAt"].(string))
			skills := []values.Skill{}
			for _, s := range d["skills"].([]interface{}) {
				skillData := s.(map[string]interface{})
				skill := values.Skill{Skill: skillData["Skill"].(string)}
				skills = append(skills, values.Skill(skill))
			}

			return events.ResumeCreated{
				ResumeId:      d["resumeId"].(string),
				FirstName:     values.FirstName{FirstName: d["firstName"].(string)},
				MiddleName:    values.MiddleName{MiddleName: d["middleName"].(string)},
				LastName:      values.LastName{LastName: d["lastName"].(string)},
				PhoneNumber:   values.PhoneNumber{PhoneNumber: d["phoneNumber"].(string)},
				Education:     values.Education{Education: d["education"].(string)},
				AboutMe:       values.AboutMe{AboutMe: d["aboutMe"].(string)},
				Skills:        values.Skills{Skills: skills},
				Photo:         values.Photo{Url: d["photo"].(string)},
				Direction:     values.Direction{Direction: d["direction"].(string)},
				AboutProjects: values.AboutProjects{AboutProjects: d["aboutProjects"].(string)},
				Portfolio:     values.Portfolio{Portfolio: d["portfolio"].(string)},
				CreatedAt:     createdAt,
			}
		},
		func(v interface{}) (string, map[string]interface{}) {
			t := v.(events.ResumeCreated)
			return "resumeCreated",
				map[string]interface{}{
					"resumeId":      t.ResumeId,
					"firstName":     t.FirstName.GetFirstName(),
					"middleName":    t.MiddleName.GetMiddleName(),
					"lastName":      t.LastName.GetLastName(),
					"phoneNumber":   t.PhoneNumber.GetPhoneNumber(),
					"education":     t.Education.GetEducation(),
					"aboutMe":       t.AboutMe.GetAboutMe(),
					"skills":        t.Skills.GetSkills(),
					"photo":         t.Photo.GetUrl(),
					"direction":     t.Direction.GetDirection(),
					"aboutProjects": t.AboutProjects.GetAboutProjects(),
					"portfolio":     t.Portfolio.GetPortfolio(),
					"createdAt":     t.CreatedAt.String(),
				}
		})
}
