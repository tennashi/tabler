package storage

import (
	"fmt"
)

// RunMigrations applies database schema migrations
func (s *Storage) RunMigrations() error {
	// Check current schema version
	version, err := s.getSchemaVersion()
	if err != nil {
		return fmt.Errorf("failed to get schema version: %w", err)
	}

	// Apply migrations based on version
	if version < 1 {
		if err := s.migrateTo1(); err != nil {
			return fmt.Errorf("failed to migrate to version 1: %w", err)
		}
	}

	return nil
}

// getSchemaVersion returns the current schema version
func (s *Storage) getSchemaVersion() (int, error) {
	// Create version table if it doesn't exist
	query := `
	CREATE TABLE IF NOT EXISTS schema_version (
		version INTEGER PRIMARY KEY
	);
	`
	if _, err := s.db.Exec(query); err != nil {
		return 0, err
	}

	// Get current version
	var version int
	err := s.db.QueryRow("SELECT version FROM schema_version LIMIT 1").Scan(&version)
	if err != nil {
		// No version yet, this is version 0
		return 0, nil
	}

	return version, nil
}

// migrateTo1 adds parent_task_id column
func (s *Storage) migrateTo1() error {
	// Begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() { _ = tx.Rollback() }()

	// Add parent_task_id column
	query := `
	ALTER TABLE tasks ADD COLUMN parent_task_id TEXT REFERENCES tasks(id);
	`
	if _, err := tx.Exec(query); err != nil {
		return err
	}

	// Update schema version
	if _, err := tx.Exec("INSERT OR REPLACE INTO schema_version (version) VALUES (1)"); err != nil {
		return err
	}

	return tx.Commit()
}
