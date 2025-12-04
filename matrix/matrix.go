package matrix

type Matrix[T comparable] [][]T

func NewMatrix[T comparable](w, h int) Matrix[T] {
	matrix := make([][]T, h)

	for i := 0; i < h; i++ {
		matrix[i] = make([]T, w)
	}

	return matrix
}

func (m Matrix[T]) Height() int {
	return len(m)
}

func (m Matrix[T]) Width() int {
	if len(m) == 0 {
		return 0
	}

	return len(m[0])
}

// Get returns the value at the given coordinates and a boolean indicating if the coordinates are valid.
func (m Matrix[T]) Get(x, y int) (value T, ok bool) {
	if y >= len(m) || y < 0 || x < 0 || x >= len(m[y]) {
		return value, false
	}
	return m[y][x], true
}

// Set sets the value at the given coordinates. It returns true if the coordinates are valid, false otherwise.
func (m Matrix[T]) Set(x, y int, value T) bool {
	if x < 0 || y < 0 || y >= len(m) || x >= len(m[y]) {
		return false
	}

	m[y][x] = value
	return true
}

// EachAround calls the given function for each adjacent cell around the (x, y) coordinates.
func (m Matrix[T]) EachAround(x, y int, fn func(x, y int, value T)) {
	if len(m) == 0 || y < 0 || y >= len(m) || x < 0 || x >= len(m[0]) {
		return
	}
	maxY := len(m) - 1
	if maxY <= 0 {
		return
	}

	maxX := len(m[y]) - 1
	for yy := max(0, y-1); yy <= min(y+1, maxY); yy++ {
		for xx := max(0, x-1); xx <= min(x+1, maxX); xx++ {
			if x == xx && y == yy {
				continue
			}
			val, _ := m.Get(xx, yy)
			fn(xx, yy, val)
		}
	}
}
