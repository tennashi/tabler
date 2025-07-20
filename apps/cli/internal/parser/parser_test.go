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

	t.Run("parse single tag with # prefix", func(t *testing.T) {
		// Arrange
		input := "Fix bug #work"
		
		// Act
		result := Parse(input)
		
		// Assert
		if result.Title != "Fix bug" {
			t.Errorf("expected title %q, got %q", "Fix bug", result.Title)
		}
		if len(result.Tags) != 1 {
			t.Fatalf("expected 1 tag, got %d tags", len(result.Tags))
		}
		if result.Tags[0] != "work" {
			t.Errorf("expected tag %q, got %q", "work", result.Tags[0])
		}
		if result.Priority != 0 {
			t.Errorf("expected priority 0, got %d", result.Priority)
		}
		if result.Deadline != nil {
			t.Errorf("expected no deadline, got %v", result.Deadline)
		}
	})
}