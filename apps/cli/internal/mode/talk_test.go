package mode

import (
	"context"
	"testing"
)

func TestTalkHandler(t *testing.T) {
	t.Run("process", func(t *testing.T) {
		t.Run("should create task through conversational flow", func(t *testing.T) {
			// Arrange
			handler := NewTalkHandler()
			input := "prepare for meeting"

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			// For now, talk mode just creates task like quick mode
			// Will be enhanced when interactive clarification is implemented
			if result.Title != "prepare for meeting" {
				t.Errorf("expected title %q, got %q", "prepare for meeting", result.Title)
			}
		})
	})
}
