package values

import "fmt"

type Photo struct {
	photo string
}

func NewPhoto(photo string) (*Photo, error) {
	return &Photo{
		photo: photo,
	}, nil
}

func (photo *Photo) ToString() string {
	return fmt.Sprintf("%s", photo)
}
