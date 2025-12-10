package matrix_test

import (
	"slices"
	"testing"

	"github.com/mahalde/advent-of-code/assert"
	"github.com/mahalde/advent-of-code/matrix"
	"github.com/mahalde/advent-of-code/ranges"
)

func TestNewMatrix(t *testing.T) {
	m := matrix.NewMatrix[int](10, 15)

	assert.Equals(t, m.Width(), 10)
	assert.Equals(t, m.Height(), 15)

	m = matrix.NewMatrix[int](0, 0)
	assert.Equals(t, m.Width(), 0)
	assert.Equals(t, m.Height(), 0)
}

func TestMatrix_Get(t *testing.T) {
	m := matrix.NewMatrix[int](10, 10)

	success := m.Set(3, 5, 30)
	assert.True(t, success)

	got, ok := m.Get(3, 5)
	assert.True(t, ok)
	assert.Equals(t, got, 30)

	got, ok = m.Get(12, 4)
	assert.False(t, ok)
	assert.Equals(t, got, 0)
}

func TestMatrix_Set(t *testing.T) {
	m := matrix.NewMatrix[int](10, 10)

	success := m.Set(0, 0, 10)
	assert.True(t, success)
	got, _ := m.Get(0, 0)
	assert.Equals(t, got, 10)

	success = m.Set(-1, 0, 10)
	assert.False(t, success)

	success = m.Set(5, 10, 10)
	assert.False(t, success)
}

func TestMatrix_EachAround(t *testing.T) {
	testCases := []struct {
		name     string
		matrix   func() matrix.Matrix[int]
		x, y     int
		expected []struct{ x, y, val int }
	}{
		{
			name: "center of 3x3 matrix",
			matrix: func() matrix.Matrix[int] {
				m := matrix.NewMatrix[int](3, 3)
				for y := 0; y < 3; y++ {
					for x := 0; x < 3; x++ {
						m.Set(x, y, y*3+x+1)
					}
				}
				return m
			},
			x: 1, y: 1,
			expected: []struct{ x, y, val int }{
				{0, 0, 1}, {1, 0, 2}, {2, 0, 3},
				{0, 1, 4}, {2, 1, 6},
				{0, 2, 7}, {1, 2, 8}, {2, 2, 9},
			},
		},
		{
			name: "top-left corner of 3x3 matrix",
			matrix: func() matrix.Matrix[int] {
				m := matrix.NewMatrix[int](3, 3)
				for y := 0; y < 3; y++ {
					for x := 0; x < 3; x++ {
						m.Set(x, y, y*3+x+1)
					}
				}
				return m
			},
			x: 0, y: 0,
			expected: []struct{ x, y, val int }{
				{0, 1, 4}, {1, 0, 2}, {1, 1, 5},
			},
		},
		{
			name: "bottom-right corner of 3x3 matrix",
			matrix: func() matrix.Matrix[int] {
				m := matrix.NewMatrix[int](3, 3)
				for y := 0; y < 3; y++ {
					for x := 0; x < 3; x++ {
						m.Set(x, y, y*3+x+1)
					}
				}
				return m
			},
			x: 2, y: 2,
			expected: []struct{ x, y, val int }{
				{1, 1, 5}, {1, 2, 8}, {2, 1, 6},
			},
		},
		{
			name: "invalid position",
			matrix: func() matrix.Matrix[int] {
				return matrix.NewMatrix[int](3, 3)
			},
			x: -1, y: -1,
			expected: []struct{ x, y, val int }{},
		},
		{
			name: "1x1 matrix",
			matrix: func() matrix.Matrix[int] {
				m := matrix.NewMatrix[int](1, 1)
				m.Set(0, 0, 42)
				return m
			},
			x: 0, y: 0,
			expected: []struct{ x, y, val int }{},
		},
		{
			name: "2x2 matrix at (0,0)",
			matrix: func() matrix.Matrix[int] {
				m := matrix.NewMatrix[int](2, 2)
				m.Set(0, 0, 1)
				m.Set(1, 0, 2)
				m.Set(0, 1, 3)
				m.Set(1, 1, 4)
				return m
			},
			x: 0, y: 0,
			expected: []struct{ x, y, val int }{
				{0, 1, 3}, {1, 0, 2}, {1, 1, 4},
			},
		},
		{
			name: "2x2 matrix at (1,1)",
			matrix: func() matrix.Matrix[int] {
				m := matrix.NewMatrix[int](2, 2)
				m.Set(0, 0, 1)
				m.Set(1, 0, 2)
				m.Set(0, 1, 3)
				m.Set(1, 1, 4)
				return m
			},
			x: 1, y: 1,
			expected: []struct{ x, y, val int }{
				{0, 0, 1}, {0, 1, 3}, {1, 0, 2},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			m := tc.matrix()
			var calls []struct{ x, y, val int }
			m.EachAround(tc.x, tc.y, func(x, y int, val int) {
				calls = append(calls, struct{ x, y, val int }{x, y, val})
			})
			assert.Equals(t, len(calls), len(tc.expected))
			for _, exp := range tc.expected {
				found := false
				for _, call := range calls {
					if call.x == exp.x && call.y == exp.y && call.val == exp.val {
						found = true
						break
					}
				}
				if !found {
					t.Errorf("Expected call for (%d,%d)=%d not found", exp.x, exp.y, exp.val)
				}
			}
		})
	}
}

func TestMatrix_FindAdjacentValues(t *testing.T) {
	t.Run("found horizontal", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{3, 4, 5})
		expected, _ := ranges.NewRange2D(0, 2, 1, 1)
		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 1)
		assert.True(t, adjacentValues[0].Equals(expected))
	})

	t.Run("found vertical", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{0, 3, 6})
		expected, _ := ranges.NewRange2D(0, 0, 0, 2)
		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 1)
		assert.True(t, adjacentValues[0].Equals(expected))
	})

	t.Run("found diagonal", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{0, 4, 8})
		expected, _ := ranges.NewRange2D(0, 2, 0, 2)

		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 1)
		assert.True(t, adjacentValues[0].Equals(expected))
	})

	t.Run("found multiple", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{0, 1, 2})

		expected1, _ := ranges.NewRange2D(0, 0, 0, 2)
		expected2, _ := ranges.NewRange2D(1, 1, 0, 2)
		expected3, _ := ranges.NewRange2D(2, 2, 0, 2)
		expected4, _ := ranges.NewRange2D(0, 2, 0, 2)
		expected5, _ := ranges.NewRange2D(2, 0, 0, 2)
		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 5)
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected1)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected2)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected3)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected4)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected5)
		}))
	})

	t.Run("found none", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{0, 1, 3})

		assert.False(t, found)
		assert.Equals(t, len(adjacentValues), 0)
	})

	t.Run("found horizontal backwards", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{5, 4, 3})
		expected, _ := ranges.NewRange2D(2, 0, 1, 1)

		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 1)
		assert.True(t, adjacentValues[0].Equals(expected))
	})

	t.Run("found vertical backwards", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{6, 3, 0})
		expected, _ := ranges.NewRange2D(0, 0, 2, 0)

		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 1)
		assert.True(t, adjacentValues[0].Equals(expected))
	})

	t.Run("found diagonal backwards", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y*3+x)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{8, 4, 0})
		expected, _ := ranges.NewRange2D(2, 0, 2, 0)

		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 1)
		assert.True(t, adjacentValues[0].Equals(expected))
	})

	t.Run("found multiple backwards", func(t *testing.T) {
		m := matrix.NewMatrix[int](3, 3)

		for y := 0; y < 3; y++ {
			for x := 0; x < 3; x++ {
				m.Set(x, y, y)
			}
		}

		adjacentValues, found := m.FindAdjacentValues([]int{2, 1, 0})
		expected1, _ := ranges.NewRange2D(0, 0, 0, 2)
		expected2, _ := ranges.NewRange2D(1, 1, 0, 2)
		expected3, _ := ranges.NewRange2D(2, 2, 0, 2)
		expected4, _ := ranges.NewRange2D(0, 2, 0, 2)
		expected5, _ := ranges.NewRange2D(2, 0, 0, 2)

		assert.True(t, found)
		assert.Equals(t, len(adjacentValues), 5)
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected1)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected2)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected3)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected4)
		}))
		assert.True(t, slices.ContainsFunc(adjacentValues, func(r *ranges.Range2D) bool {
			return r.Equals(expected5)
		}))
	})
}
