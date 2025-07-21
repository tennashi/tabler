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

func TestSpan(t *testing.T) {
	t.Run("NewSpan should create span with context", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		ctx = WithTraceID(ctx, "test-trace-123")

		// Act
		span := NewSpan(ctx, "test-operation")

		// Assert
		if span == nil {
			t.Fatal("span should not be nil")
		}
		if span.TraceID != "test-trace-123" {
			t.Errorf("expected trace ID %q, got %q", "test-trace-123", span.TraceID)
		}
		if span.Operation != "test-operation" {
			t.Errorf("expected operation %q, got %q", "test-operation", span.Operation)
		}
		if span.StartTime.IsZero() {
			t.Error("start time should be set")
		}
	})

	t.Run("Trace should return function to end trace", func(t *testing.T) {
		// Arrange
		t.Setenv("TABLER_TRACE", "1") // Enable tracing for this test
		ctx := context.Background()
		ctx = WithTraceID(ctx, "test-trace-123")
		var capturedSpan *Span

		// Override output for testing
		originalOutput := spanOutput
		spanOutput = func(span *Span) {
			capturedSpan = span
		}
		defer func() { spanOutput = originalOutput }()

		// Act
		endFunc := Trace(ctx, "test-operation")
		// No sleep - we'll just verify that end time is set
		endFunc()

		// Assert
		if capturedSpan == nil {
			t.Fatal("span should have been output")
		}
		if capturedSpan.Operation != "test-operation" {
			t.Errorf("expected operation %q, got %q", "test-operation", capturedSpan.Operation)
		}
		if capturedSpan.EndTime.IsZero() {
			t.Error("end time should be set")
		}
		// EndTime should be equal to or after StartTime
		if capturedSpan.EndTime.Before(capturedSpan.StartTime) {
			t.Error("end time should not be before start time")
		}
	})
}

func TestTraceEnvironmentControl(t *testing.T) {
	t.Run("IsTraceEnabled should return false when TABLER_TRACE is not set", func(t *testing.T) {
		// Arrange
		t.Setenv("TABLER_TRACE", "")

		// Act
		enabled := IsTraceEnabled()

		// Assert
		if enabled {
			t.Error("trace should be disabled when TABLER_TRACE is not set")
		}
	})

	t.Run("IsTraceEnabled should return true when TABLER_TRACE is 1", func(t *testing.T) {
		// Arrange
		t.Setenv("TABLER_TRACE", "1")

		// Act
		enabled := IsTraceEnabled()

		// Assert
		if !enabled {
			t.Error("trace should be enabled when TABLER_TRACE is 1")
		}
	})

	t.Run("Trace should not output when tracing is disabled", func(t *testing.T) {
		// Arrange
		t.Setenv("TABLER_TRACE", "")
		ctx := context.Background()
		ctx = WithTraceID(ctx, "test-trace-123")

		outputCalled := false
		originalOutput := spanOutput
		spanOutput = func(_ *Span) {
			outputCalled = true
		}
		defer func() { spanOutput = originalOutput }()

		// Act
		endFunc := Trace(ctx, "test-operation")
		endFunc()

		// Assert
		if outputCalled {
			t.Error("spanOutput should not be called when tracing is disabled")
		}
	})
}
