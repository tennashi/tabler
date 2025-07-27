package mode

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/decomposition"
	"github.com/tennashi/tabler/internal/task"
)

// StorageWithDecomposition extends storage interface for parent-child relationships
type StorageWithDecomposition interface {
	Create(t *task.Task) error
	CreateWithParent(t *task.Task, parentID string) error
}

// Decomposer interface for task decomposition
type Decomposer interface {
	Decompose(ctx context.Context, task string) (*decomposition.DecompositionResult, error)
}

// Presenter interface for interactive presentation
type Presenter interface {
	Present(result *decomposition.DecompositionResult) string
	ParseSelection(input string, total int) ([]int, error)
}

// PlanningHandlerWithDecomposition implements planning mode with smart decomposition
type PlanningHandlerWithDecomposition struct {
	storage    StorageWithDecomposition
	detector   *decomposition.ComplexityDetector
	decomposer Decomposer
	presenter  Presenter
	input      io.Reader // For testing, defaults to os.Stdin
}

// NewPlanningHandlerWithDecomposition creates a new handler with decomposition support
func NewPlanningHandlerWithDecomposition(
	storage StorageWithDecomposition,
	detector *decomposition.ComplexityDetector,
	decomposer Decomposer,
	presenter Presenter,
) *PlanningHandlerWithDecomposition {
	return &PlanningHandlerWithDecomposition{
		storage:    storage,
		detector:   detector,
		decomposer: decomposer,
		presenter:  presenter,
		input:      os.Stdin,
	}
}

// SetInput sets the input reader (for testing)
func (h *PlanningHandlerWithDecomposition) SetInput(input io.Reader) {
	h.input = input
}

// Process creates a task with optional decomposition
func (h *PlanningHandlerWithDecomposition) Process(ctx context.Context, input string) (*task.Task, error) {
	// Check if task is complex
	isComplex, reason := h.detector.DetectComplexity(input)

	if !isComplex {
		// Simple task - create directly
		return h.createSimpleTask(input)
	}

	// Complex task - offer decomposition
	fmt.Printf("📋 Planning mode: This looks like a complex task (%s)\n", reason)
	fmt.Println("Let me help you break it down...")

	// Try to decompose
	result, err := h.decomposer.Decompose(ctx, input)
	if err != nil {
		// Fall back to simple task creation
		fmt.Printf("⚠️  Could not decompose task: %v\n", err)
		fmt.Println("Creating single task instead...")
		return h.createSimpleTask(input)
	}

	// Present decomposition options
	presentation := h.presenter.Present(result)
	fmt.Print(presentation)

	// Get user selection
	reader := bufio.NewReader(h.input)
	selection, _ := reader.ReadString('\n')
	selection = strings.TrimSpace(selection)

	// Parse selection
	selectedIndices, err := h.presenter.ParseSelection(selection, len(result.Subtasks))
	if err != nil {
		fmt.Printf("⚠️  Invalid selection: %v\n", err)
		fmt.Println("Creating single task instead...")
		return h.createSimpleTask(input)
	}

	// Create parent task
	parentTask, err := h.createSimpleTask(input)
	if err != nil {
		return nil, err
	}

	// Create selected subtasks
	if len(selectedIndices) > 0 {
		fmt.Printf("✅ Creating %d subtasks...\n", len(selectedIndices))
		for _, idx := range selectedIndices {
			if idx > 0 && idx <= len(result.Subtasks) {
				subtaskTitle := result.Subtasks[idx-1]
				if err := h.createSubtask(subtaskTitle, parentTask.ID); err != nil {
					fmt.Printf("⚠️  Failed to create subtask %q: %v\n", subtaskTitle, err)
				}
			}
		}
	}

	return parentTask, nil
}

// createSimpleTask creates a single task without decomposition
func (h *PlanningHandlerWithDecomposition) createSimpleTask(title string) (*task.Task, error) {
	now := time.Now()
	t := &task.Task{
		ID:        uuid.New().String(),
		Title:     title,
		Priority:  0,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := h.storage.Create(t); err != nil {
		return nil, err
	}

	return t, nil
}

// createSubtask creates a subtask with parent relationship
func (h *PlanningHandlerWithDecomposition) createSubtask(title, parentID string) error {
	now := time.Now()
	t := &task.Task{
		ID:        uuid.New().String(),
		Title:     title,
		Priority:  0,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}

	return h.storage.CreateWithParent(t, parentID)
}
