package logging

import (
	"context"
	"fmt"
	"time"
)

// TrackedError represents an error with tracking context
type TrackedError struct {
	TraceID   string
	Operation string
	Err       error
	Timestamp time.Time
}

// NewTrackedError creates a new tracked error with context
func NewTrackedError(ctx context.Context, operation string, err error) *TrackedError {
	return &TrackedError{
		TraceID:   TraceIDFromContext(ctx),
		Operation: operation,
		Err:       err,
		Timestamp: time.Now(),
	}
}

// Error implements the error interface
func (e *TrackedError) Error() string {
	return fmt.Sprintf("[%s] %v", e.Operation, e.Err)
}

// Unwrap allows errors.Is and errors.As to work
func (e *TrackedError) Unwrap() error {
	return e.Err
}
