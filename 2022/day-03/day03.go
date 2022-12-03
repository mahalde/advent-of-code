package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(3, 2022, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for _, line := range input {
		firstHalf, secondHalf := line[:len(line)/2], line[len(line)/2:]

		extraItem, err := compareCompartments(firstHalf, secondHalf)

		if err != nil {
			panic(err)
		}

		result += getPriority(extraItem)
	}

	return result
}

func compareCompartments(first, second string) (rune, error) {
	for _, char := range first {
		if strings.ContainsRune(second, char) {
			return char, nil
		}
	}

	return 0, errors.New("no duplicate found")
}

func getPriority(item rune) int {
	if unicode.IsLower(item) {
		return int(item) - int('a') + 1
	} else {
		return int(item) - int('A') + 27
	}
}

func solvePart2(input []string) int {
	result := 0

	var group [3]string

	for i := 0; i < len(input); i++ {
		group[i%3] = input[i]

		if i%3 == 2 {
			groupBadge, err := getGroupBadge(group)

			if err != nil {
				panic(err)
			}

			result += getPriority(groupBadge)
		}
	}

	return result
}

func getGroupBadge(group [3]string) (rune, error) {
	for _, char := range group[0] {
		if strings.ContainsRune(group[1], char) && strings.ContainsRune(group[2], char) {
			return char, nil
		}
	}

	return 0, errors.New("no common character found")
}