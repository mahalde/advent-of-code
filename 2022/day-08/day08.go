package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/conv"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(8, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0
	matrix := parseToMatrix(input)

	for rowIndex, row := range matrix {
		if rowIndex == 0 || rowIndex == len(matrix)-1 {
			result += len(row)
			continue
		}

		for colIndex, height := range row {
			if colIndex == 0 || colIndex == len(row)-1 {
				result++
				continue
			}

			if isVisible(matrix, colIndex, rowIndex, height) {
				result++
			}
		}
	}

	return result
}

func parseToMatrix(input []string) [][]int {
	var matrix [][]int

	for _, line := range input {
		matrix = append(matrix, conv.ToIntSlice(strings.Split(line, "")))
	}

	return matrix
}

func isVisible(matrix [][]int, colIndex, rowIndex, height int) bool {
	allBeforeHorizontal := matrix[rowIndex][:colIndex]
	allAfterHorizontal := matrix[rowIndex][colIndex+1:]
	var allBeforeVertical []int
	var allAfterVertical []int

	for i, row := range matrix {
		switch {
		case i < rowIndex:
			allBeforeVertical = append(allBeforeVertical, row[colIndex])
		case i > rowIndex:
			allAfterVertical = append(allAfterVertical, row[colIndex])
		}
	}

	return allSmaller(height, allBeforeHorizontal, allAfterHorizontal, allBeforeVertical, allAfterVertical)
}

func allSmaller(height int, slices ...[]int) bool {
	for _, slice := range slices {
		if utils.Max(slice) < height {
			return true
		}
	}

	return false
}

func solvePart2(input []string) int {
	var scenicScores []int

	matrix := parseToMatrix(input)

	for rowIndex, row := range matrix {
		if rowIndex == 0 || rowIndex == len(matrix)-1 {
			continue
		}

		for colIndex, height := range row {
			if colIndex == 0 || colIndex == len(row)-1 {
				continue
			}

			scenicScores = append(scenicScores, findScenicScore(matrix, colIndex, rowIndex, height))
		}
	}

	return utils.Max(scenicScores)
}

func findScenicScore(matrix [][]int, colIndex, rowIndex, height int) int {
	score := 1

	leftScore := 0
	for i := colIndex - 1; i >= 0; i-- {
		leftScore++
		if matrix[rowIndex][i] >= height {
			break
		}
	}
	score *= leftScore

	rightScore := 0
	for i := colIndex + 1; i < len(matrix[rowIndex]); i++ {
		rightScore++
		if matrix[rowIndex][i] >= height {
			break
		}
	}
	score *= rightScore

	topScore := 0
	for i := rowIndex - 1; i >= 0; i-- {
		topScore++
		if matrix[i][colIndex] >= height {
			break
		}
	}
	score *= topScore

	bottomScore := 0
	for i := rowIndex + 1; i < len(matrix); i++ {
		bottomScore++
		if matrix[i][colIndex] >= height {
			break
		}
	}
	score *= bottomScore

	return score
}
