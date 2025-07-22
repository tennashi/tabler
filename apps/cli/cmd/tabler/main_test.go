package main

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tennashi/tabler/internal/service"
)

// captureOutput captures stdout during function execution
func captureOutput(t *testing.T, fn func() error) (string, error) {
	t.Helper()

	origStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Ensure stdout is restored
	defer func() {
		os.Stdout = origStdout
	}()

	// Execute function
	err := fn()
	_ = w.Close()

	// Read captured output
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)

	return buf.String(), err
}

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

		t.Run("should filter tasks by tag with --tag flag", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			t.Setenv("TABLER_DATA_DIR", tmpDir)

			// Create tasks with different tags
			os.Args = []string{"tabler", "add", "Work task 1 #work"}
			_, err := captureOutput(t, run)
			if err != nil {
				t.Fatalf("failed to create work task 1: %v", err)
			}

			os.Args = []string{"tabler", "add", "Personal task #personal"}
			_, err = captureOutput(t, run)
			if err != nil {
				t.Fatalf("failed to create personal task: %v", err)
			}

			os.Args = []string{"tabler", "add", "Work task 2 #work #urgent"}
			_, err = captureOutput(t, run)
			if err != nil {
				t.Fatalf("failed to create work task 2: %v", err)
			}

			// Act - list with tag filter
			os.Args = []string{"tabler", "list", "--tag", "work"}
			output, err := captureOutput(t, run)
			// Assert
			if err != nil {
				t.Errorf("run() returned error: %v", err)
			}

			// Log output for debugging
			t.Logf("Captured output:\n%s", output)

			// Verify only work tasks are shown
			if !strings.Contains(output, "Work task 1") {
				t.Errorf("expected output to contain 'Work task 1', but it didn't")
			}
			if !strings.Contains(output, "Work task 2") {
				t.Errorf("expected output to contain 'Work task 2', but it didn't")
			}
			if strings.Contains(output, "Personal task") {
				t.Errorf("expected output NOT to contain 'Personal task', but it did")
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
