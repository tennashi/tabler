package logging

import (
	"context"
	"errors"
	"strings"
	"testing"
)

func TestTrackedError(t *testing.T) {
	t.Run("NewTrackedError should create error with context", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		ctx = WithTraceID(ctx, "test-trace-123")
		baseErr := errors.New("something went wrong")

		// Act
		trackedErr := NewTrackedError(ctx, "TestOperation", baseErr)

		// Assert
		if trackedErr == nil {
			t.Fatal("tracked error should not be nil")
		}
		if trackedErr.TraceID != "test-trace-123" {
			t.Errorf("expected trace ID %q, got %q", "test-trace-123", trackedErr.TraceID)
		}
		if trackedErr.Operation != "TestOperation" {
			t.Errorf("expected operation %q, got %q", "TestOperation", trackedErr.Operation)
		}
		if !errors.Is(trackedErr.Err, baseErr) {
			t.Error("expected base error to be preserved")
		}
	})

	t.Run("TrackedError should implement error interface", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		baseErr := errors.New("base error")

		// Act
		trackedErr := NewTrackedError(ctx, "TestOp", baseErr)

		// Assert
		errMsg := trackedErr.Error()
		if !strings.Contains(errMsg, "base error") {
			t.Errorf("error message should contain base error, got %q", errMsg)
		}
		if !strings.Contains(errMsg, "TestOp") {
			t.Errorf("error message should contain operation, got %q", errMsg)
		}
	})

	t.Run("TrackedError should support errors.Is", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		baseErr := errors.New("base error")

		// Act
		trackedErr := NewTrackedError(ctx, "TestOp", baseErr)

		// Assert
		if !errors.Is(trackedErr, baseErr) {
			t.Error("errors.Is should work with tracked error")
		}
	})
}

func TestStackTrace(t *testing.T) {
	t.Run("CaptureStackTrace should return stack frames", func(t *testing.T) {
		// Act
		frames := CaptureStackTrace(0)

		// Assert
		if len(frames) == 0 {
			t.Fatal("stack trace should not be empty")
		}

		// First frame should be this test function
		firstFrame := frames[0]
		if !strings.Contains(firstFrame.Function, "TestStackTrace") {
			t.Errorf("expected first frame to be test function, got %q", firstFrame.Function)
		}
		if firstFrame.File == "" {
			t.Error("frame should have file path")
		}
		if firstFrame.Line == 0 {
			t.Error("frame should have line number")
		}
	})

	t.Run("NewTrackedError should capture stack trace when enabled", func(t *testing.T) {
		// Arrange
		t.Setenv("TABLER_ERROR_STACK", "1")
		ctx := context.Background()
		baseErr := errors.New("test error")

		// Act
		trackedErr := NewTrackedError(ctx, "TestOp", baseErr)

		// Assert
		if trackedErr.StackTrace == nil {
			t.Fatal("stack trace should be captured when enabled")
		}
		if len(trackedErr.StackTrace) == 0 {
			t.Error("stack trace should not be empty")
		}
	})

	t.Run("NewTrackedError should not capture stack trace when disabled", func(t *testing.T) {
		// Arrange
		t.Setenv("TABLER_ERROR_STACK", "")
		ctx := context.Background()
		baseErr := errors.New("test error")

		// Act
		trackedErr := NewTrackedError(ctx, "TestOp", baseErr)

		// Assert
		if trackedErr.StackTrace != nil {
			t.Error("stack trace should not be captured when disabled")
		}
	})
}
