package mode

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/task"
)

// TalkHandler implements talk mode processing
type TalkHandler struct{}

// NewTalkHandler creates a new talk mode handler
func NewTalkHandler() *TalkHandler {
	return &TalkHandler{}
}

// Process creates a task through conversational refinement
func (h *TalkHandler) Process(_ context.Context, input string) (*task.Task, error) {
	// TODO: Implement interactive clarification when available
	// For now, just show mode indicator and create task

	fmt.Println("🤔 Talk mode: Let me help you refine this task...")
	fmt.Printf("Task: %s\n", input)
	fmt.Println("(Interactive clarification will be available soon)")

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
