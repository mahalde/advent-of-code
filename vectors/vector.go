package vectors

import (
	"math"
)

type Vector struct {
	X, Y, Z int
}

func NewVector(x, y, z int) *Vector {
	return &Vector{
		X: x,
		Y: y,
		Z: z,
	}
}

func (v *Vector) EuclideanDistance(other *Vector) float64 {
	dx := v.X - other.X
	dy := v.Y - other.Y
	dz := v.Z - other.Z
	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}
