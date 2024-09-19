package values

type Education struct {
	Education string
}

func NewEducation(education string) (*Education, error) {
	return &Education{
		Education: education,
	}, nil
}

func (ed *Education) String() string {
	return ed.Education
}

func (ed *Education) GetEducation() string {
	return ed.Education
}
