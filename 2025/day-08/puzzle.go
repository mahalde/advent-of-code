package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/mahalde/advent-of-code/datastructures"
	"github.com/mahalde/advent-of-code/files"
	"github.com/mahalde/advent-of-code/vectors"
	"golang.org/x/exp/maps"
)

func main() {
	if err := files.FetchPuzzleInputIfNotExists("2025", "08"); err != nil {
		panic(err)
	}
	input := files.ParsePuzzleInput("2025", "08", "\n")

	solution1 := SolvePart1(input, 1000)
	fmt.Printf("Part 1: %d\n", solution1)
	solution2 := SolvePart2(input)
	fmt.Printf("Part 2: %d\n", solution2)
}

func SolvePart1(input []string, totalConnections int) int {
	junctionBoxes, circuits := parseInput(input)

	for range totalConnections {
		box1, box2 := &JunctionBox{}, &JunctionBox{}
		var minDistance = math.MaxFloat64

		for _, b := range junctionBoxes {
			for _, otherB := range junctionBoxes {
				if b == otherB {
					continue
				}
				distance := b.EuclideanDistance(otherB.Vector)
				if distance < minDistance && !alreadyConnected(b, otherB) {
					minDistance = distance
					box1 = b
					box2 = otherB
				}
			}
		}

		connectBoxes(box1, box2, &circuits)
	}

	allCircuits := maps.Keys(circuits)
	slices.SortFunc(allCircuits, func(a, b *Circuit) int {
		return b.Size - a.Size
	})

	return allCircuits[0].Size * allCircuits[1].Size * allCircuits[2].Size
}

func SolvePart2(input []string) int {
	junctionBoxes, circuits := parseInput(input)
	box1, box2 := &JunctionBox{}, &JunctionBox{}
	
	for {
		var minDistance = math.MaxFloat64

		for _, b := range junctionBoxes {
			for _, otherB := range junctionBoxes {
				if b == otherB {
					continue
				}
				distance := b.EuclideanDistance(otherB.Vector)
				if distance < minDistance && !alreadyConnected(b, otherB) {
					minDistance = distance
					box1 = b
					box2 = otherB
				}
			}
		}

		connectBoxes(box1, box2, &circuits)
		if len(circuits) == 1 {
			break
		}
	}

	return box1.X * box2.X
}

func parseInput(input []string) ([]*JunctionBox, Circuits) {
	junctionBoxes := make([]*JunctionBox, len(input))
	circuits := make(Circuits)
	for i, line := range input {
		var x, y, z int
		_, err := fmt.Sscanf(line, "%d,%d,%d", &x, &y, &z)
		if err != nil {
			panic(err)
		}
		jb := &JunctionBox{
			Vector:      vectors.NewVector(x, y, z),
			connections: make(Connections),
		}
		c := &Circuit{}
		c.AppendBox(jb)

		circuits[c] = struct{}{}
		junctionBoxes[i] = jb
	}

	return junctionBoxes, circuits
}

type JunctionBox struct {
	*vectors.Vector
	connections Connections
	circuit     *Circuit
}

type Circuit struct {
	datastructures.LinkedList[*JunctionBox]
}

func (c *Circuit) AppendBox(box *JunctionBox) {
	c.Append(box)
	box.circuit = c
}

func (c *Circuit) AppendCircuit(other *Circuit) {
	for node := other.Head; node != nil; node = node.Next {
		c.AppendBox(node.Value)
	}
}

type (
	Circuits    map[*Circuit]struct{}
	Connections map[*JunctionBox]struct{}
)

func alreadyConnected(bo1, bo2 *JunctionBox) bool {
	_, connected := bo1.connections[bo2]
	return connected
}

func connectBoxes(box1, box2 *JunctionBox, circuits *Circuits) {
	box1.connections[box2] = struct{}{}
	box2.connections[box1] = struct{}{}

	if box1.circuit != box2.circuit {
		delete(*circuits, box2.circuit)
		box1.circuit.AppendCircuit(box2.circuit)
	}
}
