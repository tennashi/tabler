package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/tennashi/tabler/internal/service"
	"github.com/tennashi/tabler/internal/task"
)

const (
	idDisplayWidth   = 6
	idColumnWidth    = 7
	taskColumnWidth  = 23
	statusPending    = "[ ]"
	statusCompleted  = "[✓]"
	dateFormat       = "Jan 2, 2006"
	dateTimeFormat   = "Jan 2, 2006 3:04 PM"
	
	// Extended format column widths
	extTaskColumnWidth     = 31
	extTagsColumnWidth     = 12
	extPriorityColumnWidth = 3
	extDeadlineColumnWidth = 11
)

func formatTasksAsTable(taskItems []*service.TaskItem) string {
	// Check if any task has metadata
	hasMetadata := false
	for _, item := range taskItems {
		if len(item.Tags) > 0 || item.Task.Priority > 0 || !item.Task.Deadline.IsZero() {
			hasMetadata = true
			break
		}
	}

	if hasMetadata {
		return formatTasksWithMetadata(taskItems)
	}
	return formatTasksCompact(taskItems)
}

func formatTasksCompact(taskItems []*service.TaskItem) string {
	var result strings.Builder

	// Header
	result.WriteString("ID      Task                    Status\n")
	result.WriteString("---     --------------------    ------\n")

	// Rows
	for _, item := range taskItems {
		status := statusPending
		if item.Task.Completed {
			status = statusCompleted
		}

		// Format with fixed width columns
		result.WriteString(fmt.Sprintf("%-*s %-*s %s\n",
			idColumnWidth, item.Task.ID[:idDisplayWidth],
			taskColumnWidth, truncateString(item.Task.Title, taskColumnWidth),
			status))
	}

	// Remove trailing newline
	return strings.TrimRight(result.String(), "\n")
}

func formatTasksWithMetadata(taskItems []*service.TaskItem) string {
	var result strings.Builder

	// Header
	result.WriteString("ID      Task                             Tags          Pri  Deadline     Status\n")
	result.WriteString("------  -------------------------------  ------------  ---  -----------  ------\n")

	// Rows
	for _, item := range taskItems {
		status := statusPending
		if item.Task.Completed {
			status = statusCompleted
		}

		// Format tags
		tags := "-"
		if len(item.Tags) > 0 {
			tags = strings.Join(item.Tags, ", ")
		}

		// Format priority
		priority := "-"
		if item.Task.Priority > 0 {
			priority = strings.Repeat("!", item.Task.Priority)
		}

		// Format deadline
		deadline := "-"
		if !item.Task.Deadline.IsZero() {
			deadline = item.Task.Deadline.Format("Jan 2")
		}

		// Format with fixed width columns
		result.WriteString(fmt.Sprintf("%-*s  %-*s  %-*s  %-*s  %-*s  %s\n",
			idDisplayWidth, item.Task.ID[:idDisplayWidth],
			extTaskColumnWidth, truncateString(item.Task.Title, extTaskColumnWidth),
			extTagsColumnWidth, truncateString(tags, extTagsColumnWidth),
			extPriorityColumnWidth, priority,
			extDeadlineColumnWidth, deadline,
			status))
	}

	// Remove trailing newline
	return strings.TrimRight(result.String(), "\n")
}

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}

func formatTaskDetails(task *task.Task, tags []string) string {
	var result strings.Builder

	result.WriteString(fmt.Sprintf("ID: %s\n", task.ID))
	result.WriteString(fmt.Sprintf("Task: %s\n", task.Title))
	
	// Status
	status := "Pending"
	if task.Completed {
		status = "Completed"
	}
	result.WriteString(fmt.Sprintf("Status: %s\n", status))
	
	// Tags
	if len(tags) > 0 {
		result.WriteString(fmt.Sprintf("Tags: %s\n", strings.Join(tags, ", ")))
	}
	
	// Priority
	priorityName := getPriorityName(task.Priority)
	result.WriteString(fmt.Sprintf("Priority: %s\n", priorityName))
	
	// Deadline
	if !task.Deadline.IsZero() {
		result.WriteString(fmt.Sprintf("Deadline: %s\n", task.Deadline.Format(dateFormat)))
	}
	
	// Created
	result.WriteString(fmt.Sprintf("Created: %s\n", formatDateTime(task.CreatedAt)))
	
	// Modified
	result.WriteString(fmt.Sprintf("Modified: %s", formatDateTime(task.UpdatedAt)))
	
	return result.String()
}

func getPriorityName(priority int) string {
	switch priority {
	case 1:
		return "Low"
	case 2:
		return "Medium"
	case 3:
		return "High"
	default:
		return "None"
	}
}

func formatDateTime(t time.Time) string {
	return t.Format(dateTimeFormat)
}