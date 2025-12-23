package position_test

import (
	"testing"

	"github.com/mahalde/advent-of-code/position"
)

func TestDirection_TurnLeft(t *testing.T) {
	tests := []struct {
		name     string
		dir      position.Direction
		expected position.Direction
	}{
		{"North to West", position.North, position.West},
		{"East to North", position.East, position.North},
		{"South to East", position.South, position.East},
		{"West to South", position.West, position.South},
		{"Invalid direction", position.Direction("X"), position.Direction("X")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.TurnLeft()
			if result != tt.expected {
				t.Errorf("TurnLeft() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_TurnRight(t *testing.T) {
	tests := []struct {
		name     string
		dir      position.Direction
		expected position.Direction
	}{
		{"North to East", position.North, position.East},
		{"East to South", position.East, position.South},
		{"South to West", position.South, position.West},
		{"West to North", position.West, position.North},
		{"Invalid direction", position.Direction("X"), position.Direction("X")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.TurnRight()
			if result != tt.expected {
				t.Errorf("TurnRight() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_TurnAround(t *testing.T) {
	tests := []struct {
		name     string
		dir      position.Direction
		expected position.Direction
	}{
		{"North to South", position.North, position.South},
		{"East to West", position.East, position.West},
		{"South to North", position.South, position.North},
		{"West to East", position.West, position.East},
		{"Invalid direction", position.Direction("X"), position.Direction("X")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.TurnAround()
			if result != tt.expected {
				t.Errorf("TurnAround() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_String(t *testing.T) {
	tests := []struct {
		name     string
		dir      position.Direction
		expected string
	}{
		{"North", position.North, "N"},
		{"East", position.East, "E"},
		{"South", position.South, "S"},
		{"West", position.West, "W"},
		{"Invalid direction", position.Direction("X"), "X"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.dir.String()
			if result != tt.expected {
				t.Errorf("String() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDirection_Constants(t *testing.T) {
	if position.North != "N" {
		t.Errorf("North constant = %v, want N", position.North)
	}
	if position.East != "E" {
		t.Errorf("East constant = %v, want E", position.East)
	}
	if position.South != "S" {
		t.Errorf("South constant = %v, want S", position.South)
	}
	if position.West != "W" {
		t.Errorf("West constant = %v, want W", position.West)
	}
}
