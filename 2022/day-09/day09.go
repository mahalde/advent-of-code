package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
)

type Rope struct {
	segments      []*Coordinate
	head          *Coordinate
	tail          *Coordinate
	visitedFields map[string]bool
}

func NewRope(length int) *Rope {
	var segments []*Coordinate
	for i := 0; i < length; i++ {
		segments = append(segments, &Coordinate{})
	}
	return &Rope{
		segments:      segments,
		head:          &Coordinate{},
		tail:          &Coordinate{},
		visitedFields: map[string]bool{"0-0": true},
	}
}

func (r *Rope) Move(inst Instruction) {
	for i := 0; i < inst.amount; i++ {
		for segmentIndex := 0; segmentIndex < len(r.segments); segmentIndex++ {
			if segmentIndex == 0 {
				forceMove(r.segments[0], inst)
				continue
			}

			first, second := r.segments[segmentIndex-1], r.segments[segmentIndex]

			if isAdjacent(first, second) {
				break
			}

			r.moveToSegment(first, second)

			if segmentIndex == len(r.segments)-1 {
				r.visitedFields[second.String()] = true
			}
		}
	}
}

func forceMove(segment *Coordinate, inst Instruction) {
	switch inst.direction {
	case 'U':
		segment.y++
	case 'D':
		segment.y--
	case 'L':
		segment.x--
	case 'R':
		segment.x++
	}
}

func (r *Rope) moveToSegment(first, second *Coordinate) {
	if first.y > second.y {
		second.y++
	}
	if first.y < second.y {
		second.y--
	}
	if first.x > second.x {
		second.x++
	}
	if first.x < second.x {
		second.x--
	}
}

func isAdjacent(first, second *Coordinate) bool {
	return isVerticalAdjacent(first, second) && isHorizontalAdjacent(first, second)
}

func isVerticalAdjacent(first, second *Coordinate) bool {
	return first.y == second.y || first.y == second.y+1 || first.y == second.y-1
}

func isHorizontalAdjacent(first, second *Coordinate) bool {
	return first.x == second.x || first.x == second.x+1 || first.x == second.x-1
}

type Coordinate struct {
	x, y int
}

func (c *Coordinate) String() string {
	return fmt.Sprintf("%v-%v", c.x, c.y)
}

type Instruction struct {
	direction rune
	amount    int
}

func main() {
	input := files.ReadFile(9, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	rope := NewRope(2)
	for _, line := range input {
		rope.Move(parseInstruction(line))
	}

	return len(rope.visitedFields)
}

func parseInstruction(line string) Instruction {
	return Instruction{
		direction: rune(line[0]),
		amount:    conv.ToInt(line[2:]),
	}
}

func solvePart2(input []string) int {
	rope := NewRope(10)
	for _, line := range input {
		rope.Move(parseInstruction(line))
	}

	return len(rope.visitedFields)
}
