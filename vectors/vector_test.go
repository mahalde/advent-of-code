package vectors

import (
	"math"
	"testing"

	"github.com/mahalde/advent-of-code/assert"
)

func TestNewVector(t *testing.T) {
	v := NewVector(1, 2, 3)
	assert.Equals(t, v.X, 1)
	assert.Equals(t, v.Y, 2)
	assert.Equals(t, v.Z, 3)

	v2 := NewVector(-5, 0, 10)
	assert.Equals(t, v2.X, -5)
	assert.Equals(t, v2.Y, 0)
	assert.Equals(t, v2.Z, 10)
}

func TestVector_EuclideanDistance(t *testing.T) {
	tests := []struct {
		v1, v2   *Vector
		expected float64
	}{
		{NewVector(0, 0, 0), NewVector(0, 0, 0), 0.0},
		{NewVector(1, 0, 0), NewVector(0, 0, 0), 1.0},
		{NewVector(0, 1, 0), NewVector(0, 0, 0), 1.0},
		{NewVector(0, 0, 1), NewVector(0, 0, 0), 1.0},
		{NewVector(1, 2, 3), NewVector(0, 0, 0), math.Sqrt(14)}, // sqrt(1^2 + 2^2 + 3^2) = sqrt(14)
		{NewVector(1, 2, 3), NewVector(1, 2, 3), 0.0},
		{NewVector(1, 2, 3), NewVector(4, 6, 8), math.Sqrt(50)}, // sqrt((3^2 + 4^2 + 5^2)) = sqrt(50)
		{NewVector(-1, -2, -3), NewVector(0, 0, 0), math.Sqrt(14)},
		{NewVector(1, 1, 1), NewVector(2, 2, 2), math.Sqrt(3)}, // sqrt(1^2 + 1^2 + 1^2) = sqrt(3)
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			dist := tt.v1.EuclideanDistance(tt.v2)
			if math.Abs(dist-tt.expected) > 1e-9 {
				t.Errorf("EuclideanDistance(%v, %v) = %v, want %v", tt.v1, tt.v2, dist, tt.expected)
			}
		})
	}
}

func TestVector_Fields(t *testing.T) {
	v := &Vector{X: 10, Y: 20, Z: 30}
	assert.Equals(t, v.X, 10)
	assert.Equals(t, v.Y, 20)
	assert.Equals(t, v.Z, 30)

	// Test mutability
	v.X = 100
	assert.Equals(t, v.X, 100)
}
