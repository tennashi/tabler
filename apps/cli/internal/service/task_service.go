package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/metadata"
	"github.com/tennashi/tabler/internal/parser"
	"github.com/tennashi/tabler/internal/storage"
	"github.com/tennashi/tabler/internal/task"
)

var ErrEmptyTitle = fmt.Errorf("task title cannot be empty")

type TaskService struct {
	storage  *storage.Storage
	metadata *metadata.Service
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
		storage:  store,
		metadata: nil, // No metadata service by default
	}, nil
}

func NewTaskServiceWithMetadata(dataDir string, metadataService *metadata.Service) (*TaskService, error) {
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
		storage:  store,
		metadata: metadataService,
	}, nil
}

func (s *TaskService) Close() error {
	return s.storage.Close()
}

func (s *TaskService) CreateTaskFromInput(input string) (string, error) {
	var result *parser.ParseResult
	var tags []string

	// If we have a metadata service, use it for extraction
	if s.metadata != nil {
		ctx := context.Background()
		extracted, err := s.metadata.Extract(ctx, input)
		if err == nil && extracted != nil {
			// Use LLM-extracted metadata
			result = &parser.ParseResult{
				Title: extracted.CleanedText,
				Tags:  []string{}, // We'll use tags from extracted
			}
			tags = extracted.Tags

			// Convert priority string to int
			switch extracted.Priority {
			case "high":
				result.Priority = 3
			case "medium":
				result.Priority = 2
			case "low":
				result.Priority = 1
			default:
				result.Priority = 0
			}

			// Parse deadline if provided
			if extracted.Deadline != "" {
				if deadline, err := time.Parse("2006-01-02", extracted.Deadline); err == nil {
					result.Deadline = &deadline
				}
			}
		} else {
			// Fallback to parser if LLM extraction fails
			result = parser.Parse(input)
			tags = result.Tags
		}
	} else {
		// No metadata service, use parser only
		result = parser.Parse(input)
		tags = result.Tags
	}

	// Validate title is not empty
	if strings.TrimSpace(result.Title) == "" {
		return "", ErrEmptyTitle
	}

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
	if err := s.storage.CreateTask(task, tags); err != nil {
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

type FilterOptions struct {
	Tag     string
	Today   bool
	Overdue bool
}

func (s *TaskService) ListTasks(filter *FilterOptions) ([]*TaskItem, error) {
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

		// Apply filters
		if filter != nil {
			// Tag filter
			if filter.Tag != "" {
				hasTag := false
				for _, tag := range tags {
					if tag == filter.Tag {
						hasTag = true
						break
					}
				}
				if !hasTag {
					continue
				}
			}
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

	// Validate title is not empty
	if strings.TrimSpace(result.Title) == "" {
		return ErrEmptyTitle
	}

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
