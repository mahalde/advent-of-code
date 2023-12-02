package main

import (
	_ "embed"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/files"
	"testing"
)

var (
	//go:embed testdata/data0
	file1 string

	//go:embed testdata/data1
	file2 string

	input1 = files.ParseFile(file1, "\n")
	input2 = files.ParseFile(file2, "\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input1)

	utils.AssertIntEquals(t, solution, 142)
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input2)

	utils.AssertIntEquals(t, solution, 281)
}
