package values

import "fmt"

type Skills struct {
	skills []Skill
}

func (s *Skills) AppendSkill(sk Skill) {
	s.skills = append(s.skills, sk)
}

func (s *Skills) GetSkills() []Skill {
	return s.skills
}

func (sks *Skills) ToString() string {
	return fmt.Sprintf("[%s]", sks)
}
