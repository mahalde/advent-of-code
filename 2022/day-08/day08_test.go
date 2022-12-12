package main

import (
	_ "embed"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/files"
	"testing"
)

var (
	//go:embed testdata/data
	file string

	input = files.ParseFile(file, "\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input)
	utils.AssertIntEquals(t, solution, 21)
}

func TestFindScenicScore(t *testing.T) {
	matrix := parseToMatrix(input)

	t.Run("first", func(t *testing.T) {
		got := findScenicScore(matrix, 2, 1, 5)
		utils.AssertIntEquals(t, got, 4)
	})

	t.Run("second", func(t *testing.T) {
		got := findScenicScore(matrix, 2, 3, 5)
		utils.AssertIntEquals(t, got, 8)
	})
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)
	utils.AssertIntEquals(t, solution, 8)
}
