package matrix

import (
	"github.com/mahalde/advent-of-code/utils/assert"
	"testing"
)

func TestNewMatrix(t *testing.T) {
	matrix := NewMatrix[int](10, 10)

	assert.IntEquals(t, len(matrix), 10)
	assert.IntEquals(t, len(matrix[5]), 10)
}

func TestMatrix_Set(t *testing.T) {
	matrix := NewMatrix[int](10, 10)

	success := matrix.Set(0, 0, 10)
	assert.True(t, success)
	got, _ := matrix.Get(0, 0)
	assert.IntEquals(t, got, 10)

	success = matrix.Set(-1, 0, 10)
	assert.False(t, success)

	success = matrix.Set(5, 10, 10)
	assert.False(t, success)
}

func TestMatrix_Get(t *testing.T) {
	matrix := NewMatrix[int](10, 10)

	success := matrix.Set(3, 5, 30)
	assert.True(t, success)

	got, ok := matrix.Get(3, 5)
	assert.True(t, ok)
	assert.IntEquals(t, got, 30)

	got, ok = matrix.Get(12, 4)
	assert.False(t, ok)
	assert.IntEquals(t, got, 0)
}

func TestMatrix_EachAround(t *testing.T) {
	matrix := NewMatrix[int](3, 3)

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {
			matrix.Set(x, y, y*3+x)
		}
	}

	calls := 0
	fn := func(x, y int, value int) {
		calls++
		assert.IntEquals(t, value, y*3+x)
	}

	matrix.EachAround(1, 1, fn)
	assert.IntEquals(t, calls, 8)

	calls = 0
	matrix.EachAround(2, 2, fn)
	assert.IntEquals(t, calls, 3)

	calls = 0
	matrix.EachAround(0, 1, fn)
	assert.IntEquals(t, calls, 5)
}

func TestMatrix_EachAroundRange(t *testing.T) {
	matrix := NewMatrix[int](5, 5)

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			matrix.Set(x, y, y*5+x)
		}
	}

	calls := 0
	fn := func(x, y int, value int) {
		calls++
		assert.IntEquals(t, value, y*5+x)
	}

	xRange := NewRange(1, 3)
	yRange := NewRange(1, 3)
	matrix.EachAroundRange(xRange, yRange, fn)

	assert.IntEquals(t, calls, 16)
}
