package values

type Photo struct {
	url string
}

func NewPhoto(url string) (*Photo, error) {
	return &Photo{
		url: url,
	}, nil
}
