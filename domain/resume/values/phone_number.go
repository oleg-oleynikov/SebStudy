package values

import (
	"fmt"
	"regexp"
)

type PhoneNumber struct {
	phoneNumber string
}

func NewPhoneNumber(phoneNumber string) (*PhoneNumber, error) {
	if !isValidPhoneNumber(phoneNumber) {
		return nil, fmt.Errorf("invalid phone number: {%s}", phoneNumber)
	}

	return &PhoneNumber{
		phoneNumber: phoneNumber,
	}, nil
}

func isValidPhoneNumber(phoneNumber string) bool {
	phoneNumberRegex := regexp.MustCompile(`^\d{11}$`)

	return phoneNumberRegex.MatchString(phoneNumber)
}

func (phoneNum *PhoneNumber) ToString() string {
	return fmt.Sprintf("%s", phoneNum)
}

func (p *PhoneNumber) GetPhoneNumber() string {
	return p.phoneNumber
}
