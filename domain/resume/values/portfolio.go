package values

import (
	"errors"
	"net/url"
)

type Portfolio struct {
	portfolio string
}

func NewPortfolio(portfolio string) (*Portfolio, error) {
	if !isValidPortfolio(portfolio) {
		return nil, errors.New("incorrect link")
	}

	return &Portfolio{
		portfolio: portfolio,
	}, nil
}

func isValidPortfolio(portfolio string) bool {
	u, _ := url.ParseRequestURI(portfolio)
	return u != nil
}
