package values

import "errors"

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
