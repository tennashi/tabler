package service

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTaskService(t *testing.T) {
	t.Run("CreateTaskFromInput", func(t *testing.T) {
		t.Run("should create task with parsed shortcuts", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			// Use unique subdirectory for each test
			testDir := filepath.Join(tmpDir, "create_test")
			if err := os.MkdirAll(testDir, 0o750); err != nil {
				t.Fatalf("failed to create test directory: %v", err)
			}
			service, err := NewTaskService(testDir)
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

	t.Run("ListTasks", func(t *testing.T) {
		t.Run("should list all tasks with their tags", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			// Use unique subdirectory for each test
			testDir := filepath.Join(tmpDir, "list_test")
			if err := os.MkdirAll(testDir, 0o750); err != nil {
				t.Fatalf("failed to create test directory: %v", err)
			}
			service, err := NewTaskService(testDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service.Close()
			}()

			// Create multiple tasks
			firstID, err := service.CreateTaskFromInput("First task #work @today !")
			if err != nil {
				t.Fatalf("failed to create first task: %v", err)
			}

			time.Sleep(100 * time.Millisecond) // Ensure different creation times

			secondID, err := service.CreateTaskFromInput("Second task #personal")
			if err != nil {
				t.Fatalf("failed to create second task: %v", err)
			}

			// Act
			taskItems, err := service.ListTasks()
			// Assert
			if err != nil {
				t.Errorf("ListTasks() returned error: %v", err)
			}

			if len(taskItems) != 2 {
				t.Errorf("expected 2 tasks, got %d", len(taskItems))
			}

			// Check that both tasks exist (order should be second task first due to ORDER BY created_at DESC)
			if taskItems[0].Task.ID != secondID {
				t.Errorf("expected first item to be second task (ID: %s), got %s", secondID, taskItems[0].Task.ID)
			}
			if taskItems[0].Task.Title != "Second task" {
				t.Errorf("expected first task title %q, got %q", "Second task", taskItems[0].Task.Title)
			}
			if len(taskItems[0].Tags) != 1 || taskItems[0].Tags[0] != "personal" {
				t.Errorf("expected first task tags [personal], got %v", taskItems[0].Tags)
			}

			// Check second item in list (should be first task created)
			if taskItems[1].Task.ID != firstID {
				t.Errorf("expected second item to be first task (ID: %s), got %s", firstID, taskItems[1].Task.ID)
			}
			if taskItems[1].Task.Title != "First task" {
				t.Errorf("expected second task title %q, got %q", "First task", taskItems[1].Task.Title)
			}
			if len(taskItems[1].Tags) != 1 || taskItems[1].Tags[0] != "work" {
				t.Errorf("expected second task tags [work], got %v", taskItems[1].Tags)
			}
		})
	})

	t.Run("CompleteTask", func(t *testing.T) {
		t.Run("should mark task as completed", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			// Use unique subdirectory for each test
			testDir := filepath.Join(tmpDir, "complete_test")
			if err := os.MkdirAll(testDir, 0o750); err != nil {
				t.Fatalf("failed to create test directory: %v", err)
			}
			service, err := NewTaskService(testDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service.Close()
			}()

			// Create a task first
			taskID, err := service.CreateTaskFromInput("Test task #work")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Act
			err = service.CompleteTask(taskID)
			// Assert
			if err != nil {
				t.Errorf("CompleteTask() returned error: %v", err)
			}

			// Verify task is completed
			task, _, err := service.GetTask(taskID)
			if err != nil {
				t.Fatalf("failed to get task: %v", err)
			}

			if !task.Completed {
				t.Error("expected task to be completed")
			}
		})
	})

	t.Run("DeleteTask", func(t *testing.T) {
		t.Run("should delete task", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			// Use unique subdirectory for each test
			testDir := filepath.Join(tmpDir, "delete_test")
			if err := os.MkdirAll(testDir, 0o750); err != nil {
				t.Fatalf("failed to create test directory: %v", err)
			}
			service, err := NewTaskService(testDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service.Close()
			}()

			// Create a task first
			taskID, err := service.CreateTaskFromInput("Delete test task #temp")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Act
			err = service.DeleteTask(taskID)
			// Assert
			if err != nil {
				t.Errorf("DeleteTask() returned error: %v", err)
			}

			// Verify task is deleted
			_, _, err = service.GetTask(taskID)
			if err == nil {
				t.Error("expected error when getting deleted task")
			}
		})
	})

	t.Run("UpdateTask", func(t *testing.T) {
		t.Run("should update task from new input", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			// Use unique subdirectory for each test
			testDir := filepath.Join(tmpDir, "update_test")
			if err := os.MkdirAll(testDir, 0o750); err != nil {
				t.Fatalf("failed to create test directory: %v", err)
			}
			service, err := NewTaskService(testDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service.Close()
			}()

			// Create a task first
			taskID, err := service.CreateTaskFromInput("Old task #oldtag")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Act
			err = service.UpdateTaskFromInput(taskID, "Updated task #newtag #updated @tomorrow !!")
			// Assert
			if err != nil {
				t.Errorf("UpdateTaskFromInput() returned error: %v", err)
			}

			// Verify task is updated
			task, tags, err := service.GetTask(taskID)
			if err != nil {
				t.Fatalf("failed to get task: %v", err)
			}

			if task.Title != "Updated task" {
				t.Errorf("expected title %q, got %q", "Updated task", task.Title)
			}

			if task.Priority != 2 {
				t.Errorf("expected priority 2, got %d", task.Priority)
			}

			// Check tags (sorted alphabetically)
			expectedTags := []string{"newtag", "updated"}
			if len(tags) != len(expectedTags) {
				t.Fatalf("expected %d tags, got %d", len(expectedTags), len(tags))
			}
			for i, tag := range expectedTags {
				if tags[i] != tag {
					t.Errorf("expected tag %q, got %q", tag, tags[i])
				}
			}
		})
	})
}
