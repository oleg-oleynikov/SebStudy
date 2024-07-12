package values

import (
	"errors"
	"fmt"
)

type StudentGroup struct {
	StudentGroup string
}

func NewStudentGroup(studentGroup string) (*StudentGroup, error) {
	if !isValidStudentGroup(studentGroup) {
		return nil, errors.New("too much symbols (max value: 15)")
	}

	return &StudentGroup{
		StudentGroup: studentGroup,
	}, nil
}

func isValidStudentGroup(studentGroup string) bool {
	return len(studentGroup) <= 15
}

func (group *StudentGroup) ToString() string {
	return fmt.Sprintf("%s", group)
}

func (sg *StudentGroup) GetStudentGroup() string {
	return sg.StudentGroup
}
