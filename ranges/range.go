package ranges

type Range struct {
	Start int
	End   int
}

func NewRange(start, end int) *Range {
	return &Range{
		Start: min(start, end),
		End:   max(start, end),
	}
}

func (r *Range) Contains(v int) bool {
	return v >= r.Start && v <= r.End
}
