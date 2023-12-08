package matrix

type Range struct {
	start int
	end   int
}

func NewRange(start, end int) *Range {
	return &Range{
		start: min(start, end),
		end:   max(start, end),
	}
}

type Coordinate struct {
	X int
	Y int
}
