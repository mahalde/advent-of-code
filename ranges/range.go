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

func (r *Range) Size() int {
	return r.End - r.Start + 1
}

// Exclude returns the parts of the range that are not covered by the other range.
func (r *Range) Exclude(other *Range) []*Range {
	var result []*Range

	if other.Start > r.End || other.End < r.Start {
		// No overlap
		result = append(result, NewRange(r.Start, r.End))
		return result
	}

	if other.Start > r.Start {
		result = append(result, NewRange(r.Start, other.Start-1))
	}

	if other.End < r.End {
		result = append(result, NewRange(other.End+1, r.End))
	}

	return result
}

// Intersect returns the parts of the range that are covered by the other range.
func (r *Range) Intersect(other *Range) []*Range {
	var result []*Range

	start := max(r.Start, other.Start)
	end := min(r.End, other.End)

	if start <= end {
		result = append(result, NewRange(start, end))
	}

	return result
}
