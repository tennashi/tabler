package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/tennashi/tabler/internal/metadata"
	"github.com/tennashi/tabler/internal/mode"
	service "github.com/tennashi/tabler/internal/service"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	if len(os.Args) < 2 {
		return fmt.Errorf("usage: tabler <command> [arguments]")
	}

	command := os.Args[1]

	// Get data directory
	dataDir := os.Getenv("TABLER_DATA_DIR")
	if dataDir == "" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get home directory: %w", err)
		}
		dataDir = filepath.Join(homeDir, ".tabler")
	}

	// Ensure data directory exists
	if err := os.MkdirAll(dataDir, 0o750); err != nil {
		return fmt.Errorf("failed to create data directory: %w", err)
	}

	// Initialize service
	taskService, err := service.NewTaskService(dataDir)
	if err != nil {
		return fmt.Errorf("failed to initialize service: %w", err)
	}
	defer func() {
		_ = taskService.Close()
	}()

	switch command {
	case "add":
		return handleAddCommand(taskService, os.Args[2:])
	case "list":
		return handleListCommand(taskService, os.Args[2:])
	case "done":
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: tabler done <task-id>")
		}
		taskID := os.Args[2]
		return completeTask(taskService, taskID)
	case "show":
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: tabler show <task-id>")
		}
		taskID := os.Args[2]
		return showTask(taskService, taskID)
	case "delete":
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: tabler delete <task-id>")
		}
		taskID := os.Args[2]
		return deleteTask(taskService, taskID)
	case "update":
		if len(os.Args) < 4 {
			return fmt.Errorf("usage: tabler update <task-id> <new description>")
		}
		taskID := os.Args[2]
		newInput := os.Args[3]
		return updateTask(taskService, taskID, newInput)
	default:
		return errors.New(formatUnknownCommandError(command))
	}
}

func handleAddCommand(taskService *service.TaskService, args []string) error {
	// Parse flags for add command
	addFlags := flag.NewFlagSet("add", flag.ContinueOnError)
	useAI := addFlags.Bool("ai", false, "Use AI to extract metadata from task description")
	useTalk := addFlags.Bool("talk", false, "Use interactive dialogue to clarify vague tasks")

	// Find where the task description starts (after flags)
	var taskDescStart int
	for i, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			taskDescStart = i
			break
		}
	}

	// Parse flags up to the task description
	if taskDescStart > 0 {
		if err := addFlags.Parse(args[:taskDescStart]); err != nil {
			return fmt.Errorf("failed to parse flags: %w", err)
		}
	}

	// Get task description
	if taskDescStart >= len(args) {
		return fmt.Errorf("usage: tabler add [--ai] [--talk] <task description>")
	}

	input := strings.Join(args[taskDescStart:], " ")

	// If --ai flag is set, create a new service with metadata extraction
	if *useAI {
		// Get data directory from the existing service
		dataDir := os.Getenv("TABLER_DATA_DIR")
		if dataDir == "" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				return fmt.Errorf("failed to get home directory: %w", err)
			}
			dataDir = filepath.Join(homeDir, ".tabler")
		}

		// Create metadata service
		claude := metadata.NewClaudeClient()
		metadataService := metadata.NewService(claude)

		// Create new task service with metadata
		aiTaskService, err := service.NewTaskServiceWithMetadata(dataDir, metadataService)
		if err != nil {
			return fmt.Errorf("failed to create AI-enhanced service: %w", err)
		}
		defer func() {
			_ = aiTaskService.Close()
		}()

		// Use the AI-enhanced service
		taskService = aiTaskService

		fmt.Println("📋 Using AI to extract metadata...")
	}

	// Check if talk mode is forced via flag
	if *useTalk {
		// Prepend /talk to use talk mode with clarification
		return addTaskWithMode(taskService, "/talk "+input)
	}

	// Check if mode prefixes are used
	if strings.HasPrefix(input, "/") {
		// Use mode system for inputs with mode prefixes
		return addTaskWithMode(taskService, input)
	}

	// For backward compatibility, use original method without prefix
	return addTask(taskService, input)
}

func addTaskWithMode(service *service.TaskService, input string) error {
	// Create mode manager with enhanced features
	modeManager := mode.NewManagerBuilder().
		WithClarification().
		Build()

	// Process task with mode system
	ctx := context.Background()
	task, err := modeManager.ProcessTask(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to process task: %w", err)
	}

	// Store task
	taskID, err := service.StoreTask(task)
	if err != nil {
		return fmt.Errorf("failed to store task: %w", err)
	}

	fmt.Printf("Task created: %s\n", taskID)
	return nil
}

func addTask(service *service.TaskService, input string) error {
	taskID, err := service.CreateTaskFromInput(input)
	if err != nil {
		if strings.Contains(err.Error(), "task title cannot be empty") {
			return errors.New(formatValidationError(ErrEmptyTitle))
		}
		return fmt.Errorf("failed to create task: %w", err)
	}

	fmt.Printf("Task created: %s\n", taskID)
	return nil
}

func handleListCommand(taskService *service.TaskService, args []string) error {
	filter := &service.FilterOptions{}

	// Parse flags
	i := 0
	for i < len(args) {
		switch args[i] {
		case "--tag":
			if i+1 >= len(args) {
				return fmt.Errorf("--tag requires a value")
			}
			filter.Tag = args[i+1]
			i += 2
		default:
			return fmt.Errorf("unknown flag: %s", args[i])
		}
	}

	return listTasks(taskService, filter)
}

func listTasks(taskService *service.TaskService, filter *service.FilterOptions) error {
	taskItems, err := taskService.ListTasks(filter)
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}

	if len(taskItems) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	// Display tasks in table format
	fmt.Println(formatTasksAsTable(taskItems))

	return nil
}

func completeTask(service *service.TaskService, taskID string) error {
	err := service.CompleteTask(taskID)
	if err != nil {
		if isNotFoundError(err.Error()) {
			return errors.New(formatTaskError(ErrTaskNotFound, taskID))
		}
		return fmt.Errorf("failed to complete task: %w", err)
	}

	fmt.Printf("Task completed: %s\n", taskID)
	return nil
}

func showTask(service *service.TaskService, taskID string) error {
	task, tags, err := service.GetTask(taskID)
	if err != nil {
		if isNotFoundError(err.Error()) {
			return errors.New(formatTaskError(ErrTaskNotFound, taskID))
		}
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Display formatted task details
	fmt.Println(formatTaskDetails(task, tags))

	return nil
}

func deleteTask(service *service.TaskService, taskID string) error {
	// Get task details first to show title in confirmation
	task, _, err := service.GetTask(taskID)
	if err != nil {
		if isNotFoundError(err.Error()) {
			return errors.New(formatTaskError(ErrTaskNotFound, taskID))
		}
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Skip confirmation in non-interactive mode (for tests)
	if os.Getenv("TABLER_NON_INTERACTIVE") != "1" {
		// Confirm deletion
		if !confirmDeletion(task.Title, os.Stdin) {
			fmt.Println("Deletion cancelled.")
			return nil
		}
	}

	err = service.DeleteTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}

	fmt.Printf("Task deleted: %s\n", taskID)
	return nil
}

func updateTask(service *service.TaskService, taskID string, newInput string) error {
	err := service.UpdateTaskFromInput(taskID, newInput)
	if err != nil {
		if strings.Contains(err.Error(), "task title cannot be empty") {
			return errors.New(formatValidationError(ErrEmptyTitle))
		}
		if isNotFoundError(err.Error()) {
			return errors.New(formatTaskError(ErrTaskNotFound, taskID))
		}
		return fmt.Errorf("failed to update task: %w", err)
	}

	fmt.Printf("Task updated: %s\n", taskID)
	return nil
}
