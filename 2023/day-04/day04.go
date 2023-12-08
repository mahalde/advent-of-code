package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"regexp"
	"slices"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(4, 2023, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

var cardNumRegEx = regexp.MustCompile("\\d+")

type card struct {
	number         int
	winningNumbers []int
	givenNumbers   []int
}

func solvePart1(input []string) int {
	result := 0

	cards := getCards(input)

	for _, card := range cards {
		points := 0
		for _, winningNumber := range card.winningNumbers {
			if isWin := slices.Contains(card.givenNumbers, winningNumber); !isWin {
				continue
			}

			if points == 0 {
				points = 1
			} else {
				points *= 2
			}
		}

		result += points
	}

	return result
}

func getCards(input []string) []*card {
	var cards []*card

	for _, line := range input {
		parts := strings.Split(line, ": ")
		numbers := strings.Split(parts[1], " | ")

		winningNumbers := parseNumbers(numbers[0])
		givenNumbers := parseNumbers(numbers[1])
		cardNumber := conv.ToInt(cardNumRegEx.FindString(parts[0]))

		cards = append(cards, &card{
			number:         cardNumber,
			winningNumbers: winningNumbers,
			givenNumbers:   givenNumbers,
		})
	}

	slices.SortFunc(cards, func(a, b *card) int {
		return a.number - b.number
	})

	return cards
}

func parseNumbers(numbers string) []int {
	rawValues := strings.Split(numbers, " ")
	rawValues = slices.DeleteFunc(rawValues, func(s string) bool {
		return s == ""
	})

	return conv.ToIntSlice(rawValues)
}

func solvePart2(input []string) int {
	result := 0
	cards := getCards(input)

	var newCards []*card
	originalCards := cards

	for len(cards) != 0 {
		for _, card := range cards {
			numOfWins := 0

			for _, winningNumber := range card.winningNumbers {
				if isWin := slices.Contains(card.givenNumbers, winningNumber); isWin {
					numOfWins++
				}
			}

			for i := card.number; i < card.number+numOfWins; i++ {
				newCards = append(newCards, originalCards[i])
			}
		}

		result += len(cards)
		cards = newCards
		newCards = make([]*card, 0)
	}

	return result
}
