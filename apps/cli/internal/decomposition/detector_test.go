package decomposition

import (
	"testing"
)

func TestComplexityDetector(t *testing.T) {
	t.Run("DetectComplexity", func(t *testing.T) {
		tests := []struct {
			name            string
			input           string
			expectedComplex bool
			expectedReason  string
		}{
			{
				name:            "simple task",
				input:           "buy milk",
				expectedComplex: false,
				expectedReason:  "",
			},
			{
				name:            "complex verb - plan",
				input:           "plan company retreat",
				expectedComplex: true,
				expectedReason:  "contains complex verb: plan",
			},
			{
				name:            "complex verb - organize",
				input:           "organize conference",
				expectedComplex: true,
				expectedReason:  "contains complex verb: organize",
			},
			{
				name:            "complex verb - prepare",
				input:           "prepare quarterly report",
				expectedComplex: true,
				expectedReason:  "contains complex verb: prepare",
			},
			{
				name:            "long task",
				input:           "fix all the bugs in the authentication system and add unit tests",
				expectedComplex: true,
				expectedReason:  "task is long and may contain multiple actions",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Arrange
				detector := NewComplexityDetector()

				// Act
				isComplex, reason := detector.DetectComplexity(tt.input)

				// Assert
				if isComplex != tt.expectedComplex {
					t.Errorf("expected complexity %v, got %v", tt.expectedComplex, isComplex)
				}
				if reason != tt.expectedReason {
					t.Errorf("expected reason %q, got %q", tt.expectedReason, reason)
				}
			})
		}
	})
}
