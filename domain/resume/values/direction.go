package values

import "errors"

type Direction struct {
	direction string
}

func NewDirection(direction string) (*Direction, error) {
	if len(direction) > 50 {
		return nil, errors.New("too much symbols (max: 50)")
	}

	return &Direction{
		direction: direction,
	}, nil
}

type Directions struct {
	directions []Direction
}

func (d *Directions) Appendskill(dr Direction) {
	d.directions = append(d.directions, dr)
}
