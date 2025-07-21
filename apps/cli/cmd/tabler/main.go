package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tennashi/tabler/internal/service"
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
		if len(os.Args) < 3 {
			return fmt.Errorf("usage: tabler add <task description>")
		}
		input := os.Args[2]
		return addTask(taskService, input)
	case "list":
		return listTasks(taskService)
	default:
		return fmt.Errorf("unknown command: %s", command)
	}
}

func addTask(service *service.TaskService, input string) error {
	taskID, err := service.CreateTaskFromInput(input)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}

	fmt.Printf("Task created: %s\n", taskID)
	return nil
}

func listTasks(service *service.TaskService) error {
	taskItems, err := service.ListTasks()
	if err != nil {
		return fmt.Errorf("failed to list tasks: %w", err)
	}

	if len(taskItems) == 0 {
		fmt.Println("No tasks found.")
		return nil
	}

	// Simple output for now
	for _, item := range taskItems {
		fmt.Printf("- %s\n", item.Task.Title)
	}

	return nil
}
