package clarification

import (
	"context"
	"fmt"
	"strings"
)

// ClaudeClient interface for Claude integration
type ClaudeClient interface {
	Execute(ctx context.Context, prompt string) (string, error)
}

// QuestionGeneratorImpl creates contextual questions using Claude
type QuestionGeneratorImpl struct {
	claude ClaudeClient
}

// NewQuestionGenerator creates a new question generator
func NewQuestionGenerator(claude ClaudeClient) *QuestionGeneratorImpl {
	return &QuestionGeneratorImpl{
		claude: claude,
	}
}

// GenerateQuestion creates the next clarifying question or signals completion
func (g *QuestionGeneratorImpl) GenerateQuestion(ctx context.Context, session *DialogueSession) (string, bool, error) {
	// Build prompt for Claude
	prompt := g.buildPrompt(session)

	// Call Claude
	response, err := g.claude.Execute(ctx, prompt)
	if err != nil {
		return "", false, err
	}

	// Check if Claude indicates completion
	response = strings.TrimSpace(response)
	if strings.ToUpper(response) == "COMPLETE" {
		return "", true, nil
	}

	// Extract question from response
	question := g.extractQuestion(response)

	return question, false, nil
}

// buildPrompt creates the prompt for Claude
func (g *QuestionGeneratorImpl) buildPrompt(session *DialogueSession) string {
	var prompt strings.Builder

	prompt.WriteString("You are helping clarify a vague task. ")
	prompt.WriteString("Generate ONE clarifying question to gather missing information.\n\n")

	prompt.WriteString(fmt.Sprintf("Original task: \"%s\"\n\n", session.OriginalInput))

	// Include dialogue history
	if len(session.History) > 0 {
		prompt.WriteString("Dialogue so far:\n")
		for _, exchange := range session.History {
			prompt.WriteString(fmt.Sprintf("Q: %s\n", exchange.Question))
			if exchange.Answer != "" {
				prompt.WriteString(fmt.Sprintf("A: %s\n", exchange.Answer))
			}
		}
		prompt.WriteString("\n")
	}

	// Include extracted information
	if len(session.ExtractedInfo) > 0 {
		prompt.WriteString("Information gathered:\n")
		for key, value := range session.ExtractedInfo {
			prompt.WriteString(fmt.Sprintf("- %s: %s\n", key, value))
		}
		prompt.WriteString("\n")
	}

	prompt.WriteString("Instructions:\n")
	prompt.WriteString("- If you have enough information to create a clear task, respond with just: COMPLETE\n")
	prompt.WriteString("- Otherwise, ask ONE specific question to clarify what's missing\n")
	prompt.WriteString("- Focus on: what, when, who, or specific details\n")
	prompt.WriteString("- Keep questions short and natural\n")
	prompt.WriteString("- Do not include any explanation, just the question\n")

	return prompt.String()
}

// extractQuestion cleans up Claude's response to get just the question
func (g *QuestionGeneratorImpl) extractQuestion(response string) string {
	// Remove any leading/trailing whitespace
	question := strings.TrimSpace(response)

	// Ensure it ends with a question mark
	if !strings.HasSuffix(question, "?") {
		question += "?"
	}

	// Remove any potential prefix like "Question:" or "Q:"
	prefixes := []string{"Question:", "Q:", "Ask:"}
	for _, prefix := range prefixes {
		if strings.HasPrefix(question, prefix) {
			question = strings.TrimSpace(question[len(prefix):])
		}
	}

	return question
}
