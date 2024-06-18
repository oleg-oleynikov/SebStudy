package values

import (
	"errors"
	"fmt"
)

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

func (d *Directions) AppendDirection(dr Direction) {
	d.directions = append(d.directions, dr)
}

func (d *Directions) GetDirections() []Direction {
	return d.directions
}

func (dir *Directions) ToString() string {
	return fmt.Sprintf("[%s]", dir)
}
