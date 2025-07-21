package parser

import (
	"testing"
	"time"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name             string
		input            string
		expectedTitle    string
		expectedTags     []string
		expectedPriority int
		hasDeadline      bool
		deadlineChecker  func(*testing.T, time.Time)
	}{
		{
			name:             "parse task with just title",
			input:            "Buy groceries",
			expectedTitle:    "Buy groceries",
			expectedTags:     []string{},
			expectedPriority: 0,
			hasDeadline:      false,
		},
		{
			name:             "parse single tag with # prefix",
			input:            "Fix bug #work",
			expectedTitle:    "Fix bug",
			expectedTags:     []string{"work"},
			expectedPriority: 0,
			hasDeadline:      false,
		},
		{
			name:             "parse multiple tags from input",
			input:            "Review PR #work #urgent #frontend",
			expectedTitle:    "Review PR",
			expectedTags:     []string{"work", "urgent", "frontend"},
			expectedPriority: 0,
			hasDeadline:      false,
		},
		{
			name:             "parse priority with single ! (priority 1)",
			input:            "Deploy to production !",
			expectedTitle:    "Deploy to production",
			expectedTags:     []string{},
			expectedPriority: 1,
			hasDeadline:      false,
		},
		{
			name:             "parse priority with double !! (priority 2)",
			input:            "Fix critical bug !!",
			expectedTitle:    "Fix critical bug",
			expectedTags:     []string{},
			expectedPriority: 2,
			hasDeadline:      false,
		},
		{
			name:             "parse priority with triple !!! (priority 3)",
			input:            "Emergency hotfix !!!",
			expectedTitle:    "Emergency hotfix",
			expectedTags:     []string{},
			expectedPriority: 3,
			hasDeadline:      false,
		},
		{
			name:             "parse deadline with @today",
			input:            "Submit report @today",
			expectedTitle:    "Submit report",
			expectedTags:     []string{},
			expectedPriority: 0,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				assertDateEquals(t, time.Now(), deadline, "today")
			},
		},
		{
			name:             "parse deadline with @tomorrow",
			input:            "Prepare presentation @tomorrow",
			expectedTitle:    "Prepare presentation",
			expectedTags:     []string{},
			expectedPriority: 0,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				tomorrow := time.Now().AddDate(0, 0, 1)
				assertDateEquals(t, tomorrow, deadline, "tomorrow")
			},
		},
		{
			name:             "parse deadline with weekday @mon",
			input:            "Team meeting @mon",
			expectedTitle:    "Team meeting",
			expectedTags:     []string{},
			expectedPriority: 0,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				nextMonday := getNextWeekday(time.Monday)
				assertDateEquals(t, nextMonday, deadline, "next Monday")
			},
		},
		{
			name:             "parse deadline with specific date @2024-01-15",
			input:            "Quarterly review @2024-01-15",
			expectedTitle:    "Quarterly review",
			expectedTags:     []string{},
			expectedPriority: 0,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				expectedDate := time.Date(2024, 1, 15, 0, 0, 0, 0, time.Local)
				assertDateEquals(t, expectedDate, deadline, "2024-01-15")
			},
		},
		{
			name:             "combine multiple shortcuts in one input",
			input:            "Urgent bug fix #backend #bug @today !!",
			expectedTitle:    "Urgent bug fix",
			expectedTags:     []string{"backend", "bug"},
			expectedPriority: 2,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				assertDateEquals(t, time.Now(), deadline, "today")
			},
		},
		{
			name:             "combine shortcuts in different order",
			input:            "!!! @tomorrow Review #code #important proposal",
			expectedTitle:    "Review proposal",
			expectedTags:     []string{"code", "important"},
			expectedPriority: 3,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				tomorrow := time.Now().AddDate(0, 0, 1)
				assertDateEquals(t, tomorrow, deadline, "tomorrow")
			},
		},
		{
			name:             "handle empty input",
			input:            "",
			expectedTitle:    "",
			expectedTags:     []string{},
			expectedPriority: 0,
			hasDeadline:      false,
		},
		{
			name:             "handle only shortcuts without title",
			input:            "#work #urgent @today !!",
			expectedTitle:    "",
			expectedTags:     []string{"work", "urgent"},
			expectedPriority: 2,
			hasDeadline:      true,
			deadlineChecker: func(t *testing.T, deadline time.Time) {
				assertDateEquals(t, time.Now(), deadline, "today")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange & Act
			result := Parse(tt.input)

			// Assert
			if result.Title != tt.expectedTitle {
				t.Errorf("expected title %q, got %q", tt.expectedTitle, result.Title)
			}

			if len(result.Tags) != len(tt.expectedTags) {
				t.Fatalf("expected %d tags, got %d tags", len(tt.expectedTags), len(result.Tags))
			}
			for i, expectedTag := range tt.expectedTags {
				if result.Tags[i] != expectedTag {
					t.Errorf("expected tag[%d] %q, got %q", i, expectedTag, result.Tags[i])
				}
			}

			if result.Priority != tt.expectedPriority {
				t.Errorf("expected priority %d, got %d", tt.expectedPriority, result.Priority)
			}

			if tt.hasDeadline {
				if result.Deadline == nil {
					t.Fatal("expected deadline, got nil")
				}
				if tt.deadlineChecker != nil {
					tt.deadlineChecker(t, *result.Deadline)
				}
			} else if result.Deadline != nil {
				t.Errorf("expected no deadline, got %v", result.Deadline)
			}
		})
	}
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
	t.Helper() // This ensures error line numbers point to the calling test, not this helper
	// Everything is in UTC, direct comparison
	if expected.Year() != actual.Year() ||
		expected.Month() != actual.Month() ||
		expected.Day() != actual.Day() {
		t.Errorf("expected %s (%s), got %s",
			description,
			expected.Format("2006-01-02"),
			actual.Format("2006-01-02"))
	}
}
