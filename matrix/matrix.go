package matrix

import (
	"slices"

	"github.com/mahalde/advent-of-code/ranges"
)

type Matrix[T comparable] [][]T

func NewMatrix[T comparable](w, h int) Matrix[T] {
	matrix := make([][]T, h)

	for i := 0; i < h; i++ {
		matrix[i] = make([]T, w)
	}

	return matrix
}

func (m Matrix[T]) Height() int {
	return len(m)
}

func (m Matrix[T]) Width() int {
	if len(m) == 0 {
		return 0
	}

	return len(m[0])
}

// Get returns the value at the given coordinates and a boolean indicating if the coordinates are valid.
func (m Matrix[T]) Get(x, y int) (value T, ok bool) {
	if y >= len(m) || y < 0 || x < 0 || x >= len(m[y]) {
		return value, false
	}
	return m[y][x], true
}

// Set sets the value at the given coordinates. It returns true if the coordinates are valid, false otherwise.
func (m Matrix[T]) Set(x, y int, value T) bool {
	if x < 0 || y < 0 || y >= len(m) || x >= len(m[y]) {
		return false
	}

	m[y][x] = value
	return true
}

// EachAround calls the given function for each adjacent cell around the (x, y) coordinates.
func (m Matrix[T]) EachAround(x, y int, fn func(x, y int, value T)) {
	if len(m) == 0 || y < 0 || y >= len(m) || x < 0 || x >= len(m[0]) {
		return
	}
	maxY := len(m) - 1
	if maxY <= 0 {
		return
	}

	maxX := len(m[y]) - 1
	for yy := max(0, y-1); yy <= min(y+1, maxY); yy++ {
		for xx := max(0, x-1); xx <= min(x+1, maxX); xx++ {
			if x == xx && y == yy {
				continue
			}
			val, _ := m.Get(xx, yy)
			fn(xx, yy, val)
		}
	}
}

// FindAdjacentValues is used to search for a chain of values adjacent to each other in the matrix.
// If found, it returns the ranges of the found values and true, otherwise it returns an empty slice and false.
func (m Matrix[T]) FindAdjacentValues(values []T) (foundRanges []*ranges.Range2D, found bool) {
	for y := 0; y < m.Height(); y++ {
		for x := 0; x < m.Width(); x++ {
			f, found := m.FindAdjacentValuesForCell(x, y, values)
			if found {
				foundRanges = append(foundRanges, f...)
			}
		}
	}

	return foundRanges, len(foundRanges) > 0
}

func (m Matrix[T]) FindAdjacentValuesForCell(x, y int, values []T) (foundRanges []*ranges.Range2D, found bool) {
	val, ok := m.Get(x, y)
	if !ok || val != values[0] {
		return nil, false
	}

	for yDelta := -1; yDelta <= 1; yDelta++ {
		for xDelta := -1; xDelta <= 1; xDelta++ {
			if xDelta == 0 && yDelta == 0 {
				continue
			}

			cx, cy := x+xDelta, y+yDelta
			depth := 1
			for depth < len(values) {
				if cx < 0 || cx >= m.Width() || cy < 0 || cy >= m.Height() {
					break
				}
				val, ok := m.Get(cx, cy)
				if !ok || val != values[depth] {
					break
				}
				cx += xDelta
				cy += yDelta
				depth++
			}
			if depth == len(values) {
				endX := cx - xDelta
				endY := cy - yDelta
				foundRange, _ := ranges.NewRange2D(x, endX, y, endY)
				if !slices.ContainsFunc(foundRanges, func(r *ranges.Range2D) bool {
					return r.Equals(foundRange)
				}) {
					foundRanges = append(foundRanges, foundRange)
				}
			}
		}
	}

	return foundRanges, len(foundRanges) > 0
}
