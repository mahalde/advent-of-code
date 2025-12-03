package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/mahalde/advent-of-code/assert"
	"github.com/mahalde/advent-of-code/files"
)

var (
	//go:embed test_input.txt
	testFile string
	//go:embed test_input_2.txt
	testFile2 string

	input  = files.ParseFile(testFile, "\n")
	input2 = files.ParseFile(testFile2, "\n")
)

func TestSolvePart1(t *testing.T) {
	solution := SolvePart1(strings.Join(input, ""))
	assert.Equals(t, solution, 161)
}

func TestSolvePart2(t *testing.T) {
	solution := SolvePart2(strings.Join(input2, ""))
	assert.Equals(t, solution, 48)
}
