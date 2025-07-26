package decomposition

import (
	"context"
	"fmt"
	"strings"
)

// ClaudeClient defines the interface for Claude interaction
type ClaudeClient interface {
	Execute(ctx context.Context, prompt string) (string, error)
}

// DecompositionResult contains the decomposed subtasks
type DecompositionResult struct {
	OriginalTask string
	Subtasks     []string
	Rationale    string
}

// TaskDecomposer generates subtask suggestions using Claude
type TaskDecomposer struct {
	claude ClaudeClient
}

// NewTaskDecomposer creates a new task decomposer
func NewTaskDecomposer(claude ClaudeClient) *TaskDecomposer {
	return &TaskDecomposer{
		claude: claude,
	}
}

// Decompose breaks down a complex task into subtasks
func (d *TaskDecomposer) Decompose(ctx context.Context, task string) (*DecompositionResult, error) {
	// Create prompt for Claude
	prompt := fmt.Sprintf(`Break down this task into clear, actionable subtasks:
Task: "%s"

Please provide 3-7 specific subtasks that would complete this task.
Format each subtask as a numbered list item.
Keep each subtask concise and actionable.`, task)

	// Call Claude
	response, err := d.claude.Execute(ctx, prompt)
	if err != nil {
		return nil, fmt.Errorf("failed to get decomposition from Claude: %w", err)
	}

	// Parse response
	subtasks := d.parseSubtasks(response)

	return &DecompositionResult{
		OriginalTask: task,
		Subtasks:     subtasks,
		Rationale:    "Task broken down into actionable steps",
	}, nil
}

// parseSubtasks extracts subtasks from Claude's response
func (d *TaskDecomposer) parseSubtasks(response string) []string {
	var subtasks []string
	lines := strings.Split(response, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		// Look for numbered items (1. , 2. , etc.)
		if len(line) > 2 && line[0] >= '1' && line[0] <= '9' && line[1] == '.' {
			// Extract the task text after the number
			taskText := strings.TrimSpace(line[2:])
			if taskText != "" {
				subtasks = append(subtasks, taskText)
			}
		}
	}

	return subtasks
}
