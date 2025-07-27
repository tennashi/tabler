package clarification

import (
	"context"
	"strings"
	"testing"
)

func TestQuestionGenerator(t *testing.T) {
	t.Run("GenerateQuestion", func(t *testing.T) {
		t.Run("should generate first question for vague input", func(t *testing.T) {
			// Arrange
			claude := &mockClaudeClient{
				response: "What specific thing do you need to work on?",
			}
			generator := NewQuestionGenerator(claude)
			session := &DialogueSession{
				OriginalInput: "work on the thing",
				History:       []Exchange{},
				ExtractedInfo: make(map[string]string),
			}

			// Act
			question, isComplete, err := generator.GenerateQuestion(context.Background(), session)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if isComplete {
				t.Error("expected dialogue to continue")
			}
			if question == "" {
				t.Error("expected question to be generated")
			}
			if !strings.Contains(question, "?") {
				t.Error("expected question to contain question mark")
			}
		})

		t.Run("should generate follow-up question based on history", func(t *testing.T) {
			// Arrange
			claude := &mockClaudeClient{
				response: "When do you need to complete the presentation?",
			}
			generator := NewQuestionGenerator(claude)
			session := &DialogueSession{
				OriginalInput: "prepare for meeting",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: "team presentation"},
				},
				ExtractedInfo: map[string]string{
					"meeting_type": "team presentation",
				},
			}

			// Act
			question, isComplete, err := generator.GenerateQuestion(context.Background(), session)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if isComplete {
				t.Error("expected dialogue to continue")
			}
			if !strings.Contains(strings.ToLower(question), "when") || !strings.Contains(strings.ToLower(question), "complete") {
				t.Errorf("expected question about timing, got: %s", question)
			}
		})

		t.Run("should detect when enough information is gathered", func(t *testing.T) {
			// Arrange
			claude := &mockClaudeClient{
				response: "COMPLETE",
			}
			generator := NewQuestionGenerator(claude)
			session := &DialogueSession{
				OriginalInput: "prepare for meeting",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: "quarterly review"},
					{Question: "When is it?", Answer: "next Friday"},
					{Question: "What materials do you need?", Answer: "slides and budget report"},
				},
				ExtractedInfo: map[string]string{
					"meeting_type": "quarterly review",
					"deadline":     "next Friday",
					"materials":    "slides and budget report",
				},
			}

			// Act
			_, isComplete, err := generator.GenerateQuestion(context.Background(), session)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !isComplete {
				t.Error("expected dialogue to be complete")
			}
		})

		t.Run("should handle Claude errors gracefully", func(t *testing.T) {
			// Arrange
			claude := &mockClaudeClient{
				err: context.DeadlineExceeded,
			}
			generator := NewQuestionGenerator(claude)
			session := &DialogueSession{
				OriginalInput: "do stuff",
				History:       []Exchange{},
			}

			// Act
			question, isComplete, err := generator.GenerateQuestion(context.Background(), session)

			// Assert
			if err == nil {
				t.Error("expected error to be returned")
			}
			if question != "" {
				t.Error("expected no question on error")
			}
			if isComplete {
				t.Error("expected not complete on error")
			}
		})
	})
}

// mockClaudeClient for testing
type mockClaudeClient struct {
	response string
	err      error
}

func (m *mockClaudeClient) Execute(_ context.Context, _ string) (string, error) {
	return m.response, m.err
}
