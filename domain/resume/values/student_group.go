package values

import (
	"errors"
	"fmt"
)

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
	return len(studentGroup) <= 15
}

func (group *StudentGroup) ToString() string {
	return fmt.Sprintf("%s", group)
}
