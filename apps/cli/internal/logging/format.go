package logging

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

// FormatSpanText formats a span as human-readable text with indentation
func FormatSpanText(span *Span, depth int) string {
	duration := span.EndTime.Sub(span.StartTime)
	indent := strings.Repeat("  ", depth)

	return fmt.Sprintf("%s%s (%dms)", indent, span.Operation, duration.Milliseconds())
}

// errorJSON represents the JSON structure for error tracking
type errorJSON struct {
	UseCase    string       `json:"use_case"`
	TraceID    string       `json:"trace_id"`
	Timestamp  time.Time    `json:"timestamp"`
	Operation  string       `json:"operation"`
	ErrorType  string       `json:"error_type"`
	Message    string       `json:"message"`
	StackTrace []StackFrame `json:"stack_trace,omitempty"`
}

// FormatErrorJSON formats a TrackedError as JSON for error tracking
func FormatErrorJSON(err *TrackedError) string {
	ej := errorJSON{
		UseCase:   "error_tracking",
		TraceID:   err.TraceID,
		Timestamp: err.Timestamp,
		Operation: err.Operation,
		ErrorType: fmt.Sprintf("%T", err.Err),
		Message:   err.Err.Error(),
	}

	// Add stack trace if available
	if err.StackTrace != nil {
		ej.StackTrace = err.StackTrace
	}

	// Marshal to JSON
	data, _ := json.Marshal(ej)
	return string(data)
}
