package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(1, 2019, "\n")
	fmt.Println("Solution Part One: " + strconv.Itoa(solvePart1(input)))
	fmt.Println("Solution Part Two: " + strconv.Itoa(solvePart2(input)))
}

func solvePart1(input []string) int {
	result := 0

	intInput := conv.ToIntSlice(input)

	for _, mass := range intInput {
		result += calculateFuel(mass)
	}

	return result
}

func solvePart2(input []string) int {
	result := 0

	intInput := conv.ToIntSlice(input)

	for _, mass := range intInput {
		result += recursionCalculateFuel(mass)
	}

	return result
}

func calculateFuel(mass int) int {
	return int(math.Floor(float64(mass)/3)) - 2
}

func recursionCalculateFuel(mass int) int {
	fuel := calculateFuel(mass)
	if fuel > 0 {
		return fuel + recursionCalculateFuel(fuel)
	}
	return 0
}
