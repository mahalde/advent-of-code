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

func TestExclude(t *testing.T) {
	test := []struct {
		rStart, rEnd         int
		otherStart, otherEnd int
		expected             []*ranges.Range
	}{
		// No overlap, other is before
		{5, 10, 0, 4, []*ranges.Range{ranges.NewRange(5, 10)}},
		// No overlap, other is after
		{5, 10, 11, 15, []*ranges.Range{ranges.NewRange(5, 10)}},
		// Other covers part at start
		{5, 10, 3, 7, []*ranges.Range{ranges.NewRange(8, 10)}},
		// Other covers part at end
		{5, 10, 8, 12, []*ranges.Range{ranges.NewRange(5, 7)}},
		// Other covers middle
		{5, 10, 6, 8, []*ranges.Range{ranges.NewRange(5, 5), ranges.NewRange(9, 10)}},
		// Other covers all
		{5, 10, 3, 12, []*ranges.Range{}},
		// Other covers exactly
		{5, 10, 5, 10, []*ranges.Range{}},
		// Other is inside, no exclusion
		{5, 10, 6, 9, []*ranges.Range{ranges.NewRange(5, 5), ranges.NewRange(10, 10)}},
	}

	for _, tt := range test {
		r := ranges.NewRange(tt.rStart, tt.rEnd)
		other := ranges.NewRange(tt.otherStart, tt.otherEnd)
		result := r.Exclude(other)
		assert.Equals(t, len(result), len(tt.expected))
		for i, exp := range tt.expected {
			assert.True(t, result[i].Start == exp.Start && result[i].End == exp.End)
		}
	}
}

func TestIntersect(t *testing.T) {
	test := []struct {
		rStart, rEnd         int
		otherStart, otherEnd int
		expected             []*ranges.Range
	}{
		// No overlap
		{5, 10, 0, 4, []*ranges.Range{}},
		// Partial overlap at start
		{5, 10, 3, 7, []*ranges.Range{ranges.NewRange(5, 7)}},
		// Partial overlap at end
		{5, 10, 8, 12, []*ranges.Range{ranges.NewRange(8, 10)}},
		// Full overlap
		{5, 10, 5, 10, []*ranges.Range{ranges.NewRange(5, 10)}},
		// Other inside
		{5, 10, 6, 9, []*ranges.Range{ranges.NewRange(6, 9)}},
		// Other covers all
		{5, 10, 3, 12, []*ranges.Range{ranges.NewRange(5, 10)}},
	}

	for _, tt := range test {
		r := ranges.NewRange(tt.rStart, tt.rEnd)
		other := ranges.NewRange(tt.otherStart, tt.otherEnd)
		result := r.Intersect(other)
		assert.Equals(t, len(result), len(tt.expected))
		for i, exp := range tt.expected {
			assert.True(t, result[i].Start == exp.Start && result[i].End == exp.End)
		}
	}
}
