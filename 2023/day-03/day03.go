package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/matrix"
	"regexp"

	"github.com/mahalde/advent-of-code/utils/files"
)

var (
	numRegEx    = regexp.MustCompile("\\d+")
	symbolRegEx = regexp.MustCompile("[^\\w\\s.]")
	gearRegEx   = regexp.MustCompile("\\*")
)

func main() {
	input := files.ReadFile(3, 2023, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

type part struct {
	line   int
	start  int
	end    int
	number int
}

func solvePart1(input []string) int {
	result := 0

	var parts []part
	symbols := matrix.NewMatrix[bool](len(input), len(input))
	for lineIdx, line := range input {
		numberIndices := numRegEx.FindAllStringIndex(line, -1)
		numbers := numRegEx.FindAllString(line, -1)

		for idx, num := range numbers {
			parts = append(parts, part{
				line:   lineIdx,
				start:  numberIndices[idx][0],
				end:    numberIndices[idx][1] - 1,
				number: conv.ToInt(num),
			})
		}

		symbolIndices := symbolRegEx.FindAllStringIndex(line, -1)
		for _, idx := range symbolIndices {
			symbols.Set(idx[0], lineIdx, true)
		}
	}

	for _, part := range parts {
		xRange := matrix.NewRange(part.start, part.end)
		yRange := matrix.NewRange(part.line, part.line)
		foundSymbol := false

		symbols.EachAroundRange(xRange, yRange, func(x, y int, value bool) {
			foundSymbol = foundSymbol || value
		})

		if foundSymbol {
			result += part.number
		}
	}
	return result
}

func solvePart2(input []string) int {
	result := 0

	var symbols []matrix.Coordinate
	parts := matrix.NewMatrix[int](len(input), len(input))
	for lineIdx, line := range input {
		numberIndices := numRegEx.FindAllStringIndex(line, -1)
		numbers := numRegEx.FindAllString(line, -1)

		for idx, num := range numbers {
			for x := numberIndices[idx][0]; x < numberIndices[idx][1]; x++ {
				parts.Set(x, lineIdx, conv.ToInt(num))
			}
		}

		symbolIndices := gearRegEx.FindAllStringIndex(line, -1)
		for _, idx := range symbolIndices {
			symbols = append(symbols, matrix.Coordinate{X: idx[0], Y: lineIdx})
		}
	}

	for _, gear := range symbols {
		surroundedParts := make(map[int]bool)

		parts.EachAround(gear.X, gear.Y, func(x, y int, value int) {
			if value != 0 {
				surroundedParts[value] = true
			}
		})

		if len(surroundedParts) == 2 {
			gearRatio := 1
			for part := range surroundedParts {
				gearRatio *= part
			}

			result += gearRatio
		}
	}
	return result
}
