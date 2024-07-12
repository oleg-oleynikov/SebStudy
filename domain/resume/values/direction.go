package values

import (
	"errors"
	"fmt"
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

func (dr *Direction) ToString() string {
	return fmt.Sprintf("[%s]", dr)
}

type Directions struct {
	directions []Direction
}

func (d *Directions) AppendDirection(dr Direction) {
	d.directions = append(d.directions, dr)
}

func (d *Directions) GetDirections() []Direction {
	return d.directions
}

func (dir *Directions) ToString() string {
	return fmt.Sprintf("[%s]", dir)
}

func (d *Direction) GetDirection() string {
	return d.Direction
}
