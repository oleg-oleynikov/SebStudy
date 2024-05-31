package values

import "errors"

type Skill struct {
	skill string
}

func NewSkill(skill string) (*Skill, error) {
	if len(skill) > 30 {
		return nil, errors.New("too much symbols (max: 30)")
	}

	return &Skill{
		skill: skill,
	}, nil
}

type Skills struct {
	skills []Skill
}

func (s *Skills) Appendskill(sk Skill) {
	s.skills = append(s.skills, sk)
}
