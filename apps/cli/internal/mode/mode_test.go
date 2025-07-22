package mode

import (
	"context"
	"testing"

	"github.com/tennashi/tabler/internal/task"
)

// mockHandler is a test implementation of ModeHandler
type mockHandler struct {
	processFunc func(ctx context.Context, input string) (*task.Task, error)
}

func (m *mockHandler) Process(ctx context.Context, input string) (*task.Task, error) {
	if m.processFunc != nil {
		return m.processFunc(ctx, input)
	}
	return nil, nil
}

func TestMode(t *testing.T) {
	t.Run("prefix parsing", func(t *testing.T) {
		t.Run("should parse /quick prefix", func(t *testing.T) {
			// Arrange
			input := "/quick buy milk"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, mode)
			}
			if task != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /q shortcut", func(t *testing.T) {
			// Arrange
			input := "/q buy milk"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, mode)
			}
			if task != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /talk prefix", func(t *testing.T) {
			// Arrange
			input := "/talk prepare for meeting"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != TalkMode {
				t.Errorf("expected mode %v, got %v", TalkMode, mode)
			}
			if task != "prepare for meeting" {
				t.Errorf("expected task %q, got %q", "prepare for meeting", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /t shortcut", func(t *testing.T) {
			// Arrange
			input := "/t prepare for meeting"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != TalkMode {
				t.Errorf("expected mode %v, got %v", TalkMode, mode)
			}
			if task != "prepare for meeting" {
				t.Errorf("expected task %q, got %q", "prepare for meeting", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /plan prefix", func(t *testing.T) {
			// Arrange
			input := "/plan organize conference"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != PlanningMode {
				t.Errorf("expected mode %v, got %v", PlanningMode, mode)
			}
			if task != "organize conference" {
				t.Errorf("expected task %q, got %q", "organize conference", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should parse /p shortcut", func(t *testing.T) {
			// Arrange
			input := "/p organize conference"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != PlanningMode {
				t.Errorf("expected mode %v, got %v", PlanningMode, mode)
			}
			if task != "organize conference" {
				t.Errorf("expected task %q, got %q", "organize conference", task)
			}
			if !hasPrefix {
				t.Errorf("expected hasPrefix to be true, got false")
			}
		})

		t.Run("should return no prefix for regular input", func(t *testing.T) {
			// Arrange
			input := "buy milk"

			// Act
			mode, task, hasPrefix := ParseModePrefix(input)

			// Assert
			if mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, mode)
			}
			if task != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", task)
			}
			if hasPrefix {
				t.Errorf("expected hasPrefix to be false, got true")
			}
		})
	})
}

func TestModeManager(t *testing.T) {
	t.Run("process input", func(t *testing.T) {
		t.Run("should process quick mode input", func(t *testing.T) {
			// Arrange
			manager := NewModeManager()
			input := "/q buy milk"

			// Act
			result, err := manager.ProcessInput(input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Mode != QuickMode {
				t.Errorf("expected mode %v, got %v", QuickMode, result.Mode)
			}
			if result.TaskText != "buy milk" {
				t.Errorf("expected task %q, got %q", "buy milk", result.TaskText)
			}
		})

		t.Run("should use mode handler for processing", func(t *testing.T) {
			// Arrange
			manager := NewModeManager()
			handler := &mockHandler{
				processFunc: func(_ context.Context, input string) (*task.Task, error) {
					return &task.Task{
						ID:    "test-id",
						Title: input,
					}, nil
				},
			}
			manager.RegisterHandler(QuickMode, handler)
			input := "/q buy milk"

			// Act
			result, err := manager.ProcessTask(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.ID != "test-id" {
				t.Errorf("expected task ID %q, got %q", "test-id", result.ID)
			}
			if result.Title != "buy milk" {
				t.Errorf("expected title %q, got %q", "buy milk", result.Title)
			}
		})
	})
}
