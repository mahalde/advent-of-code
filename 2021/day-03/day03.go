package main

import (
	"fmt"
	"strconv"

	"github.com/mahalde/advent-of-code/utils/files"
)

func main() {
	input := files.ReadFile(3, 2021, "\n")
	fmt.Println("Solution Part One:", solvePart1(input))
	fmt.Println("Solution Part Two:", solvePart2(input))
}

func solvePart1(input []string) int {
	gamma, epsilon := calculateGammaAndEpsilonRate(input)

	result := parseStringByteToInt(gamma) * parseStringByteToInt(epsilon)
	return int(result)
}

func solvePart2(input []string) int {
	oxygenGenerator, co2Scrubber := calculateOxygenAndCO2(input)

	result := parseStringByteToInt(oxygenGenerator) * parseStringByteToInt(co2Scrubber)

	return int(result)
}

const ASCII_SHIFT_VALUE = 48

func calculateGammaAndEpsilonRate(input []string) (gamma, epsilon string) {
	for i := range input[0] {
		mostCommonBit, leastCommonBit := getCommonBits(input, i, 0)

		gamma += fmt.Sprint(mostCommonBit)
		epsilon += fmt.Sprint(leastCommonBit)
	}

	return
}

func getCommonBits(dataset []string, index int, valueIfEqual int) (most, least int) {
	bitArr := make([]int, 2)
	for _, line := range dataset {
		bitArr[line[index]-ASCII_SHIFT_VALUE]++
	}

	if bitArr[0] > bitArr[1] {
		return 0, 1
	} else if bitArr[1] > bitArr[0] {
		return 1, 0
	}

	return valueIfEqual, valueIfEqual
}

func parseStringByteToInt(str string) int64 {
	result, err := strconv.ParseInt(str, 2, 64)

	if err != nil {
		panic(err)
	}

	return result
}

const (
	MOST_COMMON_BIT_STR  = "MOST"
	LEAST_COMMON_BIT_STR = "LEAST"
)

func calculateOxygenAndCO2(input []string) (oxygenGenerator, co2Scrubber string) {
	oxygenGenerator = solveForPart2(input, 0, MOST_COMMON_BIT_STR, 1)
	co2Scrubber = solveForPart2(input, 0, LEAST_COMMON_BIT_STR, 0)

	return
}

func solveForPart2(dataset []string, index int, bitToFilterStr string, valueIfEqual int) string {
	if len(dataset) == 1 {
		return dataset[0]
	}

	mostCommonBit, leastCommonBit := getCommonBits(dataset, index, valueIfEqual)

	var bitToFilter int
	if bitToFilterStr == MOST_COMMON_BIT_STR {
		bitToFilter = mostCommonBit
	} else if bitToFilterStr == LEAST_COMMON_BIT_STR {
		bitToFilter = leastCommonBit
	}

	var filteredDataset []string

	for _, line := range dataset {
		if line[index]-ASCII_SHIFT_VALUE == byte(bitToFilter) {
			filteredDataset = append(filteredDataset, line)
		}
	}

	return solveForPart2(filteredDataset, index+1, bitToFilterStr, valueIfEqual)
}
