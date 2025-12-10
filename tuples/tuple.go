package tuples

type Tuple struct {
	X, Y int
}

func NewTuple(x, y int) *Tuple {
	return &Tuple{
		X: x,
		Y: y,
	}
}

// Equals checks if two tuples are equal.
// Two tuples are equal if their X and Y coordinates are the same.
func (t *Tuple) Equals(other *Tuple) bool {
	return t.X == other.X && t.Y == other.Y
}
