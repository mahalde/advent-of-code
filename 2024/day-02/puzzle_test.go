package main

import (
	_ "embed"
	"testing"

	"github.com/mahalde/advent-of-code/assert"
	"github.com/mahalde/advent-of-code/files"
)

var (
	//go:embed test_input.txt
	testFile string

	input = files.ParseFile(testFile, "\n")
)

func TestSolvePart1(t *testing.T) {
	solution := SolvePart1(input)
	assert.Equals(t, solution, 2)
}

func TestSolvePart2(t *testing.T) {
	solution := SolvePart2(input)
	assert.Equals(t, solution, 4)
}
