package values

import (
	"errors"
	"regexp"
)

type FirstName struct {
	firstName string
}

func NewFirstName(firstName string) (*FirstName, error) {
	if !isValidFirstName(firstName) {
		return nil, errors.New("invalid first name")
	}

	return &FirstName{
		firstName: firstName,
	}, nil
}

func isValidFirstName(firstName string) bool {
	firstNameRegex := regexp.MustCompile(`^[^0-9A-Za-z]+$`)

	return firstNameRegex.MatchString(firstName)
}