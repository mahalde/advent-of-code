package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/conv"
	"regexp"
	"strings"

	"github.com/mahalde/advent-of-code/utils/files"
)

type game map[string]int

func newGame() game {
	game := game{}
	game["red"] = 0
	game["green"] = 0
	game["blue"] = 0

	return game
}
func (g game) isPossible() bool {
	return g["red"] <= red && g["green"] <= green && g["blue"] <= blue
}

const (
	red   = 12
	green = 13
	blue  = 14
)

var (
	gameIDRegEx = regexp.MustCompile("Game (\\d+)")
	turnRegEx   = regexp.MustCompile("(\\d+) (\\w+)")
)

func main() {
	input := files.ReadFile(2, 2023, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	result := 0

	for _, line := range input {
		parts := strings.Split(line, ":")
		gameID := getGameID(parts[0])
		turns := strings.Split(parts[1], ";")
		game := getGame(turns)

		if game.isPossible() {
			result += gameID
		}
	}

	return result
}

func getGameID(line string) int {
	id := gameIDRegEx.FindStringSubmatch(line)[1]

	return conv.ToInt(id)
}

func getGame(turns []string) game {
	game := newGame()
	for _, turn := range turns {
		parts := strings.Split(turn, ", ")
		for _, part := range parts {
			matches := turnRegEx.FindStringSubmatch(part)
			amount := conv.ToInt(matches[1])
			if game[matches[2]] < amount {
				game[matches[2]] = amount
			}
		}
	}

	return game
}

func solvePart2(input []string) int {
	result := 0

	for _, line := range input {
		parts := strings.Split(line, ":")
		turns := strings.Split(parts[1], ";")
		game := getGame(turns)

		result += game["red"] * game["green"] * game["blue"]
	}

	return result
}
