package main

import (
	"fmt"
	"regexp"
	"slices"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/ranges"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "05"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "05", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	freshRanges, ingredients := parseInput(input)

	freshIngredients := 0
	for _, ing := range ingredients {
		for _, r := range freshRanges {
			if r.Contains(ing) {
				freshIngredients++
				break
			}
		}
	}

	return freshIngredients
}

func SolvePart2(input []string) int {
	freshRanges, _ := parseInput(input)

	slices.SortFunc(freshRanges, func(a, b *ranges.Range) int {
		return a.Start - b.Start
	})

	var merged []*ranges.Range
	current := freshRanges[0]
	for _, r := range freshRanges[1:] {
		if r.Start <= current.End {
			if r.End > current.End {
				current.End = r.End
			}
		} else {
			merged = append(merged, current)
			current = r
		}
	}
	merged = append(merged, current)

	totalFresh := 0
	for _, r := range merged {
		totalFresh += r.Size()
	}

	return totalFresh
}

var (
	ingredientRangeRegEx = regexp.MustCompile(`^(\d+)-(\d+)$`)
)

func parseInput(input []string) (freshRanges []*ranges.Range, ingredients []int) {
	for _, line := range input {
		match := ingredientRangeRegEx.FindStringSubmatch(line)
		if match != nil {
			start := conv.ToInt(match[1])
			end := conv.ToInt(match[2])

			freshRanges = append(freshRanges, ranges.NewRange(start, end))
		} else if line != "" {
			ingredients = append(ingredients, conv.ToInt(line))
		}
	}

	return
}
