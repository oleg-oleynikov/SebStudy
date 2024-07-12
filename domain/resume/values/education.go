package values

import "fmt"

type Education struct {
	Education string
}

func NewEducation(education string) (Education, error) {
	// if !isValidEducation(education) {
	// 	return nil, errors.New("invalid education")
	// }

	return Education{
		Education: education,
	}, nil
}

// func isValidEducation(education string) bool {
// 	return true
// }

func (ed *Education) ToString() string {
	return fmt.Sprintf("%s", ed)
}

func (ed *Education) GetEducation() string {
	return ed.Education
}
