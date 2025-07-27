package mode

import (
	"context"
	"io"
	"strings"
	"testing"

	"github.com/tennashi/tabler/internal/clarification"
)

func TestTalkHandlerWithClarification(t *testing.T) {
	t.Run("should use clarification dialogue for vague input", func(t *testing.T) {
		// Arrange
		// Create mock components
		claude := &mockClaudeForClarification{
			responses: []string{
				"What specific thing do you need to work on?",
				"When do you need to complete it?",
				"COMPLETE",
			},
		}

		detector := clarification.NewVaguenessDetector()
		questionGen := clarification.NewQuestionGenerator(claude)
		processor := clarification.NewResponseProcessor()
		dialogueManager := clarification.NewDialogueManager(detector, questionGen, processor)

		// Create handler with clarification
		handler := NewTalkHandlerWithClarification(dialogueManager)

		// Simulate user responses
		userInput := strings.NewReader("presentation\nFriday\n")
		handler.SetInput(userInput)
		handler.SetOutput(io.Discard) // Suppress output for testing

		// Act
		task, err := handler.Process(context.Background(), "work on the thing")
		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if task == nil {
			t.Fatal("expected task to be created")
		}

		// The final task should be more specific than the original
		if task.Title == "work on the thing" {
			t.Error("expected task title to be clarified")
		}
		if !strings.Contains(strings.ToLower(task.Title), "presentation") {
			t.Error("expected task to include clarified information")
		}
	})

	t.Run("should skip dialogue for clear input", func(t *testing.T) {
		// Arrange
		claude := &mockClaudeForClarification{}
		detector := clarification.NewVaguenessDetector()
		questionGen := clarification.NewQuestionGenerator(claude)
		processor := clarification.NewResponseProcessor()
		dialogueManager := clarification.NewDialogueManager(detector, questionGen, processor)

		handler := NewTalkHandlerWithClarification(dialogueManager)
		handler.SetOutput(io.Discard)

		// Act
		clearInput := "review Q4 budget report by Friday"
		task, err := handler.Process(context.Background(), clearInput)
		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if task.Title != clearInput {
			t.Errorf("expected task title to remain unchanged for clear input")
		}

		// Claude should not have been called
		if claude.callCount > 0 {
			t.Error("expected no Claude calls for clear input")
		}
	})

	t.Run("should handle skip request", func(t *testing.T) {
		// Arrange
		claude := &mockClaudeForClarification{
			responses: []string{
				"What kind of meeting do you need to prepare for?",
			},
		}

		detector := clarification.NewVaguenessDetector()
		questionGen := clarification.NewQuestionGenerator(claude)
		processor := clarification.NewResponseProcessor()
		dialogueManager := clarification.NewDialogueManager(detector, questionGen, processor)

		handler := NewTalkHandlerWithClarification(dialogueManager)

		// User types "skip"
		userInput := strings.NewReader("skip\n")
		handler.SetInput(userInput)
		handler.SetOutput(io.Discard)

		// Act
		task, err := handler.Process(context.Background(), "prepare for meeting")
		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		// Task should be created with original input
		if task.Title != "prepare for meeting" {
			t.Errorf("expected original title when skipped, got %q", task.Title)
		}
	})
}

// mockClaudeForClarification provides canned responses for testing
type mockClaudeForClarification struct {
	responses []string
	callCount int
}

func (m *mockClaudeForClarification) Execute(_ context.Context, _ string) (string, error) {
	if m.callCount < len(m.responses) {
		response := m.responses[m.callCount]
		m.callCount++
		return response, nil
	}
	return "COMPLETE", nil
}
