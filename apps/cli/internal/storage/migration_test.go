package storage

import (
	"os"
	"testing"
)

func TestMigrations(t *testing.T) {
	t.Run("RunMigrations", func(t *testing.T) {
		t.Run("should create parent_task_id column", func(t *testing.T) {
			// Arrange
			tmpfile, err := os.CreateTemp("", "test-*.db")
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = os.Remove(tmpfile.Name()) }()
			_ = tmpfile.Close()

			storage, err := New(tmpfile.Name())
			if err != nil {
				t.Fatal(err)
			}
			defer func() { _ = storage.Close() }()

			// Initialize with original schema
			if err := storage.Init(); err != nil {
				t.Fatal(err)
			}

			// Act - run migrations
			if err := storage.RunMigrations(); err != nil {
				t.Fatal(err)
			}

			// Assert - check if parent_task_id column exists
			var columnExists bool
			query := `
			SELECT COUNT(*) FROM pragma_table_info('tasks') 
			WHERE name = 'parent_task_id'
			`
			var count int
			err = storage.db.QueryRow(query).Scan(&count)
			if err != nil {
				t.Fatal(err)
			}
			columnExists = count > 0

			if !columnExists {
				t.Error("expected parent_task_id column to exist")
			}
		})
	})
}
