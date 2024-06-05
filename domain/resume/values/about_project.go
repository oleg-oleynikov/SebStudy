package values

import (
	"errors"
	"fmt"
)

type AboutProjects struct {
	aboutProjects string
}

func NewAboutProjects(aboutProjects string) (*AboutProjects, error) {
	if len(aboutProjects) > 400 {
		return nil, errors.New("too much symbols (max: 400)")
	}

	return &AboutProjects{
		aboutProjects: aboutProjects,
	}, nil
}

func (aProjects *AboutProjects) ToString() string {
	return fmt.Sprintf("%s", aProjects)
}
