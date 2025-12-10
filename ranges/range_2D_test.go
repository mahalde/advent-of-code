package ranges_test

import (
	"fmt"
	"testing"

	"github.com/mahalde/advent-of-code/ranges"
)

func TestNewRange2D(t *testing.T) {
	test := []struct {
		xStart   int
		xEnd     int
		yStart   int
		yEnd     int
		expected bool
	}{
		{0, 10, 0, 10, true},
		{0, 10, 0, 11, false},
		{0, 10, 1, 10, false},
		{0, 10, 1, 11, true},
		{0, 10, 10, 0, true},
		{0, 10, 11, 0, false},
		{0, 0, 12, 14, true},
		{12, 14, 2, 2, true},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("start: (%v, %v) end: (%v, %v)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			if r == nil && tt.expected {
				t.Errorf("Expected a range, but got nil")
			}
			if r != nil && !tt.expected {
				t.Errorf("Expected nil, but got a range")
			}
		})
	}
}

func TestRange2D_Equals(t *testing.T) {
	test := []struct {
		xStart   int
		xEnd     int
		yStart   int
		yEnd     int
		xStart2  int
		xEnd2    int
		yStart2  int
		yEnd2    int
		expected bool
	}{
		{0, 10, 0, 10, 0, 10, 0, 10, true},
		{0, 10, 0, 10, 0, 11, 0, 11, false},
		{0, 10, 0, 10, 1, 10, 1, 10, false},
		{0, 10, 0, 10, 14, 10, 14, 10, false},
		{0, 10, 0, 10, 1, 10, 1, 10, false},
		{0, 10, 0, 10, 10, 0, 10, 0, true},
	}

	for _, tt := range test {
		t.Run(fmt.Sprintf("start: (%v, %v) end: (%v, %v) start2: (%v, %v) end2: (%v, %v)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd, tt.xStart2, tt.yStart2, tt.xEnd2, tt.yEnd2), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			r2, _ := ranges.NewRange2D(tt.xStart2, tt.xEnd2, tt.yStart2, tt.yEnd2)
			if r.Equals(r2) != tt.expected {
				t.Errorf("Expected %v, but got %v", tt.expected, r.Equals(r2))
			}
		})
	}
}

func TestRange2D_Delta(t *testing.T) {
	tests := []struct {
		xStart, xEnd, yStart, yEnd     int
		expectedDeltaX, expectedDeltaY int
	}{
		{0, 5, 0, 5, 1, 1},
		{5, 0, 5, 0, -1, -1},
		{0, 5, 0, 0, 1, 0},
		{0, 0, 0, 5, 0, 1},
		{10, 5, 10, 5, -1, -1},
		{0, 10, 0, 10, 1, 1},
		{0, 0, 0, 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start: (%d, %d) end: (%d, %d)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			delta := r.Delta()
			if delta.X != tt.expectedDeltaX || delta.Y != tt.expectedDeltaY {
				t.Errorf("Expected delta (%d, %d), got (%d, %d)", tt.expectedDeltaX, tt.expectedDeltaY, delta.X, delta.Y)
			}
		})
	}
}

func TestRange2D_Gradient(t *testing.T) {
	tests := []struct {
		xStart, xEnd, yStart, yEnd int
		expectedGradient           float64
	}{
		{0, 5, 0, 5, 1.0},
		{0, 5, 5, 0, -1.0},
		{0, 5, 0, 0, 0.0}, // horizontal, returns 0
		{0, 0, 0, 5, 0.0}, // vertical, returns 0
		{5, 0, 5, 0, 1.0},
		{0, 10, 0, 10, 1.0},
		{0, 0, 0, 0, 0.0}, // point, but since x same, 0
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start: (%d, %d) end: (%d, %d)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			gradient := r.Gradient()
			if gradient != tt.expectedGradient {
				t.Errorf("Expected gradient %f, got %f", tt.expectedGradient, gradient)
			}
		})
	}
}

func TestRange2D_Extrema(t *testing.T) {
	tests := []struct {
		xStart, xEnd, yStart, yEnd                             int
		expectedMinX, expectedMaxX, expectedMinY, expectedMaxY int
	}{
		{0, 5, 0, 5, 0, 5, 0, 5},
		{5, 0, 5, 0, 0, 5, 0, 5},
		{0, 5, 0, 0, 0, 5, 0, 0},
		{0, 0, 0, 5, 0, 0, 0, 5},
		{10, 5, 10, 5, 5, 10, 5, 10},
		{0, 10, 0, 10, 0, 10, 0, 10},
		{0, 0, 0, 0, 0, 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start: (%d, %d) end: (%d, %d)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			extrema := r.Extrema()
			if extrema.MinX != tt.expectedMinX || extrema.MaxX != tt.expectedMaxX ||
				extrema.MinY != tt.expectedMinY || extrema.MaxY != tt.expectedMaxY {
				t.Errorf("Expected extrema MinX:%d MaxX:%d MinY:%d MaxY:%d, got MinX:%d MaxX:%d MinY:%d MaxY:%d",
					tt.expectedMinX, tt.expectedMaxX, tt.expectedMinY, tt.expectedMaxY,
					extrema.MinX, extrema.MaxX, extrema.MinY, extrema.MaxY)
			}
		})
	}
}

func TestRange2D_IsHorizontal(t *testing.T) {
	tests := []struct {
		xStart, xEnd, yStart, yEnd int
		expected                   bool
	}{
		{0, 5, 0, 0, true},
		{0, 5, 0, 5, false},
		{0, 0, 0, 5, false},
		{5, 0, 5, 5, true},
		{0, 10, 0, 10, false},
		{0, 0, 0, 0, true},
		{3, 7, 2, 2, true},
		{7, 3, 8, 8, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start: (%d, %d) end: (%d, %d)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			isHorizontal := r.IsHorizontal()
			if isHorizontal != tt.expected {
				t.Errorf("Expected IsHorizontal %v, got %v", tt.expected, isHorizontal)
			}
		})
	}
}

func TestRange2D_IsVertical(t *testing.T) {
	tests := []struct {
		xStart, xEnd, yStart, yEnd int
		expected                   bool
	}{
		{0, 0, 0, 5, true},
		{0, 5, 0, 5, false},
		{0, 5, 0, 0, false},
		{5, 5, 0, 5, true},
		{0, 10, 0, 10, false},
		{0, 0, 0, 0, true},
		{3, 3, 2, 8, true},
		{7, 7, 8, 2, true},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("start: (%d, %d) end: (%d, %d)", tt.xStart, tt.yStart, tt.xEnd, tt.yEnd), func(t *testing.T) {
			r, _ := ranges.NewRange2D(tt.xStart, tt.xEnd, tt.yStart, tt.yEnd)
			isVertical := r.IsVertical()
			if isVertical != tt.expected {
				t.Errorf("Expected IsVertical %v, got %v", tt.expected, isVertical)
			}
		})
	}
}

func TestRange2D_Intersects(t *testing.T) {
	tests := []struct {
		r1xStart, r1xEnd, r1yStart, r1yEnd int
		r2xStart, r2xEnd, r2yStart, r2yEnd int
		expectedX, expectedY               int
		expectedOk                         bool
	}{
		// Identical ranges
		{0, 2, 0, 2, 0, 2, 0, 2, 0, 0, false},
		// Parallel horizontals
		{0, 2, 0, 0, 0, 2, 1, 1, 0, 0, false},
		// Parallel verticals
		{0, 0, 0, 2, 1, 1, 0, 2, 0, 0, false},
		// Parallel diagonals
		{0, 2, 0, 2, 1, 3, 1, 3, 0, 0, false},
		// Horizontal and vertical crossing
		{0, 2, 0, 0, 1, 1, -1, 1, 1, 0, true},
		// Vertical and horizontal crossing
		{1, 1, -1, 1, 0, 2, 0, 0, 1, 0, true},
		// Diagonals crossing
		{0, 2, 0, 2, 0, 2, 2, 0, 1, 1, true},
		// No intersection, horizontal and vertical not crossing
		{0, 2, 0, 0, 3, 3, -1, 1, 0, 0, false},
		// No intersection, diagonals not crossing
		{0, 2, 0, 2, 3, 5, 3, 5, 0, 0, false},
		// Horizontal and vertical, vertical outside horizontal range
		{0, 1, 0, 0, 2, 2, -1, 1, 0, 0, false},
		// Diagonals, one not 45 but since NewRange2D prevents, but assume
		// But since all are 45, ok
		// Horizontal and vertical crossing with reversed ranges
		{2, 0, 0, 0, 1, 1, 1, -1, 1, 0, true},
		// Diagonals crossing with reversed
		{2, 0, 2, 0, 2, 0, 0, 2, 1, 1, true},
		// Horizontal and diagonal, should not intersect
		{0, 2, 0, 0, 0, 2, 0, 2, 0, 0, false},
		// Vertical and diagonal, should not intersect
		{0, 0, 0, 2, 0, 2, 0, 2, 0, 0, false},
		// Diagonals touching at end points
		{0, 2, 0, 2, 2, 4, 2, 0, 2, 2, true},
		// Falling diagonal and rising diagonal crossing
		{0, 2, 2, 0, 0, 2, 0, 2, 1, 1, true},
		// 45deg diagonals, not crossing
		{0, 2, 0, 2, 3, 5, 2, 0, 0, 0, false},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("r1:(%d,%d)-(%d,%d) r2:(%d,%d)-(%d,%d)", tt.r1xStart, tt.r1yStart, tt.r1xEnd, tt.r1yEnd, tt.r2xStart, tt.r2yStart, tt.r2xEnd, tt.r2yEnd), func(t *testing.T) {
			r1, _ := ranges.NewRange2D(tt.r1xStart, tt.r1xEnd, tt.r1yStart, tt.r1yEnd)
			r2, _ := ranges.NewRange2D(tt.r2xStart, tt.r2xEnd, tt.r2yStart, tt.r2yEnd)
			intersection, ok := r1.Intersects(r2)
			if ok != tt.expectedOk {
				t.Errorf("Expected ok %v, got %v", tt.expectedOk, ok)
			}
			if tt.expectedOk {
				if intersection == nil || intersection.X != tt.expectedX || intersection.Y != tt.expectedY {
					t.Errorf("Expected intersection (%d, %d), got %v", tt.expectedX, tt.expectedY, intersection)
				}
			} else {
				if intersection != nil {
					t.Errorf("Expected no intersection, got (%d, %d)", intersection.X, intersection.Y)
				}
			}
		})
	}
}
