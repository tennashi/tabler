package logging

import (
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
