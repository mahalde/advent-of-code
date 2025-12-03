package main

import (
	"fmt"
	"strings"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/ranges"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2024", "02"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2024", "02", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

var (
	rangeIncreasing = ranges.NewRange(1, 3)
	rangeDecreasing = ranges.NewRange(-3, -1)
)

func SolvePart1(input []string) int {
	safe := 0

	for _, report := range input {
		levels, safeRange := parseReport(report)

		if isSafeReport(levels, safeRange) {
			safe++
		}
	}

	return safe
}

func SolvePart2(input []string) int {
	safeLevels := 0
	for _, report := range input {
		levels, safeRange := parseReport(report)

		if isSafeReportWithDampener(levels, safeRange) {
			safeLevels++
		}
	}

	return safeLevels
}

func parseReport(report string) (levels []int, safeRange *ranges.Range) {
	levels = conv.ToIntSlice(strings.Split(report, " "))
	safeRange = getSafeRange(levels)
	return
}

func getSafeRange(levels []int) *ranges.Range {
	decreasing := levels[1]-levels[0] < 0
	if decreasing {
		return rangeDecreasing
	}
	return rangeIncreasing
}

func isSafeReport(levels []int, safeRange *ranges.Range) bool {
	for i := 1; i < len(levels); i++ {
		diff := levels[i] - levels[i-1]
		if !safeRange.Contains(diff) {
			return false
		}
	}
	return true
}

func isSafeReportWithDampener(levels []int, safeRange *ranges.Range) bool {
	if isSafeReport(levels, safeRange) {
		return true
	}

	for i := 0; i < len(levels); i++ {
		var newLevels []int
		newLevels = append(newLevels, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)

		newSafeRange := getSafeRange(newLevels)
		if isSafeReport(newLevels, newSafeRange) {
			return true
		}
	}

	return false
}
