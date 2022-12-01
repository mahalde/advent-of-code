package files

import (
	"fmt"
	"github.com/mahalde/advent-of-code/utils/req"
	"os"
	"strconv"
	"strings"
)

// ReadFile reads content of the input file and returns it in an array, split by the specified delimiter
// If the input file does not exist, it will be created
func ReadFile(day, year int, delimiter string) []string {
	currentDay := strconv.Itoa(day)
	currentYear := strconv.Itoa(year)

	if day < 10 {
		currentDay = "0" + currentDay
	}

	filePath := fmt.Sprintf("%v/day-%v/puzzle-input.in", currentYear, currentDay)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(year, day, filePath)
	} else {
		fmt.Println("INFO: File already exists, will not create new one")
	}

	return readFile(filePath, delimiter)
}

func readFile(filePath, delimiter string) []string {

	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	return ParseFile(fileContent, delimiter)
}

func ParseFile(content, delimiter string) []string {
	slicedContent := strings.Split(content, delimiter)

	if delimiter == "\n" {
		// fetch utils always adds a new line at the end of a file, which could lead to some problems when parsing it
		// this is why the last row is just removed if the delimiter is a newline

		return slicedContent[:len(slicedContent)-1]
	} else {
		// if the delimiter is not a newline and we split on eg. a comma, the newline will be appended to the last element
		// in the slice which then cannot be converted to an int
		// this is the reason the last element in the slice is modified (the last char is removed) so it can be worked with
		lastElement := slicedContent[len(slicedContent)-1]
		slicedContent[len(slicedContent)-1] = strings.TrimSuffix(lastElement, "\n")

		return slicedContent
	}
}

func createFile(year, day int, filePath string) {
	puzzleInput := req.GetPuzzleInput(year, day)

	err := os.WriteFile(filePath, []byte(puzzleInput), 0755)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("INFO: File successfully created")
	}
}
