package logging

import (
	"fmt"
	"strings"
)

// FormatSpanText formats a span as human-readable text with indentation
func FormatSpanText(span *Span, depth int) string {
	duration := span.EndTime.Sub(span.StartTime)
	indent := strings.Repeat("  ", depth)

	return fmt.Sprintf("%s%s (%dms)", indent, span.Operation, duration.Milliseconds())
}
