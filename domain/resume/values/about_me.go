package values

import "errors"

type AboutMe struct {
	aboutMe string
}

func NewAboutMe(aboutMe string) (*AboutMe, error) {
	if len(aboutMe) > 350 {
		return nil, errors.New("too much symbols (max: 350)")
	}

	return &AboutMe{
		aboutMe: aboutMe,
	}, nil
}
