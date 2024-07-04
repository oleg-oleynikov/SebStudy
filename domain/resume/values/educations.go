package values

import "fmt"

type Educations struct {
	educations []Education
}

func (e *Educations) AppendEducations(ed ...Education) {
	e.educations = append(e.educations, ed...)
}

func (e *Educations) ToString() string {
	return fmt.Sprintf("[%s]", e)
}

func (e *Educations) GetEducations() []Education {
	return e.educations
}
