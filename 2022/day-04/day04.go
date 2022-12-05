package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(4, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for _, line := range input {
		first, second := readLine(line)

		if isInBounds(first, second) {
			result++
		}
	}
	return result
}

func readLine(line string) (first, second [2]int) {
	parts := strings.Split(line, ",")

	first = getBoundaries(parts[0])
	second = getBoundaries(parts[1])

	return
}

func getBoundaries(task string) [2]int {
	boundaries := strings.Split(task, "-")
	return [2]int{conv.ToInt(boundaries[0]), conv.ToInt(boundaries[1])}
}

func isInBounds(first, second [2]int) bool {
	return (first[0] >= second[0] && first[1] <= second[1]) || (second[0] >= first[0] && second[1] <= first[1])
}

func solvePart2(input []string) int {
	result := 0

	for _, line := range input {
		first, second := readLine(line)

		if isOverlap(first, second) {
			result++
		}
	}

	return result
}

func isOverlap(first, second [2]int) bool {
	return isInBounds(first, second) || (first[0] >= second[0] && first[0] <= second[1]) || (first[1] >= second[0] && first[1] <= second[1])
}
