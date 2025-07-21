package service

import (
	"testing"
	"time"
)

func TestTaskService(t *testing.T) {
	t.Run("CreateTaskFromInput", func(t *testing.T) {
		t.Run("should create task with parsed shortcuts", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			service, err := NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service.Close()
			}()

			input := "Fix bug in login #work #urgent @tomorrow !!"

			// Act
			taskID, err := service.CreateTaskFromInput(input)
			// Assert
			if err != nil {
				t.Errorf("CreateTaskFromInput() returned error: %v", err)
			}

			if taskID == "" {
				t.Error("expected non-empty task ID")
			}

			// Verify task was created correctly
			task, tags, err := service.GetTask(taskID)
			if err != nil {
				t.Fatalf("failed to get created task: %v", err)
			}

			if task.Title != "Fix bug in login" {
				t.Errorf("expected title %q, got %q", "Fix bug in login", task.Title)
			}

			if task.Priority != 2 {
				t.Errorf("expected priority 2, got %d", task.Priority)
			}

			// Check that deadline is tomorrow
			tomorrow := time.Now().AddDate(0, 0, 1)
			if task.Deadline.Day() != tomorrow.Day() ||
				task.Deadline.Month() != tomorrow.Month() ||
				task.Deadline.Year() != tomorrow.Year() {
				t.Errorf("expected deadline tomorrow, got %v", task.Deadline)
			}

			// Check tags
			expectedTags := []string{"urgent", "work"} // alphabetical order
			if len(tags) != len(expectedTags) {
				t.Errorf("expected %d tags, got %d", len(expectedTags), len(tags))
			}
			for i, tag := range expectedTags {
				if i < len(tags) && tags[i] != tag {
					t.Errorf("expected tag %q, got %q", tag, tags[i])
				}
			}
		})
	})
}
