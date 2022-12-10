package utils

import "testing"

func TestUnique(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"Hello", false},
		{"World", true},
		{"aAbBcC", true},
		{"whatisthishellofatestcase", false},
	}

	for _, test := range testCases {
		t.Run(test.str, func(t *testing.T) {
			got := Unique(test.str)

			if got != test.expected {
				t.Errorf("got %v, want %v", got, test.expected)
			}
		})
	}
}
