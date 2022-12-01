package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
	"math"
	"sort"
)

func main() {
	input := files.ReadFile(1, 2022, "\n")
	fmt.Printf("Solution Part One: %v\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %v", solvePart2(input))
}

func solvePart1(input []string) int {
	return maxOfSlice(getAllCalories(input))
}

func maxOfSlice(slice []int) int {
	max := math.MinInt32

	for _, number := range slice {
		if number > max {
			max = number
		}
	}

	return max
}

func getAllCalories(input []string) []int {
	allCalories := make([]int, 0)

	calories := 0
	for _, line := range input {
		if line == "" {
			allCalories = append(allCalories, calories)
			calories = 0
		} else {
			calories += conv.ToInt(line)
		}
	}

	allCalories = append(allCalories, calories)

	return allCalories
}

func solvePart2(input []string) int {
	allCalories := getAllCalories(input)

	sort.Slice(allCalories, func(i, j int) bool {
		return allCalories[i] > allCalories[j]
	})

	topThreeCalories := allCalories[:3]

	sum := 0
	for _, calorie := range topThreeCalories {
		sum += calorie
	}

	return sum
}
