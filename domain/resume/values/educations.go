package values

type Educations struct {
	educations []Education
}

func (e *Educations) AppendEducation(ed Education) {
	e.educations = append(e.educations, ed)
}
