package main

import (
	"fmt"
	"slices"

	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/matrix"
	"github.com/mahalde/advent-of-code/ranges"
	"github.com/mahalde/advent-of-code/tuples"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2024", "04"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2024", "04", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	searchValues := []rune{'X', 'M', 'A', 'S'}
	xmasMatrix := createMatrix(input)

	foundValues, found := xmasMatrix.FindAdjacentValues(searchValues)
	if !found {
		return 0
	}

	return len(foundValues)
}

func SolvePart2(input []string) int {
	searchValues := []rune{'M', 'A', 'S'}
	xmasMatrix := createMatrix(input)

	foundValues, found := xmasMatrix.FindAdjacentValues(searchValues)
	if !found {
		return 0
	}

	var foundIntersections []*tuples.Tuple

	// Filter out horizontal and vertical ranges
	foundValues = slices.DeleteFunc(foundValues, func(r *ranges.Range2D) bool {
		return r.IsHorizontal() || r.IsVertical()
	})

	for _, value := range foundValues {
		for _, value2 := range foundValues {
			if value == value2 {
				continue
			}

			cell, intersects := value.Intersects(value2)
			if !intersects {
				continue
			}

			if val, _ := xmasMatrix.Get(cell.X, cell.Y); val != 'A' {
				continue
			}

			if slices.ContainsFunc(foundIntersections, func(t *tuples.Tuple) bool {
				return t.Equals(cell)
			}) {
				continue
			}

			foundIntersections = append(foundIntersections, cell)
		}
	}

	return len(foundIntersections)
}

func createMatrix(input []string) matrix.Matrix[rune] {
	xmasMatrix := matrix.NewMatrix[rune](len(input[0]), len(input))
	for y, line := range input {
		for x, char := range line {
			xmasMatrix.Set(x, y, char)
		}
	}
	return xmasMatrix
}
