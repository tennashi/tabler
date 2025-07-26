package mode

import (
	"context"
	"fmt"

	"github.com/tennashi/tabler/internal/task"
)

// ModeHandler defines the interface for mode-specific task processing
type ModeHandler interface {
	Process(ctx context.Context, input string) (*task.Task, error)
}

// ProcessResult contains the result of processing mode input
type ProcessResult struct {
	Mode     Mode
	TaskText string
}

// ModeManager coordinates mode selection and processing
type ModeManager struct {
	handlers map[Mode]ModeHandler
	detector *ModeDetector
}

// NewModeManager creates a new mode manager with basic handlers
func NewModeManager() *ModeManager {
	// Use builder for consistent initialization
	return NewManagerBuilder().Build()
}

// RegisterHandler registers a handler for a specific mode
func (m *ModeManager) RegisterHandler(mode Mode, handler ModeHandler) {
	m.handlers[mode] = handler
}

// ProcessInput processes the input and returns the mode and task text
func (m *ModeManager) ProcessInput(input string) (*ProcessResult, error) {
	mode, taskText, _ := ParseModePrefix(input)

	return &ProcessResult{
		Mode:     mode,
		TaskText: taskText,
	}, nil
}

// ProcessTask processes the input using the appropriate mode handler
func (m *ModeManager) ProcessTask(ctx context.Context, input string) (*task.Task, error) {
	mode, taskText, hasPrefix := ParseModePrefix(input)

	// If no prefix, use auto-detection
	if !hasPrefix && m.detector != nil {
		mode = m.detector.DetectMode(input)
	}

	handler, exists := m.handlers[mode]
	if !exists {
		return nil, fmt.Errorf("no handler registered for mode: %s", mode)
	}

	return handler.Process(ctx, taskText)
}
