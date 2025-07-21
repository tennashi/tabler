package service

import (
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/parser"
	"github.com/tennashi/tabler/internal/storage"
	"github.com/tennashi/tabler/internal/task"
)

type TaskService struct {
	storage *storage.Storage
}

func NewTaskService(dataDir string) (*TaskService, error) {
	// Create data directory if it doesn't exist
	dbPath := filepath.Join(dataDir, "tasks.db")

	// Initialize storage
	store, err := storage.New(dbPath)
	if err != nil {
		return nil, err
	}

	if err := store.Init(); err != nil {
		_ = store.Close()
		return nil, err
	}

	return &TaskService{
		storage: store,
	}, nil
}

func (s *TaskService) Close() error {
	return s.storage.Close()
}

func (s *TaskService) CreateTaskFromInput(input string) (string, error) {
	// Parse input
	result := parser.Parse(input)

	// Generate task ID
	taskID := uuid.New().String()

	// Create task
	now := time.Now()

	// Handle deadline
	var deadline time.Time
	if result.Deadline != nil {
		deadline = *result.Deadline
	}

	task := &task.Task{
		ID:        taskID,
		Title:     result.Title,
		Deadline:  deadline,
		Priority:  result.Priority,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	// Store task with tags
	if err := s.storage.CreateTask(task, result.Tags); err != nil {
		return "", err
	}

	return taskID, nil
}

func (s *TaskService) GetTask(id string) (*task.Task, []string, error) {
	return s.storage.GetTask(id)
}

// TaskItem represents a task with its associated tags
type TaskItem struct {
	Task *task.Task
	Tags []string
}

func (s *TaskService) ListTasks() ([]*TaskItem, error) {
	tasks, err := s.storage.ListTasks(nil)
	if err != nil {
		return nil, err
	}

	// Get tags for each task
	taskItems := make([]*TaskItem, 0, len(tasks))
	for _, t := range tasks {
		_, tags, err := s.storage.GetTask(t.ID)
		if err != nil {
			return nil, err
		}
		taskItems = append(taskItems, &TaskItem{
			Task: t,
			Tags: tags,
		})
	}

	return taskItems, nil
}

func (s *TaskService) CompleteTask(id string) error {
	return s.storage.UpdateTaskCompleted(id, true)
}

func (s *TaskService) DeleteTask(id string) error {
	return s.storage.DeleteTask(id)
}

func (s *TaskService) UpdateTaskFromInput(id string, input string) error {
	// Parse new input
	result := parser.Parse(input)

	// Get existing task to preserve creation time
	existingTask, _, err := s.storage.GetTask(id)
	if err != nil {
		return err
	}

	// Handle deadline
	var deadline time.Time
	if result.Deadline != nil {
		deadline = *result.Deadline
	}

	// Create updated task
	updatedTask := &task.Task{
		ID:        id,
		Title:     result.Title,
		Deadline:  deadline,
		Priority:  result.Priority,
		Completed: existingTask.Completed, // Preserve completion status
		CreatedAt: existingTask.CreatedAt, // Preserve creation time
		UpdatedAt: time.Now(),
	}

	// Update task with new tags
	return s.storage.UpdateTaskFull(updatedTask, result.Tags)
}
