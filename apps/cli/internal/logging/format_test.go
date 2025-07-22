package logging

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
	"time"
)

func TestFormatSpan(t *testing.T) {
	t.Run("FormatSpanText should format span as text", func(t *testing.T) {
		// Arrange
		span := &Span{
			TraceID:   "test-trace-123",
			SpanID:    "test-span-456",
			Operation: "TestOperation",
			StartTime: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 12, 0, 0, 100000000, time.UTC), // 100ms later
		}

		// Act
		result := FormatSpanText(span, 0)

		// Assert
		if !strings.Contains(result, "TestOperation") {
			t.Errorf("expected result to contain operation name, got %q", result)
		}
		if !strings.Contains(result, "100ms") {
			t.Errorf("expected result to contain duration, got %q", result)
		}
	})

	t.Run("FormatSpanText should indent based on depth", func(t *testing.T) {
		// Arrange
		span := &Span{
			TraceID:   "test-trace-123",
			SpanID:    "test-span-456",
			Operation: "NestedOperation",
			StartTime: time.Date(2024, 1, 1, 12, 0, 0, 0, time.UTC),
			EndTime:   time.Date(2024, 1, 1, 12, 0, 0, 50000000, time.UTC),
		}

		// Act
		result := FormatSpanText(span, 2)

		// Assert
		if !strings.HasPrefix(result, "    ") { // 2 levels = 4 spaces
			t.Errorf("expected result to start with 4 spaces, got %q", result)
		}
		if !strings.Contains(result, "NestedOperation") {
			t.Errorf("expected result to contain operation name, got %q", result)
		}
	})
}

func TestFormatErrorJSON(t *testing.T) {
	t.Run("FormatErrorJSON should output valid JSON", func(t *testing.T) {
		// Arrange
		ctx := context.Background()
		ctx = WithTraceID(ctx, "test-trace-123")
		baseErr := errors.New("something went wrong")
		trackedErr := NewTrackedError(ctx, "TestOperation", baseErr)

		// Act
		jsonStr := FormatErrorJSON(trackedErr)

		// Assert
		if jsonStr == "" {
			t.Fatal("JSON output should not be empty")
		}

		// Verify it's valid JSON
		var result map[string]interface{}
		err := json.Unmarshal([]byte(jsonStr), &result)
		if err != nil {
			t.Fatalf("should produce valid JSON: %v", err)
		}

		// Check required fields
		if result["use_case"] != "error_tracking" {
			t.Errorf("expected use_case to be error_tracking, got %v", result["use_case"])
		}
		if result["trace_id"] != "test-trace-123" {
			t.Errorf("expected trace_id to be test-trace-123, got %v", result["trace_id"])
		}
		if result["operation"] != "TestOperation" {
			t.Errorf("expected operation to be TestOperation, got %v", result["operation"])
		}
	})
}
