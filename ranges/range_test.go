package ranges_test

import (
	"testing"

	"github.com/mahalde/advent-of-code/assert"
	"github.com/mahalde/advent-of-code/ranges"
)

func TestContains(t *testing.T) {
	test := []struct {
		start    int
		end      int
		number   int
		expected bool
	}{
		{0, 10, 5, true},
		{0, 10, 0, true},
		{0, 10, 10, true},
		{0, 10, 11, false},
		{0, 10, -1, false},
		{10, 0, 5, true},
		{10, 0, 0, true},
		{10, 0, 10, true},
		{10, 0, 11, false},
		{10, 0, -1, false},
	}

	for _, tt := range test {
		r := ranges.NewRange(tt.start, tt.end)
		assert.Equals(t, r.Contains(tt.number), tt.expected)
	}
}
