package storage

import (
	"os"
	"path/filepath"
	"testing"
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
}
