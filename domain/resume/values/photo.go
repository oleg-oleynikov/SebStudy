package values

import "fmt"

type Photo struct {
	url string
}

func NewPhoto(url string) (*Photo, error) {
	return &Photo{
		url: url,
	}, nil
}

func (photo *Photo) ToString() string {
	return fmt.Sprintf("%s", photo)
}

func (p *Photo) GetUrl() string {
	return p.url
}
