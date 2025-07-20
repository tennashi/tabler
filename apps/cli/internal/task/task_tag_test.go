package task

import (
	"testing"
)

func TestTaskTag(t *testing.T) {
	t.Run("creation", func(t *testing.T) {
		t.Run("with valid fields should create task tag", func(t *testing.T) {
			// Arrange
			taskID := "task-123"
			tag := "work"

			// Act
			taskTag := NewTaskTag(taskID, tag)

			// Assert
			if taskTag.TaskID != taskID {
				t.Errorf("expected TaskID %q, got %q", taskID, taskTag.TaskID)
			}
			if taskTag.Tag != tag {
				t.Errorf("expected Tag %q, got %q", tag, taskTag.Tag)
			}
		})
	})
}
