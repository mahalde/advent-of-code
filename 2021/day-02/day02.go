package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type submarine struct {
	position int
	depth    int
	aim      int
}

type instruction struct {
	command string
	units   int
}

func main() {
	input := files.ReadFile(2, 2021, "\n")
	fmt.Println("Solution Part One: " + strconv.Itoa(solvePart1(input)))
	fmt.Println("Solution Part Two: " + strconv.Itoa(solvePart2(input)))
}

func solvePart1(input []string) int {
	submarine := submarine{}

	processInstructions(&submarine, input, moveSubmarine)
	result := submarine.depth * submarine.position

	return result
}

func solvePart2(input []string) int {
	submarine := submarine{}

	processInstructions(&submarine, input, moveSubmarineWithAim)
	result := submarine.depth * submarine.position

	return result
}

func processInstructions(submarine *submarine, input []string, moveFunc func(*submarine, instruction)) {
	for _, line := range input {
		parts := strings.Split(line, " ")
		argument, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		instruction := instruction{parts[0], argument}

		moveFunc(submarine, instruction)
	}
}

func moveSubmarine(submarine *submarine, instruction instruction) {
	switch instruction.command {
	case "forward":
		submarine.position += instruction.units
	case "up":
		submarine.depth -= instruction.units
	case "down":
		submarine.depth += instruction.units
	}
}

func moveSubmarineWithAim(submarine *submarine, instruction instruction) {
	switch instruction.command {
	case "forward":
		submarine.position += instruction.units
		submarine.depth += submarine.aim * instruction.units
	case "up":
		submarine.aim -= instruction.units
	case "down":
		submarine.aim += instruction.units
	}
}
