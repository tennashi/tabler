package main

import (
	"errors"
	"fmt"
	"strings"
)

var (
	ErrTaskNotFound  = errors.New("task not found")
	ErrDatabaseError = errors.New("database error")
	ErrEmptyTitle    = errors.New("empty title")
)

type command struct {
	name        string
	description string
}

var availableCommands = []command{
	{"add", "Add a new task"},
	{"list", "List all tasks"},
	{"done", "Mark a task as completed"},
	{"show", "Show task details"},
	{"delete", "Delete a task"},
	{"update", "Update a task"},
}

func formatTaskError(err error, taskID string) string {
	if errors.Is(err, ErrTaskNotFound) {
		return fmt.Sprintf(`Task not found: %s

The task with this ID doesn't exist. Please check the ID and try again.
You can use 'tabler list' to see all tasks.`, taskID)
	}
	return err.Error()
}

func formatStorageError(err error) string {
	if errors.Is(err, ErrDatabaseError) {
		return "Unable to access task storage. Please check if the data directory is accessible."
	}
	return err.Error()
}

func formatValidationError(err error) string {
	if errors.Is(err, ErrEmptyTitle) {
		return "Task description cannot be empty. Please provide a meaningful description for your task."
	}
	return err.Error()
}

func isNotFoundError(errMsg string) bool {
	return strings.Contains(errMsg, "sql: no rows in result set")
}

func formatUnknownCommandError(cmd string) string {
	// Simple command suggestion
	suggestion := ""
	for _, c := range availableCommands {
		if strings.HasPrefix(c.name, cmd) || strings.HasPrefix(cmd, c.name[:1]) {
			suggestion = c.name
			break
		}
	}
	
	result := fmt.Sprintf("Unknown command: %s\n\n", cmd)
	
	if suggestion != "" {
		result += fmt.Sprintf("Did you mean '%s'?\n\n", suggestion)
	}
	
	result += "Available commands:\n"
	for _, c := range availableCommands {
		result += fmt.Sprintf("  %-8s - %s\n", c.name, c.description)
	}
	
	// Remove trailing newline
	return strings.TrimRight(result, "\n")
}