package values

type Education struct {
	education string
}

func NewEducation(education string) (*Education, error) {
	// if !isValidEducation(education) {
	// 	return nil, errors.New("invalid education")
	// }

	return &Education{
		education: education,
	}, nil
}

// func isValidEducation(education string) bool {
// 	return true
// }
