package values

import "fmt"

type Skills struct {
	Skills []Skill
}

func (s *Skills) AppendSkills(sk ...Skill) error {
	if len(s.Skills)+len(sk) < 500 {
		s.Skills = append(s.Skills, sk...)
		return nil
	}

	return fmt.Errorf("the number of skills cannot be more than 500")
}

func (s *Skills) GetSkills() []Skill {
	return s.Skills
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
