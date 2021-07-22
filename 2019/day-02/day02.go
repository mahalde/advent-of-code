package main

import (
	"fmt"
	"strconv"

	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(2, 2019, ",")
	fmt.Println("Solution Part One: " + strconv.Itoa(solvePart1(input)))
	fmt.Println("Solution Part Two: " + strconv.Itoa(solvePart2(input)))
}

func solvePart1(input []string) int {
	program := conv.ToIntSlice(input)

	program[1] = 12
	program[2] = 2

	return runIntCode(program)
}

func runIntCode(program []int) int {
outer:
	for i := 0; i < len(program); {
		switch program[i] {
		case 1:
			program[program[i+3]] = program[program[i+1]] + program[program[i+2]]
		case 2:
			program[program[i+3]] = program[program[i+1]] * program[program[i+2]]
		case 99:
			break outer
		}
		i += 4
	}
	return program[0]
}

func solvePart2(input []string) int {

	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			program := conv.ToIntSlice(input)
			program[1] = noun
			program[2] = verb

			if runIntCode(program) == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return -1
}
