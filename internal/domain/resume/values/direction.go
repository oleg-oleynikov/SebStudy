package values

import (
	"errors"
	"fmt"
)

type Direction struct {
	Direction string
}

func NewDirection(direction string) (*Direction, error) {
	lengthDirection := len(direction)
	if lengthDirection == 0 {
		return nil, fmt.Errorf("direction doesnt be empty")
	}
	if lengthDirection > 50 {
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
