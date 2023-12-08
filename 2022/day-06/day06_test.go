package main

import (
	"github.com/mahalde/advent-of-code/utils/assert"
	"testing"
)

var testDataPart1 = []struct {
	input    string
	solution int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
	{"nppdvjthqldpwncqszvftbrmjlhg", 6},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
}

var testDataPart2 = []struct {
	input    string
	solution int
}{
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
	{"nppdvjthqldpwncqszvftbrmjlhg", 23},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
}

func TestPart1(t *testing.T) {
	for _, test := range testDataPart1 {
		t.Run(test.input, func(t *testing.T) {
			got := solvePart1(test.input)

			assert.IntEquals(t, got, test.solution)
		})
	}
}

func TestPart2(t *testing.T) {
	for _, test := range testDataPart2 {
		t.Run(test.input, func(t *testing.T) {
			got := solvePart2(test.input)

			assert.IntEquals(t, got, test.solution)
		})
	}
}
