package task

import (
	"testing"
	"time"
)

func TestTask(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		t.Run("with valid fields should create task", func(t *testing.T) {
			// Arrange
			id := "test-id-123"
			title := "Buy groceries"
			deadline := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
			priority := 2

			// Act
			task := NewTask(id, title, deadline, priority)

			// Assert
			if task.ID != id {
				t.Errorf("expected ID %q, got %q", id, task.ID)
			}
			if task.Title != title {
				t.Errorf("expected title %q, got %q", title, task.Title)
			}
			if !task.Deadline.Equal(deadline) {
				t.Errorf("expected deadline %v, got %v", deadline, task.Deadline)
			}
			if task.Priority != priority {
				t.Errorf("expected priority %d, got %d", priority, task.Priority)
			}
			if task.Completed {
				t.Error("expected completed to be false")
			}
			if task.CreatedAt.IsZero() {
				t.Error("expected CreatedAt to be set")
			}
			if task.UpdatedAt.IsZero() {
				t.Error("expected UpdatedAt to be set")
			}
		})
	})
}
