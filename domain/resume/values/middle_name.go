package values

import (
	"errors"
	"regexp"
)

type MiddleName struct {
	middleName string
}

func NewMiddleName(middleName string) (*MiddleName, error) {
	if !isValidMiddleName(middleName) {
		return nil, errors.New("invalid middle name")
	}

	return &MiddleName{
		middleName: middleName,
	}, nil
}

func isValidMiddleName(middleName string) bool {
	middleNameRegex := regexp.MustCompile(`^[^0-9A-Za-z]+$`)

	return middleNameRegex.MatchString(middleName)
}
