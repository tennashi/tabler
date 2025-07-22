package mode

import (
	"context"
	"testing"
)

func TestPlanningHandler(t *testing.T) {
	t.Run("process", func(t *testing.T) {
		t.Run("should create task with decomposition planning", func(t *testing.T) {
			// Arrange
			handler := NewPlanningHandler()
			input := "organize conference"

			// Act
			result, err := handler.Process(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			// For now, planning mode just creates task like quick mode
			// Will be enhanced when smart decomposition is implemented
			if result.Title != "organize conference" {
				t.Errorf("expected title %q, got %q", "organize conference", result.Title)
			}
		})
	})
}
