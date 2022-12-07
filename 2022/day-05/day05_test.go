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

	input = files.ParseFile(file, "\n\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input)

	utils.AssertStringEquals(t, solution, "CMZ")
}

func TestStack(t *testing.T) {
	t.Run("has the correct length", func(t *testing.T) {
		stack := Stack{items: []string{"Hello", "World"}}

		utils.AssertIntEquals(t, stack.Len(), 2)
	})

	t.Run("correctly push items onto the stack", func(t *testing.T) {
		stack := Stack{items: []string{"Hello", "World"}}

		stack.Push("123")
		stack.Push("456")

		utils.AssertIntEquals(t, stack.Len(), 4)
	})

	t.Run("correctly pops items from the stack", func(t *testing.T) {
		stack := Stack{items: []string{"Hello", "World"}}

		item := stack.Pop()

		utils.AssertStringEquals(t, item, "World")
		utils.AssertIntEquals(t, stack.Len(), 1)

		item = stack.Pop()

		utils.AssertStringEquals(t, item, "Hello")
		utils.AssertIntEquals(t, stack.Len(), 0)
	})
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)
	utils.AssertStringEquals(t, solution, "MCD")
}
