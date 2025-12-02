package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/mahalde/advent-of-code/conv"
	"github.com/mahalde/advent-of-code/files"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2024", "01"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2024", "01", "\n")

	solution1 := SolvePart1(input)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string) int {
	var (
		leftList  []int
		rightList []int
	)

	for _, line := range input {
		parts := strings.Split(line, "   ")
		leftList = append(leftList, conv.ToInt(parts[0]))
		rightList = append(rightList, conv.ToInt(parts[1]))
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	result := 0
	for i := 0; i < len(leftList); i++ {
		result += int(math.Abs(float64(leftList[i] - rightList[i])))
	}

	return result
}

func SolvePart2(input []string) int {
	var (
		list  []int
		times = make(map[int]int)
	)

	for _, line := range input {
		parts := strings.Split(line, "   ")
		list = append(list, conv.ToInt(parts[0]))
		times[conv.ToInt(parts[1])]++
	}

	result := 0
	for _, num := range list {
		result += num * times[num]
	}

	return result
}
