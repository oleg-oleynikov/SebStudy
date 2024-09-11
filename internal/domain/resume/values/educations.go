package values

import "fmt"

type Educations struct {
	Educations []Education
}

func (e *Educations) AppendEducations(ed ...Education) {
	e.Educations = append(e.Educations, ed...)
}

func (e *Educations) ToString() string {
	return fmt.Sprintf("[%s]", e)
}

func (e *Educations) GetEducations() []Education {
	return e.Educations
}
