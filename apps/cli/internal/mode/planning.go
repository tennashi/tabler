package mode

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/task"
)

// PlanningHandler implements planning mode processing
type PlanningHandler struct{}

// NewPlanningHandler creates a new planning mode handler
func NewPlanningHandler() *PlanningHandler {
	return &PlanningHandler{}
}

// Process creates a task with comprehensive breakdown
func (h *PlanningHandler) Process(_ context.Context, input string) (*task.Task, error) {
	// TODO: Implement smart decomposition when available
	// For now, just show mode indicator and create task

	fmt.Println("📋 Planning mode: Let me help you break down this task...")
	fmt.Printf("Task: %s\n", input)
	fmt.Println("(Smart decomposition will be available soon)")

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
