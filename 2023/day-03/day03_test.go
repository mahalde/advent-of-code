package main

import (
	_ "embed"
	"github.com/mahalde/advent-of-code/utils/assert"
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

	assert.IntEquals(t, solution, 4361)
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)

	assert.IntEquals(t, solution, 467835)
}
