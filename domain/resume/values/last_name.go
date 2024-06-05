package values

import (
	"errors"
	"fmt"
	"regexp"
)

type LastName struct {
	lastName string
}

func NewLastName(lastName string) (*LastName, error) {
	if !isValidLastName(lastName) {
		return nil, errors.New("invalid last name")
	}

	return &LastName{
		lastName: lastName,
	}, nil
}

func isValidLastName(lastName string) bool {
	lastNameRegex := regexp.MustCompile(`^[^0-9A-Za-z]+$`)

	return lastNameRegex.MatchString(lastName)
}

func (lName *LastName) ToString() string {
	return fmt.Sprintf("%s", lName)
}
