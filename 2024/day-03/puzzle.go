package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2024", "03"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2024", "03", "\n")

	solution1 := SolvePart1(strings.Join(input, ""))
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(strings.Join(input, ""))
	fmt.Printf("Part 2: %d\n", solution2)
}

var mulRegEx = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
var mulDoDontRegEx = regexp.MustCompile("mul\\((\\d+),(\\d+)\\)|don't\\(\\)|do\\(\\)")

func SolvePart1(input string) int {
	matches := mulRegEx.FindAllStringSubmatch(input, -1)
	result := 0
	for _, match := range matches {
		mul1, mul2 := conv.ToInt(match[1]), conv.ToInt(match[2])
		result += mul1 * mul2
	}

	return result
}

func SolvePart2(input string) int {
	matches := mulDoDontRegEx.FindAllStringSubmatch(input, -1)
	result := 0
	enabled := true
	for _, match := range matches {
		if match[0] == "do()" {
			enabled = true
		} else if match[0] == "don't()" {
			enabled = false
		} else if enabled {
			mul1, mul2 := conv.ToInt(match[1]), conv.ToInt(match[2])
			result += mul1 * mul2
		}
	}

	return result
}
