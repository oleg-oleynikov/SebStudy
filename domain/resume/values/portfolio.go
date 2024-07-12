package values

import (
	"errors"
	"fmt"
	"net/url"
)

type Portfolio struct {
	Portfolio string
}

func NewPortfolio(portfolio string) (*Portfolio, error) {
	if !isValidPortfolio(portfolio) {
		return nil, errors.New("incorrect link")
	}

	return &Portfolio{
		Portfolio: portfolio,
	}, nil
}

func isValidPortfolio(portfolio string) bool {
	u, _ := url.ParseRequestURI(portfolio)
	return u != nil
}

func (p Portfolio) ToString() string {
	return fmt.Sprintf("%s", p)
}

func (p *Portfolio) GetPortfolio() string {
	return p.Portfolio
}
