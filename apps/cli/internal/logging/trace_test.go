package logging

import (
	"context"
	"testing"
)

func TestTraceID(t *testing.T) {
	t.Run("WithTraceID should store trace ID in context", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		traceID := "test-trace-123"

		// Act
		newCtx := WithTraceID(ctx, traceID)

		// Assert
		got := TraceIDFromContext(newCtx)
		if got != traceID {
			t.Errorf("expected trace ID %q, got %q", traceID, got)
		}
	})

	t.Run("NewTraceID should generate unique trace ID", func(t *testing.T) {
		// Act
		id1 := NewTraceID()
		id2 := NewTraceID()

		// Assert
		if id1 == "" {
			t.Error("trace ID should not be empty")
		}
		if id2 == "" {
			t.Error("trace ID should not be empty")
		}
		if id1 == id2 {
			t.Error("trace IDs should be unique")
		}
	})
}
