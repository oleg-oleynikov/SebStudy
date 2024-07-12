package values

import "fmt"

type Photo struct {
	Url   string
	Photo []byte
}

func NewPhoto(photo []byte, url string) (*Photo, error) {
	return &Photo{
		Photo: photo,
	}, nil
}

func (photo *Photo) ToString() string {
	return fmt.Sprintf("%s", photo)
}

func (p *Photo) SetUrl(url string) {
	p.Url = url
}

func (p *Photo) GetUrl() string {
	return p.Url
}

func (p *Photo) GetPhoto() []byte {
	return p.Photo
}
