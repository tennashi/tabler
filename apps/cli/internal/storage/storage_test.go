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
			defer func() {
				if err := storage.Close(); err != nil {
					t.Errorf("failed to close storage: %v", err)
				}
			}()

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
			defer func() {
				if err := storage.Close(); err != nil {
					t.Errorf("failed to close storage: %v", err)
				}
			}()

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
}
