package main

import (
	"fmt"

	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/matrix"
	"github.com/mahalde/advent-of-code/tuples"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "07"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "07", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	m, start := parseInput(input)
	tachyonBeams := map[tuples.Tuple]struct{}{
		start: {},
	}

	finished := false
	splits := 0
	for !finished {
		nextBeams := make(map[tuples.Tuple]struct{})
		for t := range tachyonBeams {
			if t.Y == m.Height()-1 {
				finished = true
				break
			}

			nextChar, _ := m.Get(t.X, t.Y+1)
			if nextChar == '.' {
				nextBeams[tuples.Tuple{X: t.X, Y: t.Y + 1}] = struct{}{}
			} else if nextChar == '^' {
				nextBeams[tuples.Tuple{X: t.X - 1, Y: t.Y + 1}] = struct{}{}
				nextBeams[tuples.Tuple{X: t.X + 1, Y: t.Y + 1}] = struct{}{}

				splits++
			}
		}

		tachyonBeams = nextBeams
	}

	return splits
}

func SolvePart2(input []string) int {
	m, start := parseInput(input)

	return sendTachyonBeam(m, start)
}

func parseInput(input []string) (m matrix.Matrix[rune], start tuples.Tuple) {
	m = matrix.NewMatrix[rune](len(input[0]), len(input))
	for y, line := range input {
		for x, char := range line {
			m.Set(x, y, char)
			if char == 'S' {
				start = tuples.Tuple{X: x, Y: y}
			}
		}
	}
	return m, start
}

var splitterCache = make(map[tuples.Tuple]int)

func sendTachyonBeam(m matrix.Matrix[rune], position tuples.Tuple) int {
	timelines := 0

	for nextChar, _ := m.Get(position.X, position.Y+1); nextChar == '.'; nextChar, _ = m.Get(position.X, position.Y+1) {
		position.Y++
		if position.Y == m.Height()-1 {
			return 1
		}
	}

	if cached, ok := splitterCache[position]; ok {
		return cached
	}
	timelines += sendTachyonBeam(m, tuples.Tuple{X: position.X - 1, Y: position.Y + 1})
	timelines += sendTachyonBeam(m, tuples.Tuple{X: position.X + 1, Y: position.Y + 1})
	splitterCache[position] = timelines

	return timelines
}
