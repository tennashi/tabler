package main

import (
	"os"
	"path/filepath"
	"testing"
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
}
