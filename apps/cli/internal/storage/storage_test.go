package storage

import (
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/tennashi/tabler/internal/task"
)

func TestStorage(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		t.Run("should create database file", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			dbPath := filepath.Join(tmpDir, "test.db")

			// Act
			storage, err := New(dbPath)
			if err != nil {
				t.Fatalf("failed to create storage: %v", err)
			}
			t.Cleanup(func() {
				if err := storage.Close(); err != nil {
					t.Errorf("failed to close storage: %v", err)
				}
			})

			err = storage.Init()
			// Assert
			if err != nil {
				t.Errorf("Init() returned error: %v", err)
			}

			// Check if database file exists
			if _, err := os.Stat(dbPath); os.IsNotExist(err) {
				t.Error("database file was not created")
			}
		})
	})

	t.Run("CreateTask", func(t *testing.T) {
		t.Run("should store task with tags", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			dbPath := filepath.Join(tmpDir, "test.db")

			storage, err := New(dbPath)
			if err != nil {
				t.Fatalf("failed to create storage: %v", err)
			}
			t.Cleanup(func() {
				if err := storage.Close(); err != nil {
					t.Errorf("failed to close storage: %v", err)
				}
			})

			if err := storage.Init(); err != nil {
				t.Fatalf("failed to init storage: %v", err)
			}

			task := &task.Task{
				ID:        "task-123",
				Title:     "Buy groceries",
				Deadline:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Priority:  2,
				Completed: false,
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
			}
			tags := []string{"shopping", "urgent"}

			// Act
			err = storage.CreateTask(task, tags)
			// Assert
			if err != nil {
				t.Errorf("CreateTask() returned error: %v", err)
			}
		})
	})

	t.Run("GetTask", func(t *testing.T) {
		t.Run("should retrieve task with tags", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			dbPath := filepath.Join(tmpDir, "test.db")

			storage, err := New(dbPath)
			if err != nil {
				t.Fatalf("failed to create storage: %v", err)
			}
			t.Cleanup(func() {
				if err := storage.Close(); err != nil {
					t.Errorf("failed to close storage: %v", err)
				}
			})

			if err := storage.Init(); err != nil {
				t.Fatalf("failed to init storage: %v", err)
			}

			// Create a task first
			originalTask := &task.Task{
				ID:        "task-456",
				Title:     "Read book",
				Deadline:  time.Date(2024, 2, 1, 0, 0, 0, 0, time.UTC),
				Priority:  1,
				Completed: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
			originalTags := []string{"personal", "reading"} // ORDER BY tag will sort alphabetically

			if err := storage.CreateTask(originalTask, originalTags); err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Act
			retrievedTask, retrievedTags, err := storage.GetTask("task-456")
			// Assert
			if err != nil {
				t.Errorf("GetTask() returned error: %v", err)
			}

			if retrievedTask.ID != originalTask.ID {
				t.Errorf("expected ID %q, got %q", originalTask.ID, retrievedTask.ID)
			}
			if retrievedTask.Title != originalTask.Title {
				t.Errorf("expected title %q, got %q", originalTask.Title, retrievedTask.Title)
			}
			if !retrievedTask.Deadline.Equal(originalTask.Deadline) {
				t.Errorf("expected deadline %v, got %v", originalTask.Deadline, retrievedTask.Deadline)
			}
			if retrievedTask.Priority != originalTask.Priority {
				t.Errorf("expected priority %d, got %d", originalTask.Priority, retrievedTask.Priority)
			}
			if retrievedTask.Completed != originalTask.Completed {
				t.Errorf("expected completed %v, got %v", originalTask.Completed, retrievedTask.Completed)
			}

			// Check tags
			if len(retrievedTags) != len(originalTags) {
				t.Errorf("expected %d tags, got %d", len(originalTags), len(retrievedTags))
			}
			for i, tag := range originalTags {
				if i < len(retrievedTags) && retrievedTags[i] != tag {
					t.Errorf("expected tag %q, got %q", tag, retrievedTags[i])
				}
			}
		})
	})

	t.Run("ListTasks", func(t *testing.T) {
		t.Run("should list all tasks when no filter", func(t *testing.T) {
			// Arrange
			tmpDir := t.TempDir()
			dbPath := filepath.Join(tmpDir, "test.db")

			storage, err := New(dbPath)
			if err != nil {
				t.Fatalf("failed to create storage: %v", err)
			}
			t.Cleanup(func() {
				if err := storage.Close(); err != nil {
					t.Errorf("failed to close storage: %v", err)
				}
			})

			if err := storage.Init(); err != nil {
				t.Fatalf("failed to init storage: %v", err)
			}

			// Create multiple tasks
			task1 := &task.Task{
				ID:        "task-111",
				Title:     "First task",
				Deadline:  time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
				Priority:  1,
				Completed: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
			if err := storage.CreateTask(task1, []string{"work"}); err != nil {
				t.Fatalf("failed to create task1: %v", err)
			}

			task2 := &task.Task{
				ID:        "task-222",
				Title:     "Second task",
				Deadline:  time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC),
				Priority:  2,
				Completed: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
			if err := storage.CreateTask(task2, []string{"personal"}); err != nil {
				t.Fatalf("failed to create task2: %v", err)
			}

			// Act
			tasks, err := storage.ListTasks(nil)
			// Assert
			if err != nil {
				t.Errorf("ListTasks() returned error: %v", err)
			}

			if len(tasks) != 2 {
				t.Errorf("expected 2 tasks, got %d", len(tasks))
			}
		})
	})
}
