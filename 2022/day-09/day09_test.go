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
	//go:embed testdata/largerdata
	largerFile string

	input       = files.ParseFile(file, "\n")
	largerInput = files.ParseFile(largerFile, "\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input)
	utils.AssertIntEquals(t, solution, 13)
}

func TestPart2(t *testing.T) {
	solution := solvePart2(largerInput)
	utils.AssertIntEquals(t, solution, 36)
}

func TestRope(t *testing.T) {
	rope := NewRope(2)

	head := rope.segments[0]
	tail := rope.segments[1]
	rope.Move(Instruction{'R', 4})
	AssertCoordinate(t, head, &Coordinate{4, 0})
	AssertCoordinate(t, tail, &Coordinate{3, 0})

	rope.Move(Instruction{'U', 4})
	AssertCoordinate(t, head, &Coordinate{4, 4})
	AssertCoordinate(t, tail, &Coordinate{4, 3})

	rope.Move(Instruction{'L', 3})
	AssertCoordinate(t, head, &Coordinate{1, 4})
	AssertCoordinate(t, tail, &Coordinate{2, 4})

	rope.Move(Instruction{'D', 1})
	AssertCoordinate(t, head, &Coordinate{1, 3})
	AssertCoordinate(t, tail, &Coordinate{2, 4})
}

func AssertCoordinate(t testing.TB, got, want *Coordinate) {
	t.Helper()
	if got.x != want.x {
		t.Errorf("x is not correct, got %v, want %v", got.x, want.x)
	}
	if got.y != want.y {
		t.Errorf("y is not correct, got %v, want %v", got.y, want.y)
	}
}
