package values

import "fmt"

type Educations struct {
	educations []Education
}

func (e *Educations) AppendEducation(ed Education) {
	e.educations = append(e.educations, ed)
}

func (eds *Educations) ToString() string {
	return fmt.Sprintf("[%s]", eds)
}
