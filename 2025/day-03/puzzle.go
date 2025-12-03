package main

import (
	"fmt"
	"math"

	"github.com/mahalde/advent-of-code/files"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "03"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "03", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	totalJoltage := 0
	for _, line := range input {
		totalJoltage += findMaxJoltage(line, 2)
	}

	return totalJoltage
}

func SolvePart2(input []string) int {
	totalJoltage := 0
	for _, line := range input {
		totalJoltage += findMaxJoltage(line, 12)
	}

	return totalJoltage
}

func findMaxJoltage(line string, size int) int {
	digits := make([]rune, size)
	for i, char := range line {
		for di, digit := range digits {
			if char > digit && len(line)-i >= size-di {
				digits[di] = char
				copy(digits[di+1:], make([]rune, size-di-1))
				break
			}
		}
	}

	maxJoltage := 0
	for i, digit := range digits {
		maxJoltage += int(digit-'0') * int(math.Pow10(size-i-1))
	}
	return maxJoltage
}
