package values

import (
	"fmt"
	"regexp"
)

type FirstName struct {
	firstName string
}

func NewFirstName(firstName string) (*FirstName, error) {
	if !isValidFirstName(firstName) {
		return nil, fmt.Errorf("invalid first name: %s", firstName)
	}

	return &FirstName{
		firstName: firstName,
	}, nil
}

func isValidFirstName(firstName string) bool {
	firstNameRegex := regexp.MustCompile(`^[^0-9A-Za-z]+$`)

	return firstNameRegex.MatchString(firstName)
}

func (frName *FirstName) ToString() string {
	return fmt.Sprintf("%s", frName)
}
