package main

import (
	"fmt"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "01"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "01", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	dial := 50
	timesZero := 0
	for _, line := range input {
		if line[0] == 'R' {
			dial += conv.ToInt(line[1:])
			dial %= 100
		} else {
			dial -= conv.ToInt(line[1:])
			if dial < 0 {
				dial = (100 - (dial * -1 % 100)) % 100
			}
		}

		if dial == 0 {
			timesZero++
		}
	}
	return timesZero
}

func SolvePart2(input []string) int {
	dial := 50
	timesZero := 0
	for _, line := range input {
		if line[0] == 'R' {
			dial += conv.ToInt(line[1:])
			timesZero += (dial - 1) / 100
			dial %= 100
		} else {
			wasZero := dial == 0
			amount := conv.ToInt(line[1:])
			timesZero += amount / 100
			amount %= 100
			dial -= amount
			if dial < 0 {
				dial = (100 - (dial * -1)) % 100
				if !wasZero {
					timesZero++
				}
			}
		}

		if dial == 0 {
			timesZero++
		}
	}
	return timesZero
}
