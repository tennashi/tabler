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
}
