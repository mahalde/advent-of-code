package main

import (
	"fmt"
	"strings"

	"github.com/mahalde/advent-of-code/algorithms/scanline"
	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/math"
	"github.com/mahalde/advent-of-code/tuples"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "09"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "09", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	points := parseInput(input)

	maxArea := 0
	for _, p := range points {
		for _, p2 := range points {
			if p == p2 {
				continue
			}

			area := math.Abs(p2.X-p.X+1) * math.Abs(p2.Y-p.Y+1)
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func SolvePart2(input []string) int {
	points := parseInput(input)

	if len(points) == 0 {
		return 0
	}

	offsetX, offsetY, _, _, spans := scanline.BuildSpans(points)

	maxArea := 0
	for _, p1 := range points {
		for _, p2 := range points {
			if p1 == p2 {
				continue
			}
			x1, y1 := p1.X, p1.Y
			x2, y2 := p2.X, p2.Y
			minXX, maxXX := min(x1, x2), max(x1, x2)
			minYY, maxYY := min(y1, y2), max(y1, y2)

			area := (maxXX - minXX + 1) * (maxYY - minYY + 1)
			valid := true
			gx1 := minXX - offsetX
			gx2 := maxXX - offsetX
			// iterate spans overlapping [minYY-offsetY, maxYY-offsetY]
			needStart := minYY - offsetY
			needEnd := maxYY - offsetY
			coveredRows := 0
			for _, sp := range spans {
				if sp.YEnd < needStart || sp.YStart > needEnd {
					continue
				}
				// overlap range
				sy0 := max(sp.YStart, needStart)
				sy1 := min(sp.YEnd, needEnd)
				rows := sy1 - sy0 + 1
				// check if all rows in this span are covered for [gx1,gx2]
				// since segs same for whole span, just check once
				covered := 0
				for _, seg := range sp.Segs {
					if seg.Start <= gx2 && seg.End >= gx1 {
						left := max(seg.Start, gx1)
						right := min(seg.End, gx2)
						if left <= right {
							covered += right - left + 1
						}
					}
				}
				if covered != (gx2 - gx1 + 1) {
					valid = false
					break
				}
				coveredRows += rows
			}
			if valid && coveredRows == (needEnd-needStart+1) && area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func parseInput(input []string) []*tuples.Tuple {
	points := make([]*tuples.Tuple, len(input))
	for i, line := range input {
		parts := strings.Split(line, ",")
		points[i] = tuples.NewTuple(conv.ToInt(parts[0]), conv.ToInt(parts[1]))
	}
	return points
}
