package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
	"io"
	"os"
)

func main() {
	input := files.ReadFile(10, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Println("Solution Part Two:")
	solvePart2(input, os.Stdout)
}

func solvePart1(input []string) int {
	result, cycle, register := 0, 0, 1

	for _, line := range input {
		if cycle >= 220 {
			break
		}

		cycle++

		if cycle%40 == 20 {
			result += cycle * register
		}

		if line[:4] == "addx" {
			cycle++

			if cycle%40 == 20 {
				result += cycle * register
			}

			register += conv.ToInt(line[5:])
		}
	}

	fmt.Println(cycle)
	return result
}

func solvePart2(input []string, writer io.Writer) {
	cycle, register := 0, 1

	for _, line := range input {
		printToCRT(writer, cycle, register)
		cycle++

		if line[:4] == "addx" {
			printToCRT(writer, cycle, register)
			cycle++
			register += conv.ToInt(line[5:])
		}
	}
}

func printToCRT(w io.Writer, cycle, register int) {
	horizontal := cycle % 40
	if horizontal == register || horizontal == register-1 || horizontal == register+1 {
		fmt.Fprint(w, "#")
	} else {
		fmt.Fprint(w, ".")
	}

	if (cycle+1)%40 == 0 {
		fmt.Fprint(w, "\n")
	}
}
