package main

import (
	"strings"
	"testing"
	"time"

	"github.com/tennashi/tabler/internal/service"
	"github.com/tennashi/tabler/internal/task"
)

func TestFormatTaskDetails(t *testing.T) {
	t.Run("should format task details with proper labels and formatting", func(t *testing.T) {
		// Arrange
		created := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
		modified := time.Date(2024, 1, 15, 14, 45, 0, 0, time.UTC)
		deadline := time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC)
		
		task := &task.Task{
			ID:         "abc123",
			Title:      "Fix login bug",
			Priority:   3,
			Deadline:   deadline,
			Completed:  false,
			CreatedAt:  created,
			UpdatedAt:  modified,
		}
		tags := []string{"work", "urgent"}

		// Act
		result := formatTaskDetails(task, tags)

		// Assert
		expected := `ID: abc123
Task: Fix login bug
Status: Pending
Tags: work, urgent
Priority: High
Deadline: Jan 16, 2024
Created: Jan 15, 2024 10:30 AM
Modified: Jan 15, 2024 2:45 PM`

		if result != expected {
			t.Errorf("expected:\n%s\n\ngot:\n%s", expected, result)
		}
	})

	t.Run("should show Completed status when task is done", func(t *testing.T) {
		// Arrange
		created := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
		task := &task.Task{
			ID:        "abc123",
			Title:     "Fix login bug",
			Completed: true,
			CreatedAt: created,
		}
		tags := []string{}

		// Act
		result := formatTaskDetails(task, tags)

		// Assert
		if !strings.Contains(result, "Status: Completed") {
			t.Error("expected 'Status: Completed' for completed task")
		}
	})
}

func TestFormatTasksAsTable(t *testing.T) {
	t.Run("should format tasks with metadata in expanded format", func(t *testing.T) {
		// Arrange
		now := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
		deadline := time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC)
		taskItems := []*service.TaskItem{
			{
				Task: &task.Task{
					ID:        "abc123",
					Title:     "Fix login bug",
					Priority:  3,
					Deadline:  deadline,
					Completed: false,
					CreatedAt: now,
				},
				Tags: []string{"work", "urgent"},
			},
			{
				Task: &task.Task{
					ID:        "def456",
					Title:     "Review documentation",
					Priority:  1,
					Completed: true,
					CreatedAt: now,
				},
				Tags: []string{"docs"},
			},
			{
				Task: &task.Task{
					ID:        "ghi789",
					Title:     "Simple task without metadata",
					Completed: false,
					CreatedAt: now,
				},
				Tags: []string{},
			},
		}

		// Act
		result := formatTasksAsTable(taskItems)

		// Assert
		expected := `ID      Task                             Tags          Pri  Deadline     Status
------  -------------------------------  ------------  ---  -----------  ------
abc123  Fix login bug                    work, urgent  !!!  Jan 16       [ ]
def456  Review documentation             docs          !    -            [✓]
ghi789  Simple task without metadata     -             -    -            [ ]`

		if result != expected {
			t.Errorf("expected:\n%s\n\ngot:\n%s", expected, result)
		}
	})

	t.Run("should format tasks in compact format when no metadata", func(t *testing.T) {
		// Arrange
		now := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)
		taskItems := []*service.TaskItem{
			{
				Task: &task.Task{
					ID:        "abc123",
					Title:     "Fix login bug",
					Completed: false,
					CreatedAt: now,
				},
			},
			{
				Task: &task.Task{
					ID:        "def456",
					Title:     "Review documentation",
					Completed: true,
					CreatedAt: now,
				},
			},
		}

		// Act
		result := formatTasksAsTable(taskItems)

		// Assert
		expected := `ID      Task                    Status
---     --------------------    ------
abc123  Fix login bug           [ ]
def456  Review documentation    [✓]`

		if result != expected {
			t.Errorf("expected:\n%s\n\ngot:\n%s", expected, result)
		}
	})
}