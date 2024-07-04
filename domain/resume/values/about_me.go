package values

import (
	"errors"
	"fmt"
)

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

func (aMe *AboutMe) ToString() string {
	return fmt.Sprintf("%s", aMe)
}

func (a *AboutMe) GetAboutMe() string {
	return a.aboutMe
}
