package values

import "time"

type BirthDate struct {
	BirthDate time.Time
}

func NewBirthDate(birthDate time.Time) (*BirthDate, error) {
	return &BirthDate{
		BirthDate: birthDate,
	}, nil
}

func (bd *BirthDate) GetBirthDate() time.Time {
	return bd.BirthDate
}
