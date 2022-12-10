package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils"

	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(6, 2022, "\n")[0]
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input string) int {
	return solve(input, 4)
}

func solvePart2(input string) int {
	return solve(input, 14)
}

func solve(input string, distinctLength int) int {
	for i := distinctLength; i < len(input); i++ {
		isUnique := utils.Unique(input[i-distinctLength : i])

		if isUnique {
			return i
		}
	}

	panic("No unique values found")
}
