package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/joho/godotenv"
	"github.com/mahalde/advent-of-code/files"
)

type model struct {
	year  textinput.Model
	day   textinput.Model
	focus int // 0 for year, 1 for day
	done  bool
	err   error
}

func (m model) validate() error {
	year := m.year.Value()
	if len(year) < 4 {
		year = fmt.Sprintf("%04s", year)
	}
	t, err := time.Parse("2006", year)
	if err != nil {
		return err
	}
	if t.Year() < 2015 || t.Year() > time.Now().Year() {
		return fmt.Errorf("year must be between 2015 and %d", time.Now().Year())
	}

	day := m.day.Value()
	if len(day) < 2 {
		day = fmt.Sprintf("%02s", day)
	}

	if day == "00" {
		return fmt.Errorf("day must be between 1 and 25")
	}

	t, err = time.Parse("200602", m.year.Value()+day)
	if err != nil {
		return err
	}
	if time.Now().Month() != time.December && t.Year() == time.Now().Year() {
		return fmt.Errorf("cannot select days in December for the current year before December")
	}
	if t.Day() < 1 || t.Day() > 25 {
		return fmt.Errorf("day must be between 1 and 25")
	}
	if t.Year() == time.Now().Year() && t.Day() > time.Now().Day() {
		return fmt.Errorf("cannot select a day in the future")
	}
	return nil
}

func initialModel() model {
	yearInput := textinput.New()
	yearInput.Prompt = "Enter the year: "
	yearInput.Placeholder = time.Now().Format("2006")
	yearInput.Focus()
	yearInput.CharLimit = 4
	yearInput.Width = 20

	dayInput := textinput.New()
	dayInput.Prompt = "Enter the day: "
	dayInput.Placeholder = time.Now().Format("02")
	dayInput.CharLimit = 2
	dayInput.Width = 20

	return model{year: yearInput, day: dayInput, focus: 0}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		case tea.KeyEnter:
			if m.focus == 0 {
				m.focus = 1
				m.day.Focus()
				return m, textinput.Blink
			} else if m.year.Err == nil && m.day.Err == nil {
				m.done = true
				return m, tea.Quit
			}
		case tea.KeyTab:
			if m.focus == 0 {
				m.focus = 1
				m.year.Blur()
				m.day.Focus()
			} else {
				m.focus = 0
				m.day.Blur()
				m.year.Focus()
			}
			return m, textinput.Blink
		}
	}

	var cmd tea.Cmd
	if m.focus == 0 {
		m.year, cmd = m.year.Update(msg)
	} else {
		m.day, cmd = m.day.Update(msg)
	}

	m.err = m.validate()
	return m, cmd
}

func (m model) View() string {
	if m.done {
		return ""
	}
	var b strings.Builder
	b.WriteString(m.year.View())
	b.WriteString("\n\n")
	b.WriteString(m.day.View())
	b.WriteString("\n\n(esc to quit)")
	if m.err != nil {
		b.WriteString("\n\nError: " + m.err.Error())
	}
	return b.String()
}

func getSessionToken() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", fmt.Errorf("error loading .env file: %w", err)
	}
	return os.Getenv("SESSION_TOKEN"), nil
}

func main() {
	sessionToken, err := getSessionToken()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
	if sessionToken == "" {
		fmt.Println("No session token found in .env file. Please add it.")
		os.Exit(1)
	}

	p := tea.NewProgram(initialModel())
	m, err := p.Run()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	mod := m.(model)
	err = mod.validate()
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	year := mod.year.Value()
	day := fmt.Sprintf("%02s", mod.day.Value())

	err = createFolderStructure(year, day)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	err = fetchPuzzleInput(year, day, sessionToken)
	if err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}

	fmt.Println("Setup completed successfully!")
}

func createFolderStructure(year, day string) error {
	fmt.Printf("Creating folder structure for Year %s Day %s...\n", year, day)

	dayPath := fmt.Sprintf("%s/day-%s", year, day)
	err := os.MkdirAll(dayPath, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Error creating directories: %w\n", err)
	}

	// Use the templates in the template folder to create the files
	err = createFileFromTemplate(year, day, "puzzle.tmpl", fmt.Sprintf("%s/puzzle.go", dayPath))
	if err != nil {
		return fmt.Errorf("Error creating puzzle.go: %w\n", err)
	}
	err = createFileFromTemplate(year, day, "puzzle_test.tmpl", fmt.Sprintf("%s/puzzle_test.go", dayPath))
	if err != nil {
		return fmt.Errorf("Error creating puzzle_test.go: %w\n", err)
	}

	// Create empty test_input.txt
	testInputPath := fmt.Sprintf("%s/test_input.txt", dayPath)
	err = os.WriteFile(testInputPath, []byte(""), 0644)
	if err != nil {
		return fmt.Errorf("Error creating test_input.txt: %v\n", err)
	}

	return nil
}

func createFileFromTemplate(year, day, templateFile, outputFile string) (err error) {
	tmpl, err := template.New("templates").ParseFiles("templates/" + templateFile)
	if err != nil {
		return fmt.Errorf("Error parsing template: %w\n", err)
	}
	data := struct {
		Year string
		Day  string
	}{Year: year, Day: day}
	file, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("Error creating file: %v\n", err)
	}
	defer func() {
		if fileErr := file.Close(); fileErr != nil {
			err = fmt.Errorf("Error closing file: %w\n", fileErr)
		}
	}()
	err = tmpl.ExecuteTemplate(file, templateFile, data)
	if err != nil {
		return fmt.Errorf("Error executing template: %w\n", err)
	}

	return nil
}

func fetchPuzzleInput(year, day, sessionToken string) error {
	fmt.Printf("Fetching puzzle input for Year %s Day %s...\n", year, day)

	return files.FetchPuzzleInput(year, day, sessionToken)
}
