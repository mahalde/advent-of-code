package main

import (
	"fmt"
	"strconv"

	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(1, 2021, "\n")
	fmt.Println("Solution Part One: " + strconv.Itoa(solvePart1(input)))
	fmt.Println("Solution Part Two: " + strconv.Itoa(solvePart2(input)))
}

func solvePart1(input []string) int {
	heights := conv.ToIntSlice(input)
	result := addHeights(heights, 1)
	return result
}

func solvePart2(input []string) int {
	heights := conv.ToIntSlice(input)
	result := addHeights(heights, 3)

	return result
}

func addHeights(heights []int, measurement int) (result int) {
	for i := 0; i < len(heights); {
		if i < measurement {
			i++
			continue
		}

		firstBatch := 0

		for j := i - measurement; j < i; j++ {
			if j < len(heights) {
				firstBatch += heights[j]
			}
		}

		i++
		secondBatch := 0

		for j := i - measurement; j < i; j++ {
			secondBatch += heights[j]
		}

		if firstBatch < secondBatch {
			result++
		}
	}

	return
}
