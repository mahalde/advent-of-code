package scanline

import (
	"sort"

	"github.com/mahalde/advent-of-code/ranges"
	"github.com/mahalde/advent-of-code/tuples"
)

// Span groups consecutive rows that have identical merged segments.
type Span struct {
	YStart, YEnd int
	Segs         []*ranges.Range
}

type verticalEdge struct{ x, yMin, yMax int }
type horizontalEdge struct{ y, xMin, xMax int }

// BuildSpans computes the compressed row spans for a polygon defined by a list of points.
// It uses a sweep-line algorithm to determine, for each row in the bounding box, which x-ranges are "green" (interior or boundary).
// To avoid storing a full grid, it groups consecutive rows with identical green x-ranges into "spans".
//
// Parameters:
//   - points: A slice of tuples representing the vertices of the polygon in counter-clockwise order.
//     Each tuple has X and Y integer coordinates.
//
// Returns:
//   - offsetX, offsetY: Offsets applied to shift coordinates into a 0-based local grid with 1-cell padding.
//     Use these to convert world coordinates to grid indices: gridX = worldX - offsetX.
//   - width, height: Dimensions of the local grid (bounding box + padding).
//   - spans: A slice of Span structs, each representing a vertical range of rows [YStart, YEnd] (inclusive)
//     where the green x-ranges are identical. Each Span contains Segs, a list of merged inclusive x-ranges
//     (as *ranges.Range) that are allowed (green or red) on those rows.
func BuildSpans(points []*tuples.Tuple) (offsetX, offsetY, width, height int, spans []Span) {
	if len(points) == 0 {
		return 0, 0, 0, 0, nil
	}

	// compute bounding box
	minX, maxX := points[0].X, points[0].X
	minY, maxY := points[0].Y, points[0].Y
	for _, p := range points {
		if p.X < minX {
			minX = p.X
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}

	offsetX = minX - 1
	offsetY = minY - 1
	width = maxX - minX + 3
	height = maxY - minY + 3

	var verticalEdges []verticalEdge
	var horizontalEdges []horizontalEdge
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		p2 := points[(i+1)%len(points)]
		x1 := p1.X - offsetX
		y1 := p1.Y - offsetY
		x2 := p2.X - offsetX
		y2 := p2.Y - offsetY
		if x1 == x2 {
			yMin, yMax := y1, y2
			if y1 > y2 {
				yMin, yMax = y2, y1
			}
			verticalEdges = append(verticalEdges, verticalEdge{x: x1, yMin: yMin, yMax: yMax})
		} else {
			xMin, xMax := x1, x2
			if x1 > x2 {
				xMin, xMax = x2, x1
			}
			horizontalEdges = append(horizontalEdges, horizontalEdge{y: y1, xMin: xMin, xMax: xMax})
		}
	}

	eventsAdd := make(map[int][]int)
	eventsRem := make(map[int][]int)
	for _, e := range verticalEdges {
		eventsAdd[e.yMin] = append(eventsAdd[e.yMin], e.x)
		eventsRem[e.yMax] = append(eventsRem[e.yMax], e.x)
	}

	// helper to get boundary intervals
	getBoundary := func(y int) []*ranges.Range {
		var b []*ranges.Range
		for _, p := range points {
			if p.Y-offsetY == y {
				x := p.X - offsetX
				b = append(b, ranges.NewRange(x, x))
			}
		}
		for _, edge := range verticalEdges {
			if edge.yMin <= y && y <= edge.yMax {
				b = append(b, ranges.NewRange(edge.x, edge.x))
			}
		}
		for _, h := range horizontalEdges {
			if h.y == y {
				b = append(b, ranges.NewRange(h.xMin, h.xMax))
			}
		}
		return b
	}

	buildSegs := func(active map[int]int, y int) []*ranges.Range {
		var curXList []int
		for x := range active {
			curXList = append(curXList, x)
		}
		sort.Ints(curXList)
		var intervals []*ranges.Range
		for i := 0; i+1 < len(curXList); i += 2 {
			startX := curXList[i]
			endX := curXList[i+1] - 1
			if startX <= endX {
				intervals = append(intervals, ranges.NewRange(startX, endX))
			}
		}
		intervals = append(intervals, getBoundary(y)...)
		sort.Slice(intervals, func(i, j int) bool { return intervals[i].Start < intervals[j].Start })
		var merged []*ranges.Range
		for _, seg := range intervals {
			if len(merged) == 0 || merged[len(merged)-1].End+1 < seg.Start {
				merged = append(merged, seg)
			} else {
				if seg.End > merged[len(merged)-1].End {
					merged[len(merged)-1].End = seg.End
				}
			}
		}
		return merged
	}

	active := make(map[int]int)
	if adds, ok := eventsAdd[0]; ok {
		for _, x := range adds {
			active[x]++
		}
	}
	var prevSegs []*ranges.Range
	spanStart := 0
	for y := 1; y < height; y++ {
		if adds, ok := eventsAdd[y]; ok {
			for _, x := range adds {
				active[x]++
			}
		}
		if rems, ok := eventsRem[y]; ok {
			for _, x := range rems {
				active[x]--
				if active[x] <= 0 {
					delete(active, x)
				}
			}
		}
		curSegs := buildSegs(active, y)
		if !segmentsEqual(prevSegs, curSegs) {
			spans = append(spans, Span{YStart: spanStart, YEnd: y - 1, Segs: prevSegs})
			spanStart = y
			prevSegs = curSegs
		}
	}

	return offsetX, offsetY, width, height, spans
}

func segmentsEqual(a, b []*ranges.Range) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i].Start != b[i].Start || a[i].End != b[i].End {
			return false
		}
	}
	return true
}
