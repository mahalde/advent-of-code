package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "02"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "02", ",")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	sum := 0
	for _, line := range input {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		start := conv.ToInt(parts[0])
		end := conv.ToInt(parts[1])
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			if len(s)%2 == 0 {
				half := len(s) / 2
				if s[:half] == s[half:] {
					sum += i
				}
			}
		}
	}
	return sum
}

func SolvePart2(input []string) int {
	sum := 0
	for _, line := range input {
		parts := strings.Split(line, "-")
		if len(parts) != 2 {
			continue
		}
		start := conv.ToInt(parts[0])
		end := conv.ToInt(parts[1])
		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			n := len(s)
			invalid := false
			for d := 1; d < n && !invalid; d++ {
				if n%d != 0 {
					continue
				}
				k := n / d
				if k < 2 {
					break
				}

				invalid = strings.Repeat(s[:d], k) == s

			}
			if invalid {
				sum += i
			}
		}
	}
	return sum
}
