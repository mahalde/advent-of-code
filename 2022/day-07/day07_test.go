package main

import (
	_ "embed"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/files"
	"reflect"
	"testing"
)

var (
	//go:embed testdata/data
	file string

	input = files.ParseFile(file, "\n")
)

func TestPart1(t *testing.T) {
	solution := solvePart1(input)
	utils.AssertIntEquals(t, solution, 95437)
}

func TestPart2(t *testing.T) {
	solution := solvePart2(input)
	utils.AssertIntEquals(t, solution, 24933642)
}

func TestDirectory(t *testing.T) {
	t.Run("gets a child from the directory", func(t *testing.T) {
		dir := NewDirectory("/")
		file := &File{name: "Hello World.txt"}
		dir.children = append(dir.children, file)

		got := dir.GetChild("Hello World.txt")

		if !reflect.DeepEqual(got, file) {
			t.Errorf("got %+v, want %+v", got, file)
		}
	})

	t.Run("adds a child to the directory", func(t *testing.T) {
		dir := NewDirectory("/")
		child := &File{name: "Hello World.txt"}

		got := dir.GetChild("Hello World.txt")

		if got != nil {
			t.Errorf("expected nil, got %+v", got)
		}

		dir.AddChild(child)
		got = dir.GetChild("Hello World.txt")

		if !reflect.DeepEqual(got, child) {
			t.Errorf("got %+v, want %+v", got, child)
		}
	})

	t.Run("doesn't add child if the name already exists", func(t *testing.T) {
		dir := NewDirectory("/")
		file := &File{name: "Hello World.txt"}
		dir.children = append(dir.children, file)

		utils.AssertIntEquals(t, len(dir.children), 1)

		dir.AddChild(file)

		utils.AssertIntEquals(t, len(dir.children), 1)
	})

	t.Run("gets the directory size", func(t *testing.T) {
		dir := NewDirectory("/")
		otherDir := NewDirectory("other")
		otherDir.AddChild(&File{size: 10})
		otherDir.AddChild(&File{name: "a", size: 25})
		dir.AddChild(&File{size: 15})
		dir.AddChild(otherDir)

		utils.AssertIntEquals(t, otherDir.Size(), 35)
		utils.AssertIntEquals(t, dir.Size(), 50)
	})

	t.Run("it sets the parent on adding a child dir", func(t *testing.T) {
		dir := NewDirectory("/")
		otherDir := NewDirectory("other")
		dir.AddChild(otherDir)

		if otherDir.parent != dir {
			t.Errorf("got %+v as parent, want %+v", otherDir.parent, dir)
		}
	})
}

func TestParseLine(t *testing.T) {
	t.Run("it cds to the parent", func(t *testing.T) {
		dir := NewDirectory("/")
		cd := NewDirectory("hello")
		dir.AddChild(cd)

		got := parseLine("$ cd ..", cd)

		if got != dir {
			t.Errorf("got %+v, want %+v", got, dir)
		}
	})

	t.Run("it cds to a child", func(t *testing.T) {
		dir := NewDirectory("/")
		cd := NewDirectory("hello")
		dir.AddChild(cd)

		got := parseLine("$ cd hello", dir)

		if got != cd {
			t.Errorf("got %+v, want %+v", got, cd)
		}
	})

	t.Run("it adds a directory", func(t *testing.T) {
		dir := NewDirectory("/")

		cd := parseLine("dir other", dir)

		if cd != dir {
			t.Fatalf("didn't expect the cd to change, but got %+v", cd)
		}

		child := dir.GetChild("other")

		if child == nil {
			t.Error("did not get a child")
		}
		if _, ok := child.(*Directory); !ok {
			t.Errorf("child is not a directory, but %t", child)
		}
	})

	t.Run("it adds a file", func(t *testing.T) {
		dir := NewDirectory("/")

		cd := parseLine("100 hello.txt", dir)

		if cd != dir {
			t.Fatalf("didn't expect the cd to change, but got %+v", cd)
		}

		child := dir.GetChild("hello.txt")

		if child == nil {
			t.Error("dit not get a child")
		}

		file, ok := child.(*File)
		if !ok {
			t.Errorf("child is not a file, but type %t", child)
		}

		utils.AssertIntEquals(t, file.Size(), 100)
		utils.AssertStringEquals(t, file.Name(), "hello.txt")
	})
}
