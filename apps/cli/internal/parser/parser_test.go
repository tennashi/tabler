package parser

import (
	"testing"
	"time"
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

	t.Run("parse multiple tags from input", func(t *testing.T) {
		// Arrange
		input := "Review PR #work #urgent #frontend"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Review PR" {
			t.Errorf("expected title %q, got %q", "Review PR", result.Title)
		}
		if len(result.Tags) != 3 {
			t.Fatalf("expected 3 tags, got %d tags", len(result.Tags))
		}
		expectedTags := []string{"work", "urgent", "frontend"}
		for i, expectedTag := range expectedTags {
			if result.Tags[i] != expectedTag {
				t.Errorf("expected tag[%d] %q, got %q", i, expectedTag, result.Tags[i])
			}
		}
		if result.Priority != 0 {
			t.Errorf("expected priority 0, got %d", result.Priority)
		}
		if result.Deadline != nil {
			t.Errorf("expected no deadline, got %v", result.Deadline)
		}
	})

	t.Run("parse priority with single ! (priority 1)", func(t *testing.T) {
		// Arrange
		input := "Deploy to production !"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Deploy to production" {
			t.Errorf("expected title %q, got %q", "Deploy to production", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 1 {
			t.Errorf("expected priority 1, got %d", result.Priority)
		}
		if result.Deadline != nil {
			t.Errorf("expected no deadline, got %v", result.Deadline)
		}
	})

	t.Run("parse priority with double !! (priority 2)", func(t *testing.T) {
		// Arrange
		input := "Fix critical bug !!"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Fix critical bug" {
			t.Errorf("expected title %q, got %q", "Fix critical bug", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 2 {
			t.Errorf("expected priority 2, got %d", result.Priority)
		}
		if result.Deadline != nil {
			t.Errorf("expected no deadline, got %v", result.Deadline)
		}
	})

	t.Run("parse priority with triple !!! (priority 3)", func(t *testing.T) {
		// Arrange
		input := "Emergency hotfix !!!"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Emergency hotfix" {
			t.Errorf("expected title %q, got %q", "Emergency hotfix", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 3 {
			t.Errorf("expected priority 3, got %d", result.Priority)
		}
		if result.Deadline != nil {
			t.Errorf("expected no deadline, got %v", result.Deadline)
		}
	})

	t.Run("parse deadline with @today", func(t *testing.T) {
		// Arrange
		input := "Submit report @today"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Submit report" {
			t.Errorf("expected title %q, got %q", "Submit report", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 0 {
			t.Errorf("expected priority 0, got %d", result.Priority)
		}
		if result.Deadline == nil {
			t.Fatal("expected deadline, got nil")
		}
		// Check if deadline is today
		assertDateEquals(t, time.Now(), *result.Deadline, "today")
	})

	t.Run("parse deadline with @tomorrow", func(t *testing.T) {
		// Arrange
		input := "Prepare presentation @tomorrow"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Prepare presentation" {
			t.Errorf("expected title %q, got %q", "Prepare presentation", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 0 {
			t.Errorf("expected priority 0, got %d", result.Priority)
		}
		if result.Deadline == nil {
			t.Fatal("expected deadline, got nil")
		}
		// Check if deadline is tomorrow
		tomorrow := time.Now().AddDate(0, 0, 1)
		assertDateEquals(t, tomorrow, *result.Deadline, "tomorrow")
	})

	t.Run("parse deadline with weekday @mon", func(t *testing.T) {
		// Arrange
		input := "Team meeting @mon"

		// Act
		result := Parse(input)

		// Assert
		if result.Title != "Team meeting" {
			t.Errorf("expected title %q, got %q", "Team meeting", result.Title)
		}
		if len(result.Tags) != 0 {
			t.Errorf("expected no tags, got %v", result.Tags)
		}
		if result.Priority != 0 {
			t.Errorf("expected priority 0, got %d", result.Priority)
		}
		if result.Deadline == nil {
			t.Fatal("expected deadline, got nil")
		}
		// Check if deadline is next Monday
		nextMonday := getNextWeekday(time.Monday)
		assertDateEquals(t, nextMonday, *result.Deadline, "next Monday")
	})
}

// Helper functions for tests
func getNextWeekday(weekday time.Weekday) time.Time {
	now := time.Now()
	daysUntilWeekday := (int(weekday) - int(now.Weekday()) + 7) % 7
	if daysUntilWeekday == 0 {
		daysUntilWeekday = 7
	}
	return now.AddDate(0, 0, daysUntilWeekday)
}

func assertDateEquals(t *testing.T, expected, actual time.Time, description string) {
	t.Helper()
	if expected.Year() != actual.Year() ||
		expected.Month() != actual.Month() ||
		expected.Day() != actual.Day() {
		t.Errorf("expected %s (%s), got %s",
			description,
			expected.Format("2006-01-02"),
			actual.Format("2006-01-02"))
	}
}
