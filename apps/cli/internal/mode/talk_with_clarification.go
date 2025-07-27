package mode

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/tennashi/tabler/internal/clarification"
	"github.com/tennashi/tabler/internal/task"
)

// TalkHandlerWithClarification implements talk mode with dialogue-based clarification
type TalkHandlerWithClarification struct {
	dialogueManager *clarification.DialogueManager
	input           io.Reader
	output          io.Writer
}

// NewTalkHandlerWithClarification creates a new talk handler with clarification
func NewTalkHandlerWithClarification(dialogueManager *clarification.DialogueManager) *TalkHandlerWithClarification {
	return &TalkHandlerWithClarification{
		dialogueManager: dialogueManager,
		input:           os.Stdin,
		output:          os.Stdout,
	}
}

// SetInput sets the input reader (for testing)
func (h *TalkHandlerWithClarification) SetInput(input io.Reader) {
	h.input = input
}

// SetOutput sets the output writer (for testing)
func (h *TalkHandlerWithClarification) SetOutput(output io.Writer) {
	h.output = output
}

// Process creates a task through conversational clarification
func (h *TalkHandlerWithClarification) Process(ctx context.Context, input string) (*task.Task, error) {
	// Start dialogue if needed
	session, err := h.dialogueManager.StartDialogue(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("failed to start dialogue: %w", err)
	}

	// If no dialogue needed (clear input), create task directly
	if session == nil {
		return h.createTask(input), nil
	}

	// Show initial greeting
	_, _ = fmt.Fprintln(h.output, "🤔 I'd like to help clarify this task.")

	// Conduct dialogue
	reader := bufio.NewReader(h.input)

	for !session.IsComplete {
		// Show current question
		_, _ = fmt.Fprintln(h.output, session.CurrentQuestion)
		_, _ = fmt.Fprint(h.output, "> ")

		// Get user response
		response, err := reader.ReadString('\n')
		if err != nil {
			// If error reading, skip dialogue
			session.SkipRequested = true
			break
		}
		response = strings.TrimSpace(response)

		// Process response
		if err := h.dialogueManager.ProcessResponse(ctx, session, response); err != nil {
			return nil, fmt.Errorf("failed to process response: %w", err)
		}

		// Add blank line between exchanges
		if !session.IsComplete {
			_, _ = fmt.Fprintln(h.output)
		}
	}

	// Get final task
	finalTaskTitle := h.dialogueManager.GetFinalTask(session)

	// Show result
	if !session.SkipRequested && finalTaskTitle != input {
		_, _ = fmt.Fprintf(h.output, "\n✅ Got it! Creating task: \"%s\"\n", finalTaskTitle)

		// Show extracted details if any
		if len(session.ExtractedInfo) > 0 {
			if deadline, ok := session.ExtractedInfo["deadline"]; ok {
				_, _ = fmt.Fprintf(h.output, "📅 Deadline: %s\n", deadline)
			}
			if audience, ok := session.ExtractedInfo["audience"]; ok {
				_, _ = fmt.Fprintf(h.output, "👥 For: %s\n", audience)
			}
			if materials, ok := session.ExtractedInfo["materials"]; ok {
				_, _ = fmt.Fprintf(h.output, "📋 Materials: %s\n", materials)
			}
		}
	}

	return h.createTask(finalTaskTitle), nil
}

// createTask creates a new task with the given title
func (h *TalkHandlerWithClarification) createTask(title string) *task.Task {
	now := time.Now()
	return &task.Task{
		ID:        uuid.New().String(),
		Title:     title,
		Priority:  0,
		Completed: false,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
