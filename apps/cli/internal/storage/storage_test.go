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

	t.Run("UpdateTask", func(t *testing.T) {
		t.Run("should update task completed status", func(t *testing.T) {
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
				ID:        "task-789",
				Title:     "Complete project",
				Deadline:  time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC),
				Priority:  2,
				Completed: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
			if err := storage.CreateTask(originalTask, []string{"work"}); err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Act
			err = storage.UpdateTaskCompleted("task-789", true)
			// Assert
			if err != nil {
				t.Errorf("UpdateTaskCompleted() returned error: %v", err)
			}

			// Verify the update
			updatedTask, _, err := storage.GetTask("task-789")
			if err != nil {
				t.Fatalf("failed to get updated task: %v", err)
			}

			if !updatedTask.Completed {
				t.Error("expected task to be completed")
			}
		})
	})

	t.Run("DeleteTask", func(t *testing.T) {
		t.Run("should delete task and its tags", func(t *testing.T) {
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
			task := &task.Task{
				ID:        "task-999",
				Title:     "Task to delete",
				Deadline:  time.Date(2024, 4, 1, 0, 0, 0, 0, time.UTC),
				Priority:  1,
				Completed: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
			if err := storage.CreateTask(task, []string{"temp", "delete"}); err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Act
			err = storage.DeleteTask("task-999")
			// Assert
			if err != nil {
				t.Errorf("DeleteTask() returned error: %v", err)
			}

			// Verify task is deleted
			_, _, err = storage.GetTask("task-999")
			if err == nil {
				t.Error("expected error when getting deleted task")
			}
		})
	})

	t.Run("UpdateTaskFull", func(t *testing.T) {
		t.Run("should update all task fields", func(t *testing.T) {
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
				ID:        "task-888",
				Title:     "Original title",
				Deadline:  time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC),
				Priority:  1,
				Completed: false,
				CreatedAt: time.Now().UTC(),
				UpdatedAt: time.Now().UTC(),
			}
			if err := storage.CreateTask(originalTask, []string{"old", "tags"}); err != nil {
				t.Fatalf("failed to create task: %v", err)
			}

			// Prepare updated task
			updatedTask := &task.Task{
				ID:        "task-888",
				Title:     "Updated title",
				Deadline:  time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC),
				Priority:  3,
				Completed: true,
			}
			newTags := []string{"new", "tags", "updated"} // ORDER BY tag will sort alphabetically

			// Act
			err = storage.UpdateTaskFull(updatedTask, newTags)
			// Assert
			if err != nil {
				t.Errorf("UpdateTaskFull() returned error: %v", err)
			}

			// Verify the update
			retrievedTask, retrievedTags, err := storage.GetTask("task-888")
			if err != nil {
				t.Fatalf("failed to get updated task: %v", err)
			}

			if retrievedTask.Title != updatedTask.Title {
				t.Errorf("expected title %q, got %q", updatedTask.Title, retrievedTask.Title)
			}
			if !retrievedTask.Deadline.Equal(updatedTask.Deadline) {
				t.Errorf("expected deadline %v, got %v", updatedTask.Deadline, retrievedTask.Deadline)
			}
			if retrievedTask.Priority != updatedTask.Priority {
				t.Errorf("expected priority %d, got %d", updatedTask.Priority, retrievedTask.Priority)
			}
			if retrievedTask.Completed != updatedTask.Completed {
				t.Errorf("expected completed %v, got %v", updatedTask.Completed, retrievedTask.Completed)
			}

			// Check tags
			if len(retrievedTags) != len(newTags) {
				t.Errorf("expected %d tags, got %d", len(newTags), len(retrievedTags))
			}
			for i, tag := range newTags {
				if i < len(retrievedTags) && retrievedTags[i] != tag {
					t.Errorf("expected tag %q, got %q", tag, retrievedTags[i])
				}
			}
		})
	})
}
