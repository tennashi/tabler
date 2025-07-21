package storage

import (
	"database/sql"
	"time"

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

func (s *Storage) GetTask(id string) (*task.Task, []string, error) {
	// Get task
	var t task.Task
	var deadlineUnix, createdAtUnix, updatedAtUnix int64
	var completed bool

	query := `
	SELECT id, title, deadline, priority, completed, created_at, updated_at
	FROM tasks
	WHERE id = ?
	`

	err := s.db.QueryRow(query, id).Scan(
		&t.ID, &t.Title, &deadlineUnix, &t.Priority,
		&completed, &createdAtUnix, &updatedAtUnix,
	)
	if err != nil {
		return nil, nil, err
	}

	// Convert Unix timestamps to time.Time
	t.Deadline = time.Unix(deadlineUnix, 0).UTC()
	t.CreatedAt = time.Unix(createdAtUnix, 0).UTC()
	t.UpdatedAt = time.Unix(updatedAtUnix, 0).UTC()
	t.Completed = completed

	// Get tags
	tagQuery := `SELECT tag FROM task_tags WHERE task_id = ? ORDER BY tag`
	rows, err := s.db.Query(tagQuery, id)
	if err != nil {
		return nil, nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var tags []string
	for rows.Next() {
		var tag string
		if err := rows.Scan(&tag); err != nil {
			return nil, nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, nil, err
	}

	return &t, tags, nil
}

func (s *Storage) ListTasks(_ map[string]interface{}) ([]*task.Task, error) {
	query := `
	SELECT id, title, deadline, priority, completed, created_at, updated_at
	FROM tasks
	ORDER BY created_at DESC, id DESC
	`

	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = rows.Close()
	}()

	var tasks []*task.Task
	for rows.Next() {
		var t task.Task
		var deadlineUnix, createdAtUnix, updatedAtUnix int64
		var completed bool

		err := rows.Scan(
			&t.ID, &t.Title, &deadlineUnix, &t.Priority,
			&completed, &createdAtUnix, &updatedAtUnix,
		)
		if err != nil {
			return nil, err
		}

		// Convert Unix timestamps to time.Time
		t.Deadline = time.Unix(deadlineUnix, 0).UTC()
		t.CreatedAt = time.Unix(createdAtUnix, 0).UTC()
		t.UpdatedAt = time.Unix(updatedAtUnix, 0).UTC()
		t.Completed = completed

		tasks = append(tasks, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tasks, nil
}
