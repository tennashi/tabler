package decomposition

import (
	"fmt"
	"strings"
	"testing"
)

func TestInteractivePresenter(t *testing.T) {
	t.Run("Present", func(t *testing.T) {
		t.Run("should format decomposition result for display", func(t *testing.T) {
			// Arrange
			result := &DecompositionResult{
				OriginalTask: "organize conference",
				Subtasks: []string{
					"Book venue for conference",
					"Create conference schedule and agenda",
					"Invite speakers and confirm attendance",
					"Setup registration system",
					"Arrange catering and refreshments",
					"Prepare conference materials and badges",
				},
				Rationale: "Task broken down into actionable steps",
			}
			presenter := NewInteractivePresenter()

			// Act
			output := presenter.Present(result)

			// Assert
			// Check that output contains the original task
			if !strings.Contains(output, "organize conference") {
				t.Error("output should contain original task")
			}

			// Check that output contains all subtasks
			for i, subtask := range result.Subtasks {
				expectedLine := formatSubtaskLine(i+1, subtask)
				if !strings.Contains(output, expectedLine) {
					t.Errorf("output should contain subtask %d: %s", i+1, subtask)
				}
			}

			// Check that output contains instructions
			if !strings.Contains(output, "Select subtasks") {
				t.Error("output should contain selection instructions")
			}
		})
	})

	t.Run("ParseSelection", func(t *testing.T) {
		presenter := NewInteractivePresenter()

		tests := []struct {
			name     string
			input    string
			total    int
			expected []int
			wantErr  bool
		}{
			{
				name:     "single selection",
				input:    "1",
				total:    5,
				expected: []int{1},
				wantErr:  false,
			},
			{
				name:     "multiple selections",
				input:    "1,3,5",
				total:    5,
				expected: []int{1, 3, 5},
				wantErr:  false,
			},
			{
				name:     "range selection",
				input:    "1-3",
				total:    5,
				expected: []int{1, 2, 3},
				wantErr:  false,
			},
			{
				name:     "mixed selection",
				input:    "1,3-5",
				total:    5,
				expected: []int{1, 3, 4, 5},
				wantErr:  false,
			},
			{
				name:     "all selection",
				input:    "all",
				total:    3,
				expected: []int{1, 2, 3},
				wantErr:  false,
			},
			{
				name:     "none selection",
				input:    "none",
				total:    3,
				expected: []int{},
				wantErr:  false,
			},
			{
				name:     "out of range",
				input:    "6",
				total:    5,
				expected: nil,
				wantErr:  true,
			},
			{
				name:     "invalid format",
				input:    "abc",
				total:    5,
				expected: nil,
				wantErr:  true,
			},
			{
				name:     "zero index",
				input:    "0",
				total:    5,
				expected: nil,
				wantErr:  true,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Act
				selected, err := presenter.ParseSelection(tt.input, tt.total)

				// Assert
				if tt.wantErr {
					if err == nil {
						t.Error("expected error but got none")
					}
				} else {
					if err != nil {
						t.Errorf("unexpected error: %v", err)
					}
					if !equalIntSlices(selected, tt.expected) {
						t.Errorf("expected %v, got %v", tt.expected, selected)
					}
				}
			})
		}
	})
}

// formatSubtaskLine formats a subtask for display
func formatSubtaskLine(num int, task string) string {
	return fmt.Sprintf("[%d] %s", num, task)
}

// equalIntSlices checks if two int slices are equal
func equalIntSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
