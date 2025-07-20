package parser

import (
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("parse task with just title", func(t *testing.T) {
		// Arrange
		input := "Buy groceries"
		
		// Act
		result := Parse(input)
		
		// Assert
		if result.Title != "Buy groceries" {
			t.Errorf("expected title %q, got %q", "Buy groceries", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 0 {
			t.Errorf("expected priority 0, got %d", result.Priority)
		}
		if result.Deadline != nil {
			t.Errorf("expected no deadline, got %v", result.Deadline)
		}
	})
}