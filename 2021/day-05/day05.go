package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type Line struct {
	start Coordinate
	end   Coordinate
}

type Coordinate struct {
	x, y int
}

func main() {
	input := files.ReadFile(5, 2021, "\n")
	fmt.Println("Solution Part One:", solve(input, false))
	fmt.Println("Solution Part Two:", solve(input, true))
}

func solve(input []string, diagonals bool) int {
	var field [1000][1000]int
	lines := parseInput(input)

	for _, line := range lines {
		markLineInField(&line, &field, diagonals)
	}

	result := 0

	for _, fieldLine := range field {
		for _, field := range fieldLine {
			if field >= 2 {
				result++
			}
		}
	}

	return result
}

func parseInput(input []string) (lines []Line) {
	for _, inputLine := range input {
		parts := strings.Split(inputLine, " -> ")
		beginParts := strings.Split(parts[0], ",")
		endParts := strings.Split(parts[1], ",")

		startCoordinate := Coordinate{convertToInt(beginParts[0]), convertToInt(beginParts[1])}
		endCoordinate := Coordinate{convertToInt(endParts[0]), convertToInt(endParts[1])}

		lines = append(lines, Line{startCoordinate, endCoordinate})
	}

	return
}

func markLineInField(line *Line, field *[1000][1000]int, diagonals bool) {

	if line.start.x != line.end.x && line.start.y != line.end.y {
		if !diagonals {
			return
		}
		for x, y := line.start.x, line.start.y; compareFunc(x, line.start.x, line.end.x) && compareFunc(y, line.start.y, line.end.y); func() { setVariableFunc(&x, line.start.x, line.end.x); setVariableFunc(&y, line.start.y, line.end.y) }() {
			field[int(y)][int(x)]++
		}
		return
	}

	for y := line.start.y; compareFunc(y, line.start.y, line.end.y); setVariableFunc(&y, line.start.y, line.end.y) {
		for x := line.start.x; compareFunc(x, line.start.x, line.end.x); setVariableFunc(&x, line.start.x, line.end.x) {
			field[int(y)][int(x)]++
		}
	}
}

func convertToInt(str string) int {
	num, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return num
}

func compareFunc(num, start, end int) bool {
	if start <= end {
		return num <= end
	} else {
		return num >= end
	}
}

func setVariableFunc(numPointer *int, start, end int) {
	if start <= end {
		*numPointer++
	} else {
		*numPointer--
	}
}
