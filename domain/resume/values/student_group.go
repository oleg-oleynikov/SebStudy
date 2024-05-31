package values

import "errors"

type StudentGroup struct {
	studentGroup string
}

func NewStudentGroup(studentGroup string) (*StudentGroup, error) {
	if !isValidStudentGroup(studentGroup) {
		return nil, errors.New("too much symbols (max value: 15)")
	}

	return &StudentGroup{
		studentGroup: studentGroup,
	}, nil
}

func isValidStudentGroup(studentGroup string) bool {
	if len(studentGroup) > 15 {
		return false
	}
	return true
}
