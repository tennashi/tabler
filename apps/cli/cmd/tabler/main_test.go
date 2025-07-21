package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tennashi/tabler/internal/service"
)

func TestCLI(t *testing.T) {
	t.Run("add command", func(t *testing.T) {
		t.Run("should create task from input", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()

			// Set up test arguments
			os.Args = []string{
				"tabler",
				"add",
				"Fix bug in login #work #urgent @tomorrow !!",
			}

			// Set data directory environment variable (automatically cleaned up)
			t.Setenv("TABLER_DATA_DIR", tmpDir)

			// Act
			err := run()
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}

			// Check that database file was created
			dbPath := filepath.Join(tmpDir, "tasks.db")
			if _, err := os.Stat(dbPath); os.IsNotExist(err) {
				t.Error("database file was not created")
			}
		})
	})

	t.Run("list command", func(t *testing.T) {
		t.Run("should list all tasks", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()

			// Set data directory environment variable
			t.Setenv("TABLER_DATA_DIR", tmpDir)

			// First create some tasks
			os.Args = []string{"tabler", "add", "First task #work"}
			if err := run(); err != nil {
				t.Fatalf("failed to create first task: %v", err)
			}

			os.Args = []string{"tabler", "add", "Second task #personal"}
			if err := run(); err != nil {
				t.Fatalf("failed to create second task: %v", err)
			}

			// Set up test arguments for list
			os.Args = []string{"tabler", "list"}

			// Act
			err := run()
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}
		})
	})

	t.Run("done command", func(t *testing.T) {
		t.Run("should complete task successfully", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			t.Setenv("TABLER_DATA_DIR", tmpDir)

			// Create a task directly using service to get real ID
			taskService, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			taskID, err := taskService.CreateTaskFromInput("Test task #work")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}
			_ = taskService.Close()

			// Set up test arguments for done with real task ID
			os.Args = []string{"tabler", "done", taskID}

			// Act
			err = run()
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}

			// Verify task is completed
			service2, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service2.Close()
			}()

			task, _, err := service2.GetTask(taskID)
			if err != nil {
				t.Fatalf("failed to get task: %v", err)
			}

			if !task.Completed {
				t.Error("expected task to be completed")
			}
		})
	})

	t.Run("show command", func(t *testing.T) {
		t.Run("should show task details", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			t.Setenv("TABLER_DATA_DIR", tmpDir)

			// Create a task directly to get real ID
			taskService, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			taskID, err := taskService.CreateTaskFromInput("Show test task #work #urgent @tomorrow !!")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}
			_ = taskService.Close()

			// Set up test arguments for show
			os.Args = []string{"tabler", "show", taskID}

			// Act
			err = run()
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}
		})
	})

	t.Run("delete command", func(t *testing.T) {
		t.Run("should delete task", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			t.Setenv("TABLER_DATA_DIR", tmpDir)
			t.Setenv("TABLER_NON_INTERACTIVE", "1") // Skip confirmation prompt in test

			// Create a task directly to get real ID
			taskService, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			taskID, err := taskService.CreateTaskFromInput("Delete test task #temp")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}
			_ = taskService.Close()

			// Set up test arguments for delete
			os.Args = []string{"tabler", "delete", taskID}

			// Act
			err = run()
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}

			// Verify task is deleted
			service2, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service2.Close()
			}()

			_, _, err = service2.GetTask(taskID)
			if err == nil {
				t.Error("expected error when getting deleted task")
			}
		})
	})

	t.Run("update command", func(t *testing.T) {
		t.Run("should update task", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			t.Setenv("TABLER_DATA_DIR", tmpDir)

			// Create a task directly to get real ID
			taskService, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			taskID, err := taskService.CreateTaskFromInput("Original task #old")
			if err != nil {
				t.Fatalf("failed to create task: %v", err)
			}
			_ = taskService.Close()

			// Set up test arguments for update
			os.Args = []string{"tabler", "update", taskID, "Updated task #new #updated !"}

			// Act
			err = run()
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}

			// Verify task is updated
			service2, err := service.NewTaskService(tmpDir)
			if err != nil {
				t.Fatalf("failed to create service: %v", err)
			}
			defer func() {
				_ = service2.Close()
			}()

			task, tags, err := service2.GetTask(taskID)
			if err != nil {
				t.Fatalf("failed to get task: %v", err)
			}

			if task.Title != "Updated task" {
				t.Errorf("expected title %q, got %q", "Updated task", task.Title)
			}

			expectedTags := []string{"new", "updated"}
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
