package decomposition

import (
	"context"
	"testing"
)

func TestTaskDecomposer(t *testing.T) {
	t.Run("Decompose", func(t *testing.T) {
		t.Run("should decompose complex task into subtasks", func(t *testing.T) {
			// Arrange
			claude := &mockClaudeClient{
				executeFunc: func(_ context.Context, _ string) (string, error) {
					// Simulate Claude response
					return `Here are the subtasks for "organize conference":

1. Book venue for conference
2. Create conference schedule and agenda
3. Invite speakers and confirm attendance
4. Setup registration system
5. Arrange catering and refreshments
6. Prepare conference materials and badges`, nil
				},
			}
			decomposer := NewTaskDecomposer(claude)
			input := "organize conference"

			// Act
			result, err := decomposer.Decompose(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result.Subtasks) != 6 {
				t.Errorf("expected 6 subtasks, got %d", len(result.Subtasks))
			}
			if result.Subtasks[0] != "Book venue for conference" {
				t.Errorf("expected first subtask to be %q, got %q", "Book venue for conference", result.Subtasks[0])
			}
		})
	})
}

// mockClaudeClient for testing
type mockClaudeClient struct {
	executeFunc func(ctx context.Context, prompt string) (string, error)
}

func (m *mockClaudeClient) Execute(ctx context.Context, prompt string) (string, error) {
	if m.executeFunc != nil {
		return m.executeFunc(ctx, prompt)
	}
	return "", nil
}
