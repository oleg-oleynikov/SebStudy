package values

import (
	"errors"
)

type Direction struct {
	Direction string
}

func NewDirection(direction string) (*Direction, error) {
	if len(direction) > 50 {
		return nil, errors.New("too much symbols (max: 50)")
	}

	return &Direction{
		Direction: direction,
	}, nil
}

func (dr *Direction) String() string {
	return dr.Direction
}

func (d *Direction) GetDirection() string {
	return d.Direction
}
