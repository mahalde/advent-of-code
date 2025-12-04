package main

import (
	"fmt"

	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/matrix"
	"github.com/mahalde/advent-of-code/tuples"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "04"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "04", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	m := setupMatrix(input)

	accessibleRolls := findAccessibleRolls(m)

	return len(accessibleRolls)
}

func SolvePart2(input []string) int {
	m := setupMatrix(input)

	rolls := 0
	for accessibleRolls := findAccessibleRolls(m); len(accessibleRolls) > 0; accessibleRolls = findAccessibleRolls(m) {
		for _, roll := range accessibleRolls {
			m.Set(roll.X, roll.Y, '.')
		}
		rolls += len(accessibleRolls)
	}

	return rolls
}

func setupMatrix(input []string) matrix.Matrix[rune] {
	m := matrix.NewMatrix[rune](len(input[0]), len(input))
	for y, line := range input {
		for x, char := range line {
			m.Set(x, y, char)
		}
	}
	return m
}

func findAccessibleRolls(m matrix.Matrix[rune]) []tuples.Tuple {
	var accessibleRolls []tuples.Tuple

	for x := 0; x < m.Width(); x++ {
		for y := 0; y < m.Height(); y++ {
			if v, _ := m.Get(x, y); v != '@' {
				continue
			}
			adjacentRolls := 0
			m.EachAround(x, y, func(x, y int, value rune) {
				if value == '@' {
					adjacentRolls++
				}
			})
			if adjacentRolls < 4 {
				accessibleRolls = append(accessibleRolls, tuples.Tuple{X: x, Y: y})
			}
		}
	}

	return accessibleRolls
}
