package files

import (
	"os"
	"strings"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func ReadFileAsString(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

func ReadAndParseFile(path, delimiter string) []string {
	content := ReadFileAsString(path)
	return ParseFile(content, delimiter)
}

// ParseFile parses the content of a file and splits it by the specified delimiter
func ParseFile(content, delimiter string) []string {
	slicedContent := strings.Split(content, delimiter)
	lastElement := slicedContent[len(slicedContent)-1]

	if lastElement == delimiter || lastElement == "" {
		return slicedContent[:len(slicedContent)-1]
	}

	slicedContent[len(slicedContent)-1] = strings.TrimRight(lastElement, delimiter+"\n")

	return slicedContent
}
