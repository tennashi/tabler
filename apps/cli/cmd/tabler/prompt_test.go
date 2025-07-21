package main

import (
	"strings"
	"testing"
)

func TestConfirmDeletion(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"should return true when user inputs y", "y\n", true},
		{"should return true when user inputs Y", "Y\n", true},
		{"should return false when user inputs n", "n\n", false},
		{"should return false when user inputs N", "N\n", false},
		{"should return false when user presses enter (default)", "\n", false},
		{"should return false for any other input", "maybe\n", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			taskTitle := "Fix login bug"
			reader := strings.NewReader(tt.input)

			// Act
			result := confirmDeletion(taskTitle, reader)

			// Assert
			if result != tt.expected {
				t.Errorf("expected %v for input %q, got %v", tt.expected, tt.input, result)
			}
		})
	}
}