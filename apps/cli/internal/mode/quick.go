package mode

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/task"
)

// QuickHandler implements quick mode processing
type QuickHandler struct{}

// NewQuickHandler creates a new quick mode handler
func NewQuickHandler() *QuickHandler {
	return &QuickHandler{}
}

// Process creates a task with minimal processing
func (h *QuickHandler) Process(_ context.Context, input string) (*task.Task, error) {
	// Quick mode: just create the task with the raw input
	// No parsing, no AI, no enhancements
	now := time.Now()
	return &task.Task{
		ID:        uuid.New().String(),
		Title:     input,
		Priority:  0,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}
