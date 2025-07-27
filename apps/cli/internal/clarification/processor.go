package clarification

import (
	"fmt"
	"strings"
)

// ResponseProcessorImpl extracts information from user responses
type ResponseProcessorImpl struct{}

// NewResponseProcessor creates a new response processor
func NewResponseProcessor() *ResponseProcessorImpl {
	return &ResponseProcessorImpl{}
}

// ProcessResponse updates the session based on user response
func (p *ResponseProcessorImpl) ProcessResponse(_ *DialogueSession, _ string) error {
	// Response is already recorded in History by DialogueManager
	// This method is for any additional processing if needed
	return nil
}

// ExtractInfo analyzes dialogue history to extract structured information
func (p *ResponseProcessorImpl) ExtractInfo(session *DialogueSession) map[string]string {
	info := make(map[string]string)

	// Copy existing extracted info
	for k, v := range session.ExtractedInfo {
		info[k] = v
	}

	// Analyze each exchange
	for _, exchange := range session.History {
		if exchange.Answer == "" {
			continue
		}

		questionLower := strings.ToLower(exchange.Question)
		answer := strings.TrimSpace(exchange.Answer)

		// Extract based on question patterns
		switch {
		case strings.Contains(questionLower, "what kind") || strings.Contains(questionLower, "what type"):
			// Extract type information
			if strings.Contains(questionLower, "meeting") {
				info["type"] = answer + " meeting"
			} else {
				info["type"] = answer
			}

		case strings.Contains(questionLower, "what") && !strings.Contains(questionLower, "what kind"):
			// Extract what/subject
			info["what"] = answer

		case strings.Contains(questionLower, "when") || strings.Contains(questionLower, "deadline"):
			// Extract timing
			info["deadline"] = p.normalizeDeadline(answer)

		case strings.Contains(questionLower, "who") || strings.Contains(questionLower, "whom"):
			// Extract audience/recipient
			info["audience"] = answer

		case strings.Contains(questionLower, "which"):
			// Extract specific selection
			if strings.Contains(questionLower, "project") {
				info["project"] = answer
			} else {
				info["selection"] = answer
			}

		case strings.Contains(questionLower, "prepare") || strings.Contains(questionLower, "need"):
			// Extract materials/requirements
			info["materials"] = answer
		}
	}

	return info
}

// normalizeDeadline cleans up deadline responses
func (p *ResponseProcessorImpl) normalizeDeadline(deadline string) string {
	deadline = strings.TrimSpace(deadline)

	// Remove common prefixes
	prefixes := []string{"by ", "on ", "at ", "before "}
	for _, prefix := range prefixes {
		if strings.HasPrefix(strings.ToLower(deadline), prefix) {
			deadline = deadline[len(prefix):]
		}
	}

	return deadline
}

// BuildFinalTask constructs a clear task from the gathered information
func (p *ResponseProcessorImpl) BuildFinalTask(session *DialogueSession) string {
	info := p.ExtractInfo(session)

	// Start building the task description
	var taskBuilder strings.Builder

	// Determine the main action based on context
	originalLower := strings.ToLower(session.OriginalInput)

	// Handle different types of tasks
	if meetingType, ok := info["type"]; ok && strings.Contains(meetingType, "meeting") {
		// Meeting preparation
		taskBuilder.WriteString(fmt.Sprintf("Prepare for %s", meetingType))

		// Add materials if specified
		if materials, ok := info["materials"]; ok {
			taskBuilder.WriteString(fmt.Sprintf(" - prepare %s", materials))
		}

		// Add deadline
		if deadline, ok := info["deadline"]; ok {
			taskBuilder.WriteString(fmt.Sprintf(" by %s", deadline))
		}
	} else if project, ok := info["project"]; ok {
		// Project work
		taskBuilder.WriteString(fmt.Sprintf("Work on %s", project))

		// Add specifics if available
		if what, ok := info["what"]; ok {
			taskBuilder.WriteString(fmt.Sprintf(" - %s", what))
		}
	} else if what, ok := info["what"]; ok {
		// General task with specific item
		switch {
		case strings.Contains(originalLower, "prepar"):
			taskBuilder.WriteString(fmt.Sprintf("Prepare %s", what))
		case strings.Contains(originalLower, "creat") || strings.Contains(originalLower, "make"):
			taskBuilder.WriteString(fmt.Sprintf("Create %s", what))
		default:
			taskBuilder.WriteString(fmt.Sprintf("Complete %s", what))
		}

		// Add audience if specified
		if audience, ok := info["audience"]; ok {
			taskBuilder.WriteString(fmt.Sprintf(" for %s", audience))
		}

		// Add deadline
		if deadline, ok := info["deadline"]; ok {
			taskBuilder.WriteString(fmt.Sprintf(" by %s", deadline))
		}
	} else {
		// Fallback: improve original with any extracted info
		taskBuilder.WriteString(p.improveOriginal(session.OriginalInput))

		// Add any deadline
		if deadline, ok := info["deadline"]; ok {
			taskBuilder.WriteString(fmt.Sprintf(" by %s", deadline))
		}
	}

	return taskBuilder.String()
}

// improveOriginal makes the original input slightly better
func (p *ResponseProcessorImpl) improveOriginal(original string) string {
	// Simple improvements
	replacements := map[string]string{
		"do the thing":  "Complete the task",
		"work on stuff": "Work on project",
		"prepare":       "Prepare materials",
		"fix it":        "Fix the issue",
		"handle it":     "Handle the request",
	}

	lower := strings.ToLower(strings.TrimSpace(original))
	if improved, ok := replacements[lower]; ok {
		return improved
	}

	// Capitalize first letter
	if len(original) > 0 {
		return strings.ToUpper(original[:1]) + original[1:]
	}

	return original
}

// DetectsSkip checks if the user wants to skip clarification
func (p *ResponseProcessorImpl) DetectsSkip(response string) bool {
	response = strings.TrimSpace(response)

	// Empty response means skip
	if response == "" {
		return true
	}

	// Explicit skip keywords
	skipWords := []string{"skip", "cancel", "stop", "exit", "quit"}
	responseLower := strings.ToLower(response)

	for _, skip := range skipWords {
		if responseLower == skip {
			return true
		}
	}

	return false
}
