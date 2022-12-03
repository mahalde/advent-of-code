package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/files"
)

var part1Map = map[string]int{
	"A X": 3 + 1,
	"A Y": 6 + 2,
	"A Z": 0 + 3,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 6 + 1,
	"C Y": 0 + 2,
	"C Z": 3 + 3,
}

var part2Map = map[string]int{
	"A X": 0 + 3,
	"A Y": 3 + 1,
	"A Z": 6 + 2,
	"B X": 0 + 1,
	"B Y": 3 + 2,
	"B Z": 6 + 3,
	"C X": 0 + 2,
	"C Y": 3 + 3,
	"C Z": 6 + 1,
}

func main() {
	input := files.ReadFile(2, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	return solveWithMap(part1Map, input)
}

func solvePart2(input []string) int {
	return solveWithMap(part2Map, input)
}

func solveWithMap(pointMap map[string]int, input []string) int {
	points := 0
	for _, line := range input {
		points += pointMap[line]
	}
	return points
}
