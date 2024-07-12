package values

import (
	"errors"
	"fmt"
)

type Skill struct {
	Skill string
}

func NewSkill(skill string) (*Skill, error) {
	if len(skill) > 30 {
		return nil, errors.New("too much symbols (max: 30)")
	}

	return &Skill{
		Skill: skill,
	}, nil
}

func (sk *Skill) ToString() string {
	return fmt.Sprintf("%s", sk)
}

func (s *Skill) GetSkill() string {
	return s.Skill
}
