package matrix

type Matrix[T any] [][]T

func NewMatrix[T any](w, h int) Matrix[T] {
	matrix := make([][]T, h)

	for i := 0; i < h; i++ {
		matrix[i] = make([]T, w)
	}

	return matrix
}
func (m Matrix[T]) Get(x, y int) (value T, ok bool) {
	if y >= len(m) || y < 0 || x < 0 || x >= len(m[y]) {
		return value, false
	}
	return m[y][x], true
}
func (m Matrix[T]) Set(x, y int, value T) bool {
	if x < 0 || y < 0 || y >= len(m) || x >= len(m[y]) {
		return false
	}

	m[y][x] = value
	return true
}

func (m Matrix[T]) EachAround(x, y int, fn func(x, y int, value T)) {
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

func (m Matrix[T]) EachAroundRange(x, y *Range, fn func(x, y int, value T)) {
	maxY := len(m) - 1
	if maxY <= 0 {
		return
	}

	maxX := len(m[0]) - 1
	for yy := max(0, y.start-1); yy <= min(y.end+1, maxY); yy++ {
		for xx := max(0, x.start-1); xx <= min(x.end+1, maxX); xx++ {
			if xx >= x.start && xx <= x.end && yy >= y.start && yy <= y.end {
				continue
			}

			val, _ := m.Get(xx, yy)
			fn(xx, yy, val)
		}
	}
}
