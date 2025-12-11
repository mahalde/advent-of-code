package main

import (
	"fmt"
	"regexp"
	"slices"
	"strings"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "06"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "06", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

type Problem struct {
	operands []int
	operator rune
}

func SolvePart1(input []string) int {
	problems := parseInputPart1(input)

	total := 0
	for _, problem := range problems {
		total += calculateTotal(problem.operands, problem.operator)
	}
	return total
}

func SolvePart2(input []string) int {
	problems := parseInputPart2(input)

	total := 0
	for _, problem := range problems {
		total += calculateTotal(problem.operands, problem.operator)
	}
	return total
}

var multiSpaceRegEx = regexp.MustCompile(`\s+`)

func parseInputPart1(input []string) []Problem {
	var problems []Problem
	for i, line := range input {
		line = multiSpaceRegEx.ReplaceAllString(line, " ")
		parts := strings.Split(line, " ")
		parts = slices.DeleteFunc(parts, func(s string) bool {
			return s == ""
		})
		if i == 0 {
			problems = make([]Problem, len(parts))
		}
		if parts[0] == "*" || parts[0] == "+" {
			for i, part := range parts {
				problems[i].operator = rune(part[0])
			}
		} else {
			for i, part := range parts {
				problems[i].operands = append(problems[i].operands, conv.ToInt(part))
			}
		}
	}
	return problems
}

func parseInputPart2(input []string) []Problem {
	lineLength := len(input[0])
	problems := make([]Problem, lineLength)
	i := 0
	for c := range lineLength {
		var operand, operator string
		for _, line := range input {
			if line[c] == ' ' {
				continue
			}
			if line[c] == '+' || line[c] == '*' {
				operator = string(line[c])
			} else {
				operand += string(line[c])
			}
		}
		if operand == "" {
			i++
		} else {
			problems[i].operands = append(problems[i].operands, conv.ToInt(operand))
		}

		if operator != "" {
			problems[i].operator = rune(operator[0])
		}
	}

	return problems
}

func calculateTotal(operands []int, operator rune) int {
	switch operator {
	case '+':
		sum := 0
		for _, operand := range operands {
			sum += operand
		}
		return sum
	case '*':
		product := 1
		for _, operand := range operands {
			product *= operand
		}
		return product
	default:
		return 0
	}
}
