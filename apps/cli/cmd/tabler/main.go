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

func completeTask(service *service.TaskService, taskID string) error {
	err := service.CompleteTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to complete task: %w", err)
	}

	fmt.Printf("Task completed: %s\n", taskID)
	return nil
}

func showTask(service *service.TaskService, taskID string) error {
	task, tags, err := service.GetTask(taskID)
	if err != nil {
		return fmt.Errorf("failed to get task: %w", err)
	}

	// Simple output for now
	fmt.Printf("ID: %s\n", task.ID)
	fmt.Printf("Title: %s\n", task.Title)
	fmt.Printf("Priority: %d\n", task.Priority)
	fmt.Printf("Deadline: %s\n", task.Deadline.Format("2006-01-02"))
	fmt.Printf("Completed: %v\n", task.Completed)
	if len(tags) > 0 {
		fmt.Printf("Tags: %v\n", tags)
	}

	return nil
}
