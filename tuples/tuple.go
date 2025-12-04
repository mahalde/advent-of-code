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
