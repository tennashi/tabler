package service

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestTaskService(t *testing.T) {
	t.Run("CreateTaskFromInput", func(t *testing.T) {
		t.Run("should return error for empty title", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			testDir := filepath.Join(tmpDir, "empty_title_test")
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

			// Test cases for empty titles
			emptyInputs := []string{
				"",                  // completely empty
				"   ",               // only spaces
				"#tag",              // only tag
				"@tomorrow",         // only deadline
				"!!",                // only priority
				"#tag @tomorrow !!", // only shortcuts, no title
			}

			for _, input := range emptyInputs {
				// Act
				_, err := service.CreateTaskFromInput(input)

				// Assert
				if err == nil {
					t.Errorf("expected error for input %q, but got none", input)
				}
			}
		})

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

			secondID, err := service.CreateTaskFromInput("Second task #personal")
			if err != nil {
				t.Fatalf("failed to create second task: %v", err)
			}

			// Act
			taskItems, err := service.ListTasks(nil)
			// Assert
			if err != nil {
				t.Errorf("ListTasks() returned error: %v", err)
			}

			if len(taskItems) != 2 {
				t.Errorf("expected 2 tasks, got %d", len(taskItems))
			}

			// Order-independent verification using map
			tasksByID := make(map[string]*TaskItem)
			for _, item := range taskItems {
				tasksByID[item.Task.ID] = item
			}

			// Check first task
			firstTask, ok := tasksByID[firstID]
			if !ok {
				t.Errorf("first task (ID: %s) not found in results", firstID)
			} else {
				if firstTask.Task.Title != "First task" {
					t.Errorf("expected first task title %q, got %q", "First task", firstTask.Task.Title)
				}
				if len(firstTask.Tags) != 1 || firstTask.Tags[0] != "work" {
					t.Errorf("expected first task tags [work], got %v", firstTask.Tags)
				}
			}

			// Check second task
			secondTask, ok := tasksByID[secondID]
			if !ok {
				t.Errorf("second task (ID: %s) not found in results", secondID)
			} else {
				if secondTask.Task.Title != "Second task" {
					t.Errorf("expected second task title %q, got %q", "Second task", secondTask.Task.Title)
				}
				if len(secondTask.Tags) != 1 || secondTask.Tags[0] != "personal" {
					t.Errorf("expected second task tags [personal], got %v", secondTask.Tags)
				}
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
		t.Run("should return error for empty title", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			testDir := filepath.Join(tmpDir, "update_empty_test")
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
			taskID, err := service.CreateTaskFromInput("Original task #tag")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Test cases for empty titles
			emptyInputs := []string{
				"",          // completely empty
				"   ",       // only spaces
				"#newtag",   // only tag
				"@tomorrow", // only deadline
				"!!!",       // only priority
			}

			for _, input := range emptyInputs {
				// Act
				err := service.UpdateTaskFromInput(taskID, input)

				// Assert
				if err == nil {
					t.Errorf("expected error for input %q, but got none", input)
				}
			}
		})

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

	t.Run("ListTasks", func(t *testing.T) {
		t.Run("with tag filter should return tasks with specific tag", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			testDir := filepath.Join(tmpDir, "list_tag_filter_test")
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

			// Create tasks with different tags
			_, err = service.CreateTaskFromInput("Task 1 #work")
			if err != nil {
				t.Fatalf("failed to create task 1: %v", err)
			}
			_, err = service.CreateTaskFromInput("Task 2 #personal")
			if err != nil {
				t.Fatalf("failed to create task 2: %v", err)
			}
			_, err = service.CreateTaskFromInput("Task 3 #work #urgent")
			if err != nil {
				t.Fatalf("failed to create task 3: %v", err)
			}

			// Act
			filter := &FilterOptions{Tag: "work"}
			tasks, err := service.ListTasks(filter)
			// Assert
			if err != nil {
				t.Fatalf("ListTasks() returned error: %v", err)
			}

			if len(tasks) != 2 {
				t.Errorf("expected 2 tasks with tag 'work', got %d", len(tasks))
			}

			// Verify both tasks have 'work' tag
			for _, task := range tasks {
				hasWorkTag := false
				for _, tag := range task.Tags {
					if tag == "work" {
						hasWorkTag = true
						break
					}
				}
				if !hasWorkTag {
					t.Errorf("task %q doesn't have 'work' tag", task.Task.Title)
				}
			}
		})
	})
}
