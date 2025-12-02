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

	input = files.ParseFile(testFile, ",")
)

func TestSolvePart1(t *testing.T) {
	solution := SolvePart1(input)
	assert.IntEquals(t, solution, 1227775554)
}

func TestSolvePart2(t *testing.T) {
	solution := SolvePart2(input)
	assert.IntEquals(t, solution, 4174379265)
}
