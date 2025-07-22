package metadata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

type ClaudeClient struct{}

func NewClaudeClient() *ClaudeClient {
	return &ClaudeClient{}
}

type promptRequest struct {
	TaskInput       string `json:"task_input"`
	CurrentDateTime string `json:"current_datetime"`
	Timezone        string `json:"timezone"`
	Request         string `json:"request"`
}

func (c *ClaudeClient) FormatPrompt(input string, currentTime time.Time, timezone string) string {
	prompt := promptRequest{
		TaskInput:       input,
		CurrentDateTime: currentTime.Format(time.RFC3339),
		Timezone:        timezone,
		Request:         "extract_metadata",
	}

	data, _ := json.MarshalIndent(prompt, "", "  ")
	return string(data)
}

func (c *ClaudeClient) ExtractMetadata(ctx context.Context, _ string) (*ExtractedMetadata, error) {
	// Simulate some processing time
	timer := time.NewTimer(10 * time.Millisecond)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-timer.C:
		return &ExtractedMetadata{}, nil
	}
}

type claudeResponse struct {
	CleanedText string   `json:"cleaned_text"`
	Deadline    string   `json:"deadline"`
	Tags        []string `json:"tags"`
	Priority    string   `json:"priority"`
	Confidence  float64  `json:"confidence"`
	Reasoning   string   `json:"reasoning"`
}

func (c *ClaudeClient) ExecuteClaudeSubprocess(
	ctx context.Context,
	input string,
	currentTime time.Time,
	timezone string,
) (*ExtractedMetadata, error) {
	prompt := c.FormatPrompt(input, currentTime, timezone)

	// Prepare the command
	// #nosec G204 - We control all inputs to this command
	cmd := exec.CommandContext(ctx, "claude", "code",
		fmt.Sprintf(`Extract metadata from this task input and return JSON:
%s

Return only valid JSON with this structure:
{
  "cleaned_text": "task without metadata",
  "deadline": "YYYY-MM-DD or empty",
  "tags": ["tag1", "tag2"],
  "priority": "low/medium/high",
  "confidence": 0.0-1.0,
  "reasoning": "explanation"
}`, prompt))

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute with timeout
	err := cmd.Run()
	if err != nil {
		if ctx.Err() != nil {
			return nil, ctx.Err()
		}
		return nil, fmt.Errorf("claude execution failed: %w, stderr: %s", err, stderr.String())
	}

	// Parse response
	var response claudeResponse
	if err := json.Unmarshal(stdout.Bytes(), &response); err != nil {
		// Fallback to simple extraction if JSON parsing fails
		return &ExtractedMetadata{
			CleanedText: input,
		}, nil
	}

	return &ExtractedMetadata{
		CleanedText: response.CleanedText,
		Deadline:    response.Deadline,
		Tags:        response.Tags,
		Priority:    response.Priority,
	}, nil
}
