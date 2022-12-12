package main

import (
	"bytes"
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
	utils.AssertIntEquals(t, solution, 13140)
}

func TestPart2(t *testing.T) {
	got := &bytes.Buffer{}
	want :=
		`##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....
`

	solvePart2(input, got)
	if got.String() != want {
		t.Errorf("got:\n%v\nwant:\n%v", got.String(), want)
	}
}
