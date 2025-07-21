package main

import (
	"fmt"
	"strings"

	"github.com/tennashi/tabler/internal/service"
)

const (
	idDisplayWidth   = 6
	idColumnWidth    = 7
	taskColumnWidth  = 23
	statusPending    = "[ ]"
	statusCompleted  = "[✓]"
)

func formatTasksAsTable(taskItems []*service.TaskItem) string {
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

func truncateString(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}