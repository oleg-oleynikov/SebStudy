package values

type Photo struct {
	photo string
}

func NewPhoto(photo string) (*Photo, error) {
	return &Photo{
		photo: photo,
	}, nil
}
