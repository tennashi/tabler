package main

import (
	"testing"
	"time"

	"github.com/tennashi/tabler/internal/service"
	"github.com/tennashi/tabler/internal/task"
)

func TestFormatTasksAsTable(t *testing.T) {
	t.Run("should format tasks in table format with ID, title and status", func(t *testing.T) {
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