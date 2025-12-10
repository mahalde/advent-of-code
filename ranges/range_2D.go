package ranges

import (
	"github.com/mahalde/advent-of-code/math"
	"github.com/mahalde/advent-of-code/tuples"
)

type Range2D struct {
	Start *tuples.Tuple
	End   *tuples.Tuple
}

type Extrema2D struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

// NewRange2D creates a new 2D range.
// It supports horizontal, vertical and diagonal ranges.
func NewRange2D(xStart, xEnd, yStart, yEnd int) (r *Range2D, ok bool) {
	xDelta := math.Abs(xEnd - xStart)
	yDelta := math.Abs(yEnd - yStart)
	if xDelta != yDelta && xDelta > 0 && yDelta > 0 {
		return nil, false
	}

	return &Range2D{
		Start: tuples.NewTuple(xStart, yStart),
		End:   tuples.NewTuple(xEnd, yEnd),
	}, true
}

// Equals checks if two Range2D are equal, regardless of the order of start and end.
func (r *Range2D) Equals(other *Range2D) bool {
	forwardCheck := r.Start.Equals(other.Start) && r.End.Equals(other.End)
	reverseCheck := r.Start.Equals(other.End) && r.End.Equals(other.Start)

	return forwardCheck || reverseCheck
}

// Delta returns the delta of the range in both dimensions.
func (r *Range2D) Delta() *tuples.Tuple {
	return tuples.NewTuple(math.Delta(r.End.X-r.Start.X), math.Delta(r.End.Y-r.Start.Y))
}

func (r *Range2D) Gradient() float64 {
	if r.Start.X == r.End.X {
		return 0
	}

	return float64(r.End.Y-r.Start.Y) / float64(r.End.X-r.Start.X)
}

func (r *Range2D) Extrema() *Extrema2D {
	return &Extrema2D{
		MinX: min(r.Start.X, r.End.X),
		MaxX: max(r.Start.X, r.End.X),
		MinY: min(r.Start.Y, r.End.Y),
		MaxY: max(r.Start.Y, r.End.Y),
	}
}

// IsHorizontal checks if the range is horizontal.
func (r *Range2D) IsHorizontal() bool {
	return r.Start.Y == r.End.Y
}

// IsVertical checks if the range is vertical.
func (r *Range2D) IsVertical() bool {
	return r.Start.X == r.End.X
}

// Intersects checks if two Range2D intersect and returns the intersection point if they do.
// If the ranges have multiple intersection points, nil is returned.
// If the ranges are not 45-degree aligned, nil is returned.
func (r *Range2D) Intersects(other *Range2D) (intersection *tuples.Tuple, ok bool) {
	if r.Equals(other) {
		return nil, false
	}

	if (max(other.Start.X, other.End.X) < min(r.Start.X, r.End.X) &&
		max(other.Start.Y, other.End.Y) < min(r.Start.Y, r.End.Y)) ||
		(min(other.Start.Y, other.End.Y) > max(r.Start.Y, r.End.Y) &&
			min(other.Start.X, other.End.X) > max(r.Start.X, r.End.X)) {
		return nil, false
	}

	rExtrema := r.Extrema()
	otherExtrema := other.Extrema()

	if r.IsHorizontal() && other.IsVertical() &&
		other.Start.X >= rExtrema.MinX && other.Start.X <= rExtrema.MaxX &&
		r.Start.Y >= otherExtrema.MinY && r.Start.Y <= otherExtrema.MaxY {
		return tuples.NewTuple(other.Start.X, r.Start.Y), true
	} else if r.IsVertical() && other.IsHorizontal() &&
		r.Start.X >= otherExtrema.MinX && r.Start.X <= otherExtrema.MaxX &&
		other.Start.Y >= rExtrema.MinY && other.Start.Y <= rExtrema.MaxY {
		return tuples.NewTuple(r.Start.X, other.Start.Y), true
	}

	if r.Gradient() == other.Gradient() {
		return nil, false
	}

	if (r.Gradient() != -1 && r.Gradient() != 1) || (other.Gradient() != -1 && other.Gradient() != 1) {
		return nil, false
	}

	for x := rExtrema.MinX; x <= rExtrema.MaxX; x++ {
		var y int
		if r.Gradient() == 1 {
			y = r.Start.Y + (x - r.Start.X)
		} else {
			y = r.Start.Y - (x - r.Start.X)
		}

		if x >= otherExtrema.MinX && x <= otherExtrema.MaxX {
			var otherY int
			if other.Gradient() == 1 {
				otherY = other.Start.Y + (x - other.Start.X)
			} else {
				otherY = other.Start.Y - (x - other.Start.X)
			}

			if y == otherY {
				return tuples.NewTuple(x, y), true
			}
		}
	}

	return nil, false
}
