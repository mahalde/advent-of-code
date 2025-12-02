package files

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

const (
	adventOfCodeURL             = "https://adventofcode.com"
	adventOfCodePuzzleInputPath = adventOfCodeURL + "/%s/day/%s/input"
)

func FetchPuzzleInputIfNotExists(year, day string) error {
	inputPath := fmt.Sprintf("%s/day-%s/input.txt", year, day)
	if Exists(inputPath) {
		return nil // exists
	}

	err := godotenv.Load()
	if err != nil {
		return fmt.Errorf("error loading .env: %w", err)
	}

	sessionToken := os.Getenv("SESSION_TOKEN")
	if sessionToken == "" {
		return fmt.Errorf("no SESSION_TOKEN in .env")
	}

	return FetchPuzzleInput(year, day, sessionToken)
}

func FetchPuzzleInput(year, day, sessionToken string) error {
	url := fmt.Sprintf(adventOfCodePuzzleInputPath, year, strings.TrimLeft(day, "0"))
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return fmt.Errorf("Error creating HTTP request: %w\n", err)
	}

	req.AddCookie(&http.Cookie{
		Name:  "session",
		Value: sessionToken,
	})

	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("Error making HTTP request: %w\n", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non-OK HTTP status: %s\n", resp.Status)
	}

	inputFile, err := os.Create(fmt.Sprintf("%s/day-%s/input.txt", year, day))
	if err != nil {
		return fmt.Errorf("Error creating input.txt file: %w\n", err)
	}
	defer inputFile.Close()

	_, err = inputFile.ReadFrom(resp.Body)
	if err != nil {
		return fmt.Errorf("Error writing to input.txt file: %w\n", err)
	}

	return nil
}

func ParsePuzzleInput(year, day, delimiter string) []string {
	inputPath := fmt.Sprintf("%s/day-%s/input.txt", year, day)
	return ReadAndParseFile(inputPath, delimiter)
}
