package values

import "fmt"

type Skills struct {
	skills []Skill
}

func (s *Skills) AppendSkills(sk ...Skill) error {
	if len(s.skills)+len(sk) < 500 {
		s.skills = append(s.skills, sk...)
		return nil
	}

	return fmt.Errorf("The number of skills cannot be more than 500")
}

func (s *Skills) GetSkills() []Skill {
	return s.skills
}

func (s *Skills) ToString() string {
	return fmt.Sprintf("[%s]", s)
}

func NewSkills(sk ...Skill) (*Skills, error) {
	s := &Skills{}
	if err := s.AppendSkills(sk...); err != nil {
		return nil, err
	}

	return s, nil
}
