package main

import (
	"strings"
	"testing"
)

func TestUserFriendlyErrors(t *testing.T) {
	t.Run("should format task not found error", func(t *testing.T) {
		// Arrange
		err := ErrTaskNotFound
		taskID := "abc123"

		// Act
		result := formatTaskError(err, taskID)

		// Assert
		expected := `Task not found: abc123

The task with this ID doesn't exist. Please check the ID and try again.
You can use 'tabler list' to see all tasks.`

		if result != expected {
			t.Errorf("expected:\n%s\n\ngot:\n%s", expected, result)
		}
	})

	t.Run("should format database error", func(t *testing.T) {
		// Arrange
		err := ErrDatabaseError

		// Act
		result := formatStorageError(err)

		// Assert
		expected := "Unable to access task storage. Please check if the data directory is accessible."

		if result != expected {
			t.Errorf("expected:\n%s\n\ngot:\n%s", expected, result)
		}
	})

	t.Run("should format empty title error", func(t *testing.T) {
		// Arrange
		err := ErrEmptyTitle

		// Act
		result := formatValidationError(err)

		// Assert
		expected := "Task description cannot be empty. Please provide a meaningful description for your task."

		if result != expected {
			t.Errorf("expected:\n%s\n\ngot:\n%s", expected, result)
		}
	})
}

func TestErrorWrapping(t *testing.T) {
	t.Run("should detect SQL no rows error", func(t *testing.T) {
		// Arrange
		sqlError := "sql: no rows in result set"

		// Act
		isNotFound := isNotFoundError(sqlError)

		// Assert
		if !isNotFound {
			t.Error("expected SQL no rows error to be detected as not found")
		}
	})

	t.Run("should show helpful message for unknown command", func(t *testing.T) {
		// Arrange
		command := "lst" // typo for "list"

		// Act
		result := formatUnknownCommandError(command)

		// Assert
		if !strings.Contains(result, "Unknown command: lst") {
			t.Error("expected error to contain unknown command")
		}
		if !strings.Contains(result, "Did you mean 'list'?") {
			t.Error("expected error to suggest similar command")
		}
		if !strings.Contains(result, "Available commands:") {
			t.Error("expected error to list available commands")
		}
	})
}
