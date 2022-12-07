package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/conv"
	"regexp"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type Stack struct {
	items []string
}

func (s *Stack) Len() int {
	return len(s.items)
}

func (s *Stack) Push(item string) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() string {
	lastItemIndex := len(s.items) - 1
	item := s.items[lastItemIndex]
	s.items = s.items[:lastItemIndex]

	return item
}

type MoveInstruction struct {
	from   int
	to     int
	amount int
}

func main() {
	input := files.ReadFile(5, 2022, "\n\n")
	fmt.Printf("Solution Part One: %q\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %q", solvePart2(input))
}

func solvePart1(input []string) string {
	return solve(input, applyInstructionsPart1)
}

func solvePart2(input []string) string {
	return solve(input, applyInstructionsPart2)
}

func solve(input []string, instructionApply func([]MoveInstruction, []*Stack)) string {
	result := ""
	stacks := parseCrateStacks(input[0])
	instructions := parseInstructions(input[1])

	instructionApply(instructions, stacks)

	for _, stack := range stacks {
		result += stack.Pop()
	}
	return result
}

func parseCrateStacks(input string) []*Stack {
	var stacks []*Stack

	lines := strings.Split(input, "\n")

	// ignore the last line, start with the containers
	for i := len(lines) - 2; i >= 0; i-- {
		crateStack, line := 0, lines[i]
		for {
			nextStack := crateStack*3 + crateStack

			if nextStack >= len(line) {
				break
			}

			crate := line[nextStack : nextStack+3]
			if crate == "   " {
				crateStack++
				continue
			}

			if len(stacks) == crateStack {
				stacks = append(stacks, &Stack{})
			}

			stacks[crateStack].Push(string(crate[1]))

			crateStack++
		}
	}

	return stacks
}

func parseInstructions(input string) []MoveInstruction {
	var instructions []MoveInstruction
	commandRegEx := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)

	lines := strings.Split(input, "\n")
	for _, line := range lines {
		matches := commandRegEx.FindStringSubmatch(line)[1:]
		instruction := MoveInstruction{
			from:   conv.ToInt(matches[1]) - 1,
			to:     conv.ToInt(matches[2]) - 1,
			amount: conv.ToInt(matches[0]),
		}

		instructions = append(instructions, instruction)
	}
	return instructions
}

func applyInstructionsPart1(instructions []MoveInstruction, stacks []*Stack) {
	for _, instruction := range instructions {
		for i := 0; i < instruction.amount; i++ {
			stacks[instruction.to].Push(stacks[instruction.from].Pop())
		}
	}
}

func applyInstructionsPart2(instructions []MoveInstruction, stacks []*Stack) {
	for _, instruction := range instructions {
		var cratesToMove []string
		for i := 0; i < instruction.amount; i++ {
			cratesToMove = append(cratesToMove, stacks[instruction.from].Pop())
		}
		utils.ReverseSlice(cratesToMove)
		for _, crate := range cratesToMove {
			stacks[instruction.to].Push(crate)
		}
	}
}
