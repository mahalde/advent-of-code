package math_test

import (
	"testing"

	"github.com/mahalde/advent-of-code/math"
)

func TestAbs(t *testing.T) {
	test := []struct {
		number   int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{10, 10},
		{-10, 10},
	}

	for _, tt := range test {
		if math.Abs(tt.number) != tt.expected {
			t.Errorf("Expected %v, but got %v", tt.expected, math.Abs(tt.number))
		}
	}
}
