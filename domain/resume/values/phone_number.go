package values

import (
	"errors"
	"regexp"
)

type PhoneNumber struct {
	phoneNumber string
}

func NewPhoneNumber(phoneNumber string) (*PhoneNumber, error) {
	if !isValidPhoneNumber(phoneNumber) {
		return nil, errors.New("invalid phone number")
	}

	return &PhoneNumber{
		phoneNumber: phoneNumber,
	}, nil
}

func isValidPhoneNumber(phoneNumber string) bool {
	phoneNumberRegex := regexp.MustCompile(`^\d{11}$`)

	return phoneNumberRegex.MatchString(phoneNumber)
}
