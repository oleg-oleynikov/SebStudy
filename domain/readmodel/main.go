package main

import (
	"SebStudy/domain/resume/values"
	"fmt"
)

func main() {
	f, err := values.NewFirstName("Апельсин")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Printf("%v\n", f)
}
