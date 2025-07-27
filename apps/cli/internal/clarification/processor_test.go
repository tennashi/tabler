package clarification

import (
	"strings"
	"testing"
)

func TestResponseProcessor(t *testing.T) {
	t.Run("ProcessResponse", func(t *testing.T) {
		t.Run("should update session with response information", func(t *testing.T) {
			// Arrange
			processor := NewResponseProcessor()
			session := &DialogueSession{
				OriginalInput: "prepare for meeting",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: ""},
				},
				ExtractedInfo: make(map[string]string),
			}

			// Act
			err := processor.ProcessResponse(session, "team standup")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			// ExtractedInfo will be updated by ExtractInfo method
		})
	})

	t.Run("ExtractInfo", func(t *testing.T) {
		t.Run("should extract meeting type from response", func(t *testing.T) {
			// Arrange
			processor := NewResponseProcessor()
			session := &DialogueSession{
				OriginalInput: "prepare for meeting",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: "quarterly review"},
				},
				ExtractedInfo: make(map[string]string),
			}

			// Act
			info := processor.ExtractInfo(session)

			// Assert
			if info["type"] != "quarterly review meeting" {
				t.Errorf("expected meeting type to be extracted")
			}
		})

		t.Run("should extract deadline from response", func(t *testing.T) {
			// Arrange
			processor := NewResponseProcessor()
			session := &DialogueSession{
				OriginalInput: "finish the report",
				History: []Exchange{
					{Question: "When do you need it?", Answer: "by Friday"},
				},
				ExtractedInfo: make(map[string]string),
			}

			// Act
			info := processor.ExtractInfo(session)

			// Assert
			if info["deadline"] != "Friday" {
				t.Errorf("expected deadline to be extracted")
			}
		})

		t.Run("should accumulate information across exchanges", func(t *testing.T) {
			// Arrange
			processor := NewResponseProcessor()
			session := &DialogueSession{
				OriginalInput: "do the thing",
				History: []Exchange{
					{Question: "What thing?", Answer: "presentation"},
					{Question: "For whom?", Answer: "the board"},
					{Question: "When?", Answer: "next Tuesday"},
				},
				ExtractedInfo: make(map[string]string),
			}

			// Act
			info := processor.ExtractInfo(session)

			// Assert
			if len(info) < 3 {
				t.Errorf("expected multiple pieces of info, got %d", len(info))
			}
			if info["what"] != "presentation" {
				t.Errorf("expected 'what' to be presentation")
			}
			if info["audience"] != "the board" {
				t.Errorf("expected audience to be the board")
			}
			if info["deadline"] != "next Tuesday" {
				t.Errorf("expected deadline to be next Tuesday")
			}
		})
	})

	t.Run("BuildFinalTask", func(t *testing.T) {
		t.Run("should construct clear task from session", func(t *testing.T) {
			// Arrange
			processor := NewResponseProcessor()
			session := &DialogueSession{
				OriginalInput: "prepare for meeting",
				History: []Exchange{
					{Question: "What kind of meeting?", Answer: "quarterly planning"},
					{Question: "When is it?", Answer: "Thursday 2pm"},
					{Question: "What do you need to prepare?", Answer: "slides and budget"},
				},
				ExtractedInfo: map[string]string{
					"type":      "quarterly planning meeting",
					"deadline":  "Thursday 2pm",
					"materials": "slides and budget",
				},
			}

			// Act
			finalTask := processor.BuildFinalTask(session)

			// Assert
			if !strings.Contains(finalTask, "quarterly planning") {
				t.Error("expected final task to include meeting type")
			}
			if !strings.Contains(finalTask, "Thursday") {
				t.Error("expected final task to include deadline")
			}
			if !strings.Contains(finalTask, "slides") || !strings.Contains(finalTask, "budget") {
				t.Error("expected final task to include materials")
			}
		})

		t.Run("should handle minimal information", func(t *testing.T) {
			// Arrange
			processor := NewResponseProcessor()
			session := &DialogueSession{
				OriginalInput: "work on project",
				History: []Exchange{
					{Question: "Which project?", Answer: "website redesign"},
				},
				ExtractedInfo: map[string]string{
					"project": "website redesign",
				},
			}

			// Act
			finalTask := processor.BuildFinalTask(session)

			// Assert
			if !strings.Contains(finalTask, "website redesign") {
				t.Error("expected final task to include project name")
			}
		})
	})

	t.Run("DetectsSkip", func(t *testing.T) {
		processor := NewResponseProcessor()

		tests := []struct {
			response string
			expected bool
		}{
			{"skip", true},
			{"Skip", true},
			{"SKIP", true},
			{"", true},
			{"  ", true},
			{"quarterly review", false},
			{"I don't know", false},
		}

		for _, tt := range tests {
			t.Run(tt.response, func(t *testing.T) {
				result := processor.DetectsSkip(tt.response)
				if result != tt.expected {
					t.Errorf("for response %q: expected %v, got %v", tt.response, tt.expected, result)
				}
			})
		}
	})
}
