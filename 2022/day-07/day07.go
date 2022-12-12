package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"sort"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type Child interface {
	Size() int
	Name() string
}

type Directory struct {
	name     string
	children []Child
	parent   *Directory
	size     int
}

func NewDirectory(name string) *Directory {
	return &Directory{name: name}
}

func (d *Directory) Size() int {
	if d.size != 0 {
		return d.size
	}

	size := 0
	for _, child := range d.children {
		size += child.Size()
	}
	d.size = size
	return size
}

func (d *Directory) Name() string {
	return d.name
}

func (d *Directory) AddChild(child Child) {
	if d.GetChild(child.Name()) != nil {
		return
	}

	d.children = append(d.children, child)

	if dir, ok := child.(*Directory); ok {
		dir.parent = d
	}
}

func (d *Directory) GetChild(name string) Child {
	for _, child := range d.children {
		if child.Name() == name {
			return child
		}
	}

	return nil
}

type File struct {
	name string
	size int
}

func (f *File) Size() int {
	return f.size
}

func (f *File) Name() string {
	return f.name
}

func main() {
	input := files.ReadFile(7, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	// ignore first line as it changes the cd to root
	input = input[1:]
	root := NewDirectory("/")
	cd := root

	for _, line := range input {
		cd = parseLine(line, cd)
	}

	return getSizesUnder10000(root)
}

func parseLine(line string, cd *Directory) *Directory {
	if line[0] == '$' {
		cmd := line[2:4]

		if cmd == "ls" {
			return cd
		}

		if cmd == "cd" {
			dir := line[5:]

			if dir == ".." {
				return cd.parent
			}

			return cd.GetChild(dir).(*Directory)
		}
	}

	if line[:3] == "dir" {
		dir := NewDirectory(line[4:])
		cd.AddChild(dir)
	} else {
		parts := strings.Split(line, " ")
		file := &File{
			name: parts[1],
			size: conv.ToInt(parts[0]),
		}
		cd.AddChild(file)
	}
	return cd
}

func getSizesUnder10000(dir *Directory) int {
	total := 0
	for _, child := range dir.children {
		childDir, ok := child.(*Directory)
		if !ok {
			continue
		}
		size := childDir.Size()
		if size <= 100_000 {
			total += size
		}
		total += getSizesUnder10000(childDir)
	}

	return total
}

func solvePart2(input []string) int {
	// ignore first line as it changes the cd to root
	input = input[1:]
	root := NewDirectory("/")
	cd := root

	for _, line := range input {
		cd = parseLine(line, cd)
	}

	remainingDifference := 30_000_000 - (70_000_000 - root.Size())
	allDirectorySizes := getAllDirectorySizes(root)
	sort.Ints(allDirectorySizes)

	for _, size := range allDirectorySizes {
		if size >= remainingDifference {
			return size
		}
	}

	return 0
}

func getAllDirectorySizes(dir *Directory) []int {
	sizes := []int{dir.Size()}
	for _, child := range dir.children {
		childDir, ok := child.(*Directory)
		if !ok {
			continue
		}
		childDirSizes := getAllDirectorySizes(childDir)
		sizes = append(sizes, childDirSizes...)
	}

	return sizes
}
