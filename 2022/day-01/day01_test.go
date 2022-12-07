package main

import (
	_ "embed"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/files"
	"testing"
)

var (
	//go:embed testdata/data_0
	file string

	input = files.ParseFile(file, "\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input)

	utils.AssertIntEquals(t, solution, 24000)
}

func TestMaxOfSlice(t *testing.T) {
	slice := []int{2000, 1, 230, -50, 300, 5400, 5399}

	got := maxOfSlice(slice)

	utils.AssertIntEquals(t, got, 5400)
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)

	utils.AssertIntEquals(t, solution, 45000)
}
