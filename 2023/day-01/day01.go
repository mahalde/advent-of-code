package main

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils"
	"github.com/mahalde/advent-of-code/utils/conv"
	"github.com/mahalde/advent-of-code/utils/files"
	"regexp"
	"slices"
	"strings"
)

const numString = "0123456789"

var writtenNums = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

func main() {
	input := files.ReadFile(1, 2023, "\n")
	fmt.Printf("Solution Part One: %d\n", solvePart1(input))
	fmt.Printf("Solution Part Two: %d", solvePart2(input))
}

func solvePart1(input []string) int {
	var nums []int

	for _, line := range input {
		nums = append(nums, getNumFromLine(line))
	}

	return utils.Sum(nums)
}

func solvePart2(input []string) int {
	var nums []int
	for _, line := range input {
		hits := getHitsFromLine(line)
		firstNum, lastNum := getFirstAndLastFromHits(hits)

		nums = append(nums, conv.ToInt(firstNum+lastNum))
	}

	return utils.Sum(nums)
}

type hit struct {
	index int
	value string
}

func getHitsFromLine(line string) []hit {
	var hits []hit
	for key, value := range writtenNums {
		searchStr := fmt.Sprintf("(?:%v)|%v", key, value)
		matches := regexp.MustCompile(searchStr).FindAllStringIndex(line, -1)
		if matches == nil {
			continue
		}

		for _, match := range matches {
			hits = append(hits, hit{index: match[0], value: value})
		}
	}

	return hits
}

func getFirstAndLastFromHits(hits []hit) (string, string) {
	slices.SortFunc(hits, func(a, b hit) int {
		return a.index - b.index
	})
	firstNum := hits[0].value
	lastNum := hits[len(hits)-1].value

	return firstNum, lastNum
}

func getNumFromLine(line string) int {
	firstNum := string(line[strings.IndexAny(line, numString)])
	lastNum := string(line[strings.LastIndexAny(line, numString)])

	return conv.ToInt(firstNum + lastNum)
}
