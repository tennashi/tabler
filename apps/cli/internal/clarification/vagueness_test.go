package clarification

import (
	"testing"
)

func TestVaguenessDetector(t *testing.T) {
	t.Run("DetectVagueness", func(t *testing.T) {
		tests := []struct {
			name          string
			input         string
			expectedVague bool
			expectedScore float64
		}{
			{
				name:          "very vague task",
				input:         "do the thing",
				expectedVague: true,
				expectedScore: 0.9,
			},
			{
				name:          "somewhat vague task",
				input:         "prepare for meeting",
				expectedVague: true,
				expectedScore: 0.7,
			},
			{
				name:          "clear task with details",
				input:         "review Q4 budget report by Friday",
				expectedVague: false,
				expectedScore: 0.2,
			},
			{
				name:          "specific task with context",
				input:         "send email to john about project status",
				expectedVague: false,
				expectedScore: 0.3,
			},
			{
				name:          "very short vague input",
				input:         "stuff",
				expectedVague: true,
				expectedScore: 1.0,
			},
			{
				name:          "question as task",
				input:         "what about the deadline?",
				expectedVague: true,
				expectedScore: 0.8,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				// Arrange
				detector := NewVaguenessDetector()

				// Act
				isVague, score := detector.DetectVagueness(tt.input)

				// Assert
				if isVague != tt.expectedVague {
					t.Errorf("expected vague=%v, got %v", tt.expectedVague, isVague)
				}

				// Allow some tolerance for score
				tolerance := 0.3
				if score < tt.expectedScore-tolerance || score > tt.expectedScore+tolerance {
					t.Errorf("expected score around %.1f, got %.1f", tt.expectedScore, score)
				}
			})
		}
	})

	t.Run("vagueness indicators", func(t *testing.T) {
		t.Run("should detect generic words", func(t *testing.T) {
			// Arrange
			detector := NewVaguenessDetector()
			inputs := []string{
				"handle stuff",
				"do things",
				"work on something",
				"fix it",
			}

			// Act & Assert
			for _, input := range inputs {
				isVague, _ := detector.DetectVagueness(input)
				if !isVague {
					t.Errorf("expected %q to be detected as vague", input)
				}
			}
		})

		t.Run("should detect missing context", func(t *testing.T) {
			// Arrange
			detector := NewVaguenessDetector()
			
			// Act
			isVague, _ := detector.DetectVagueness("send the report")
			
			// Assert
			if !isVague {
				t.Error("expected missing context to be detected as vague")
			}
		})
	})
}