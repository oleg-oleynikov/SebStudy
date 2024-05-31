package main

import (
	"SebStudy/domain/resume/values"
	"fmt"
)

func main() {
	portfolio, r := values.NewPortfolio("https://github.com")

	fmt.Println(portfolio, r)
}
