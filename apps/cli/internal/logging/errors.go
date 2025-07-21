package logging

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"time"
)

// TrackedError represents an error with tracking context
type TrackedError struct {
	TraceID    string
	Operation  string
	Err        error
	Timestamp  time.Time
	StackTrace []StackFrame
}

// NewTrackedError creates a new tracked error with context
func NewTrackedError(ctx context.Context, operation string, err error) *TrackedError {
	te := &TrackedError{
		TraceID:   TraceIDFromContext(ctx),
		Operation: operation,
		Err:       err,
		Timestamp: time.Now(),
	}

	// Capture stack trace if enabled
	if os.Getenv("TABLER_ERROR_STACK") == "1" {
		te.StackTrace = CaptureStackTrace(1) // Skip NewTrackedError frame
	}

	return te
}

// Error implements the error interface
func (e *TrackedError) Error() string {
	return fmt.Sprintf("[%s] %v", e.Operation, e.Err)
}

// Unwrap allows errors.Is and errors.As to work
func (e *TrackedError) Unwrap() error {
	return e.Err
}

// StackFrame represents a single frame in a stack trace
type StackFrame struct {
	Function string
	File     string
	Line     int
}

// CaptureStackTrace captures the current stack trace, skipping the specified number of frames
func CaptureStackTrace(skip int) []StackFrame {
	var frames []StackFrame

	// Start from the caller of CaptureStackTrace
	pc := make([]uintptr, 32)
	n := runtime.Callers(skip+2, pc)

	for i := 0; i < n; i++ {
		fn := runtime.FuncForPC(pc[i])
		if fn == nil {
			continue
		}

		file, line := fn.FileLine(pc[i])
		frames = append(frames, StackFrame{
			Function: fn.Name(),
			File:     file,
			Line:     line,
		})
	}

	return frames
}
