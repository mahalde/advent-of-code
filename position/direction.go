package position

import (
	"github.com/mahalde/advent-of-code/tuples"
)

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

var (
	AllDirections      = []Direction{North, East, South, West}
	AllDirectionTuples = []*tuples.Tuple{
		tuples.NewTuple(0, -1), // North
		tuples.NewTuple(1, 0),  // East
		tuples.NewTuple(0, 1),  // South
		tuples.NewTuple(-1, 0), // West
	}
)

func (d Direction) TurnLeft() Direction {
	switch d {
	case North:
		return West
	case East:
		return North
	case South:
		return East
	case West:
		return South
	}

	return d
}

func (d Direction) TurnRight() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}

	return d
}

func (d Direction) TurnAround() Direction {
	switch d {
	case North:
		return South
	case East:
		return West
	case South:
		return North
	case West:
		return East
	}

	return d
}

func (d Direction) String() string {
	return string(d)
}
