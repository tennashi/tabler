package mode

import (
	"context"
	"testing"
)

func TestQuickHandler(t *testing.T) {
	t.Run("process", func(t *testing.T) {
		t.Run("should create task with minimal processing", func(t *testing.T) {
			// Arrange
			handler := NewQuickHandler()
			input := "buy milk #shopping @tomorrow !2"

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Title != "buy milk #shopping @tomorrow !2" {
				t.Errorf("expected title %q, got %q", "buy milk #shopping @tomorrow !2", result.Title)
			}
		})
	})
}
