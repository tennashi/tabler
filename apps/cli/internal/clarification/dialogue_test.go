package clarification

import (
	"context"
	"strings"
	"testing"
)

func TestDialogueManager(t *testing.T) {
	t.Run("StartDialogue", func(t *testing.T) {
		t.Run("should initiate dialogue for vague input", func(t *testing.T) {
			// Arrange
			questionGen := &mockQuestionGenerator{
				question: "What kind of meeting do you need to prepare for?",
			}
			processor := &mockResponseProcessor{}
			detector := NewVaguenessDetector()

			manager := NewDialogueManager(detector, questionGen, processor)
			input := "prepare for meeting"

			// Act
			dialogue, err := manager.StartDialogue(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dialogue == nil {
				t.Fatal("expected dialogue to be created")
			}
			if dialogue.OriginalInput != input {
				t.Errorf("expected original input %q, got %q", input, dialogue.OriginalInput)
			}
			if dialogue.CurrentQuestion == "" {
				t.Error("expected current question to be set")
			}
		})

		t.Run("should skip dialogue for clear input", func(t *testing.T) {
			// Arrange
			questionGen := &mockQuestionGenerator{}
			processor := &mockResponseProcessor{}
			detector := NewVaguenessDetector()

			manager := NewDialogueManager(detector, questionGen, processor)
			input := "send Q4 budget report to john@example.com by Friday"

			// Act
			dialogue, err := manager.StartDialogue(context.Background(), input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dialogue != nil {
				t.Error("expected no dialogue for clear input")
			}
		})
	})

	t.Run("ProcessResponse", func(t *testing.T) {
		t.Run("should handle user response and generate next question", func(t *testing.T) {
			// Arrange
			questionGen := &mockQuestionGenerator{
				question: "When is the deadline?",
			}
			processor := &mockResponseProcessor{
				updatedContext: map[string]string{
					"meeting_type": "team standup",
				},
			}
			detector := NewVaguenessDetector()

			manager := NewDialogueManager(detector, questionGen, processor)
			dialogue := &DialogueSession{
				OriginalInput:   "prepare for meeting",
				CurrentQuestion: "What kind of meeting?",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: ""},
				},
			}

			// Act
			err := manager.ProcessResponse(context.Background(), dialogue, "team standup")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if dialogue.History[0].Answer != "team standup" {
				t.Errorf("expected answer to be recorded")
			}
			if dialogue.CurrentQuestion != "When is the deadline?" {
				t.Errorf("expected new question, got %q", dialogue.CurrentQuestion)
			}
		})

		t.Run("should detect skip intent", func(t *testing.T) {
			// Arrange
			questionGen := &mockQuestionGenerator{}
			processor := &mockResponseProcessor{}
			detector := NewVaguenessDetector()

			manager := NewDialogueManager(detector, questionGen, processor)
			dialogue := &DialogueSession{
				OriginalInput:   "do the thing",
				CurrentQuestion: "What thing?",
				History:         []Exchange{},
			}

			// Act
			err := manager.ProcessResponse(context.Background(), dialogue, "skip")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !dialogue.SkipRequested {
				t.Error("expected skip to be detected")
			}
		})

		t.Run("should complete after exchange limit", func(t *testing.T) {
			// Arrange
			questionGen := &mockQuestionGenerator{
				complete: true,
			}
			processor := &mockResponseProcessor{}
			detector := NewVaguenessDetector()

			manager := NewDialogueManager(detector, questionGen, processor)
			dialogue := &DialogueSession{
				OriginalInput: "task",
				History: []Exchange{
					{Question: "Q1", Answer: "A1"},
					{Question: "Q2", Answer: "A2"},
				},
			}

			// Act
			err := manager.ProcessResponse(context.Background(), dialogue, "answer 3")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !dialogue.IsComplete {
				t.Error("expected dialogue to be complete after 3 exchanges")
			}
		})
	})

	t.Run("GetFinalTask", func(t *testing.T) {
		t.Run("should construct final task from dialogue", func(t *testing.T) {
			// Arrange
			questionGen := &mockQuestionGenerator{}
			processor := &mockResponseProcessor{
				finalTask: "Prepare slides for Q4 planning team standup on Thursday",
			}
			detector := NewVaguenessDetector()

			manager := NewDialogueManager(detector, questionGen, processor)
			dialogue := &DialogueSession{
				OriginalInput: "prepare for meeting",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: "team standup"},
					{Question: "What's the topic?", Answer: "Q4 planning"},
					{Question: "When is it?", Answer: "Thursday"},
				},
				ExtractedInfo: map[string]string{
					"meeting_type": "team standup",
					"topic":        "Q4 planning",
					"deadline":     "Thursday",
				},
				IsComplete: true,
			}

			// Act
			finalTask := manager.GetFinalTask(dialogue)

			// Assert
			if !strings.Contains(finalTask, "Q4 planning") {
				t.Errorf("expected final task to include topic")
			}
			if !strings.Contains(finalTask, "Thursday") {
				t.Errorf("expected final task to include deadline")
			}
		})
	})
}

// Mock types for testing
type mockQuestionGenerator struct {
	question string
	complete bool
	err      error
}

func (m *mockQuestionGenerator) GenerateQuestion(_ context.Context, _ *DialogueSession) (string, bool, error) {
	return m.question, m.complete, m.err
}

type mockResponseProcessor struct {
	updatedContext map[string]string
	finalTask      string
}

func (m *mockResponseProcessor) ProcessResponse(_ *DialogueSession, _ string) error {
	return nil
}

func (m *mockResponseProcessor) ExtractInfo(_ *DialogueSession) map[string]string {
	return m.updatedContext
}

func (m *mockResponseProcessor) BuildFinalTask(_ *DialogueSession) string {
	return m.finalTask
}

func (m *mockResponseProcessor) DetectsSkip(response string) bool {
	return response == "skip" || response == ""
}
