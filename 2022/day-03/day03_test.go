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

	utils.AssertIntEquals(t, solution, 157)
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)

	utils.AssertIntEquals(t, solution, 70)
}
