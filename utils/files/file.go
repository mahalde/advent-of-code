package files

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/mahalde/advent-of-code/utils/req"
)

// Reads content of the input file and returns it in an array, split by the specified delimiter
// If the input file does not exist, it will be created
func ReadFile(day, year int, delimiter string) []string {
	currentDay := strconv.Itoa(day)
	currentYear := strconv.Itoa(year)

	if len(currentDay) == 1 {
		currentDay = "0" + currentDay
	}

	filePath := fmt.Sprintf("%v/day-%v/puzzle-input.in", currentYear, currentDay)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		createFile(year, day, filePath)
	} else {
		fmt.Println("INFO: File already exists, will not create new one")
	}

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, delimiter)

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

	err := ioutil.WriteFile(filePath, []byte(puzzleInput), 0755)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("INFO: File successfully created")
	}
}
