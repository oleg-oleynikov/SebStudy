package values

import (
	"errors"
	"fmt"
)

type AboutProjects struct {
	AboutProjects string
}

func NewAboutProjects(aboutProjects string) (*AboutProjects, error) {
	if len(aboutProjects) > 400 {
		return nil, errors.New("too much symbols (max: 400)")
	}

	return &AboutProjects{
		AboutProjects: aboutProjects,
	}, nil
}

func (aProjects *AboutProjects) ToString() string {
	return fmt.Sprintf("%s", aProjects)
}

func (a *AboutProjects) GetAboutProjects() string {
	return a.AboutProjects
}
