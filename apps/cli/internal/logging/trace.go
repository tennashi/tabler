package logging

import (
	"context"

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
