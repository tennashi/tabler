package mode

import (
	"context"
	"fmt"
	"testing"

	"github.com/tennashi/tabler/internal/decomposition"
	"github.com/tennashi/tabler/internal/task"
)

func TestPlanningHandlerWithDecomposition(t *testing.T) {
	t.Run("process with decomposition", func(t *testing.T) {
		t.Run("should detect complex task and offer decomposition", func(t *testing.T) {
			// Arrange
			storage := &mockStorageWithDecomposition{
				tasks: make(map[string]*task.Task),
			}
			detector := decomposition.NewComplexityDetector()
			decomposer := &mockDecomposer{
				result: &decomposition.DecompositionResult{
					OriginalTask: "organize conference",
					Subtasks: []string{
						"Book venue",
						"Invite speakers",
						"Setup registration",
					},
				},
			}
			presenter := &mockPresenter{
				selectedIndices: []int{1, 3}, // Select first and third subtask
			}

			handler := NewPlanningHandlerWithDecomposition(storage, detector, decomposer, presenter)
			input := "organize conference"

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result == nil {
				t.Fatal("expected result to be non-nil")
			}

			// Check parent task was created
			if result.Title != "organize conference" {
				t.Errorf("expected parent title %q, got %q", "organize conference", result.Title)
			}

			// Check subtasks were created
			if len(storage.createdTasks) != 3 { // parent + 2 selected subtasks
				t.Errorf("expected 3 tasks created, got %d", len(storage.createdTasks))
			}

			// Verify decomposer was called
			if !decomposer.called {
				t.Error("expected decomposer to be called")
			}

			// Verify presenter was called
			if !presenter.called {
				t.Error("expected presenter to be called")
			}
		})

		t.Run("should skip decomposition for simple task", func(t *testing.T) {
			// Arrange
			storage := &mockStorageWithDecomposition{
				tasks: make(map[string]*task.Task),
			}
			detector := decomposition.NewComplexityDetector()
			decomposer := &mockDecomposer{}
			presenter := &mockPresenter{}

			handler := NewPlanningHandlerWithDecomposition(storage, detector, decomposer, presenter)
			input := "buy milk" // Simple task

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Title != "buy milk" {
				t.Errorf("expected title %q, got %q", "buy milk", result.Title)
			}

			// Check only one task was created
			if len(storage.createdTasks) != 1 {
				t.Errorf("expected 1 task created, got %d", len(storage.createdTasks))
			}

			// Verify decomposer was NOT called
			if decomposer.called {
				t.Error("expected decomposer NOT to be called for simple task")
			}
		})

		t.Run("should handle decomposition errors gracefully", func(t *testing.T) {
			// Arrange
			storage := &mockStorageWithDecomposition{
				tasks: make(map[string]*task.Task),
			}
			detector := decomposition.NewComplexityDetector()
			decomposer := &mockDecomposer{
				err: fmt.Errorf("Claude API timeout"),
			}
			presenter := &mockPresenter{}

			handler := NewPlanningHandlerWithDecomposition(storage, detector, decomposer, presenter)
			input := "organize conference"

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Should fall back to creating single task
			if result.Title != "organize conference" {
				t.Errorf("expected title %q, got %q", "organize conference", result.Title)
			}

			if len(storage.createdTasks) != 1 {
				t.Errorf("expected 1 task created on error, got %d", len(storage.createdTasks))
			}
		})

		t.Run("should handle user selecting 'none'", func(t *testing.T) {
			// Arrange
			storage := &mockStorageWithDecomposition{
				tasks: make(map[string]*task.Task),
			}
			detector := decomposition.NewComplexityDetector()
			decomposer := &mockDecomposer{
				result: &decomposition.DecompositionResult{
					OriginalTask: "organize conference",
					Subtasks: []string{
						"Book venue",
						"Invite speakers",
					},
				},
			}
			presenter := &mockPresenter{
				selectedIndices: []int{}, // User selected 'none'
			}

			handler := NewPlanningHandlerWithDecomposition(storage, detector, decomposer, presenter)
			input := "organize conference"

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Should create only parent task
			if result.Title != "organize conference" {
				t.Errorf("expected title %q, got %q", "organize conference", result.Title)
			}
			if len(storage.createdTasks) != 1 {
				t.Errorf("expected 1 task when user selects none, got %d", len(storage.createdTasks))
			}
		})
	})
}

// Mock types for testing
type mockStorageWithDecomposition struct {
	tasks        map[string]*task.Task
	createdTasks []*task.Task
}

func (m *mockStorageWithDecomposition) Create(t *task.Task) error {
	m.tasks[t.ID] = t
	m.createdTasks = append(m.createdTasks, t)
	return nil
}

func (m *mockStorageWithDecomposition) CreateWithParent(t *task.Task, _ string) error {
	m.tasks[t.ID] = t
	m.createdTasks = append(m.createdTasks, t)
	return nil
}

type mockDecomposer struct {
	called bool
	result *decomposition.DecompositionResult
	err    error
}

func (m *mockDecomposer) Decompose(_ context.Context, _ string) (*decomposition.DecompositionResult, error) {
	m.called = true
	return m.result, m.err
}

type mockPresenter struct {
	called          bool
	selectedIndices []int
}

func (m *mockPresenter) Present(_ *decomposition.DecompositionResult) string {
	m.called = true
	return "mocked presentation"
}

func (m *mockPresenter) ParseSelection(_ string, _ int) ([]int, error) {
	return m.selectedIndices, nil
}
