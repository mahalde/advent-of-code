package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type Bingo struct {
	Lines [5]*Line
}

type Line struct {
	Numbers [5]*BingoNumber
}

type BingoNumber struct {
	Value    int
	IsMarked bool
}

func (bingo *Bingo) getSumOfUnmarkedNumbers() int {
	sum := 0
	for _, line := range bingo.Lines {
		for _, num := range line.Numbers {
			if !num.IsMarked {
				sum += num.Value
			}
		}
	}

	return sum
}

func (bingo *Bingo) markNumber(number int) {
	for _, line := range bingo.Lines {
		line.markNumber(number)
	}
}

func (line *Line) markNumber(number int) {
	for _, bingoNumber := range line.Numbers {
		if bingoNumber.Value == number {
			bingoNumber.IsMarked = true
		}
	}
}

func (bingo *Bingo) isFinished() bool {
	finished := false
	for _, line := range bingo.Lines {
		finished = finished || line.isFinished()
	}
	for i := 0; i < 5; i++ {
		finished = finished || (bingo.Lines[0].Numbers[i].IsMarked &&
			bingo.Lines[1].Numbers[i].IsMarked &&
			bingo.Lines[2].Numbers[i].IsMarked &&
			bingo.Lines[3].Numbers[i].IsMarked &&
			bingo.Lines[4].Numbers[i].IsMarked)
	}

	return finished
}

func (line *Line) isFinished() bool {
	finished := true
	for _, number := range line.Numbers {
		finished = finished && number.IsMarked
	}

	return finished
}

func main() {
	input := files.ReadFile(4, 2021, "\n")
	fmt.Println("Solution Part One:", solvePart1(input))
	fmt.Println("Solution Part Two:", solvePart2(input))
}

func solvePart1(input []string) int {
	chosenNumbers, bingoFields := parseInput(input)

	for _, num := range chosenNumbers {
		for _, bingo := range bingoFields {
			bingo.markNumber(num)

			if bingo.isFinished() {
				return bingo.getSumOfUnmarkedNumbers() * num
			}
		}
	}

	return 0
}

func parseInput(input []string) (chosenNumbers []int, bingoFields []*Bingo) {
	for _, number := range strings.Split(input[0], ",") {
		if num, err := strconv.Atoi(number); err == nil {
			chosenNumbers = append(chosenNumbers, num)
		}
	}

	for i := 2; i < len(input); {
		if input[i] == "" {
			i++
			continue
		}

		bingo := &Bingo{}

		for j := 0; j < 5; j++ {
			line := &Line{}
			bingoNumbersStr := strings.Fields(input[i])

			for k, numberStr := range bingoNumbersStr {
				if num, err := strconv.Atoi(numberStr); err == nil {
					line.Numbers[k] = &BingoNumber{num, false}
				}
			}
			i++
			bingo.Lines[j] = line
		}
		bingoFields = append(bingoFields, bingo)
	}

	return
}

func solvePart2(input []string) int {
	chosenNumbers, bingoFields := parseInput(input)

	nonFinishedFields := bingoFields

	for {
		lastBingo := false
		for _, num := range chosenNumbers {
			var tempArr []*Bingo
			for _, bingo := range nonFinishedFields {
				bingo.markNumber(num)

				if lastBingo && bingo.isFinished() {
					return bingo.getSumOfUnmarkedNumbers() * num
				}

				if !bingo.isFinished() {
					tempArr = append(tempArr, bingo)
				}
			}

			if len(tempArr) == 1 {
				lastBingo = true
			}

			nonFinishedFields = tempArr
		}
	}
}
