package logging

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

type contextKey struct{}

var traceIDKey = contextKey{}

// WithTraceID stores a trace ID in the context
func WithTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, traceIDKey, traceID)
}

// TraceIDFromContext retrieves the trace ID from the context
func TraceIDFromContext(ctx context.Context) string {
	if id, ok := ctx.Value(traceIDKey).(string); ok {
		return id
	}
	return ""
}

// NewTraceID generates a new unique trace ID
func NewTraceID() string {
	return uuid.New().String()
}

// Span represents a single operation in a trace
type Span struct {
	TraceID   string
	SpanID    string
	Operation string
	StartTime time.Time
	EndTime   time.Time
}

// NewSpan creates a new span with the given operation name
func NewSpan(ctx context.Context, operation string) *Span {
	return &Span{
		TraceID:   TraceIDFromContext(ctx),
		SpanID:    uuid.New().String(),
		Operation: operation,
		StartTime: time.Now(),
	}
}

// spanOutput is a variable for testing purposes
var spanOutput = func(_ *Span) {
	// Default implementation will be added later
}

// Trace starts a new trace span and returns a function to end it
func Trace(ctx context.Context, operation string) func() {
	if !IsTraceEnabled() {
		return func() {} // no-op
	}

	span := NewSpan(ctx, operation)

	return func() {
		span.EndTime = time.Now()
		spanOutput(span)
	}
}

// IsTraceEnabled checks if tracing is enabled via environment variable
func IsTraceEnabled() bool {
	return os.Getenv("TABLER_TRACE") == "1"
}

// InitializeOutput sets up the default span output function
func InitializeOutput() {
	spanOutput = func(span *Span) {
		if IsTraceEnabled() {
			fmt.Fprintln(os.Stderr, FormatSpanText(span, 0))
		}
	}
}
