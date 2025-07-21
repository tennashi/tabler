package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"github.com/tennashi/tabler/internal/task"
)

type Storage struct {
	db *sql.DB
}

func New(dbPath string) (*Storage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Init() error {
	// Create tables
	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT NOT NULL,
		deadline INTEGER,
		priority INTEGER DEFAULT 0,
		completed INTEGER DEFAULT 0,
		created_at INTEGER NOT NULL,
		updated_at INTEGER NOT NULL
	);
	
	CREATE TABLE IF NOT EXISTS task_tags (
		task_id TEXT NOT NULL,
		tag TEXT NOT NULL,
		PRIMARY KEY (task_id, tag),
		FOREIGN KEY (task_id) REFERENCES tasks(id)
	);
	`

	_, err := s.db.Exec(query)
	return err
}

func (s *Storage) CreateTask(t *task.Task, tags []string) error {
	// Start transaction
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		_ = tx.Rollback()
	}()

	// Insert task
	query := `
	INSERT INTO tasks (id, title, deadline, priority, completed, created_at, updated_at)
	VALUES (?, ?, ?, ?, ?, ?, ?)
	`
	_, err = tx.Exec(query,
		t.ID, t.Title, t.Deadline.Unix(), t.Priority,
		t.Completed, t.CreatedAt.Unix(), t.UpdatedAt.Unix())
	if err != nil {
		return err
	}

	// Insert tags
	for _, tag := range tags {
		tagQuery := `INSERT INTO task_tags (task_id, tag) VALUES (?, ?)`
		_, err = tx.Exec(tagQuery, t.ID, tag)
		if err != nil {
			return err
		}
	}

	// Commit transaction
	return tx.Commit()
}
