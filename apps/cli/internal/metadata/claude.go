package metadata

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
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

func (c *ClaudeClient) ExtractMetadata(ctx context.Context, input string) (*ExtractedMetadata, error) {
	// Use current time and default timezone
	currentTime := time.Now()
	timezone := "UTC"
	if tz := time.Local.String(); tz != "" {
		timezone = tz
	}

	return c.ExecuteClaudeSubprocess(ctx, input, currentTime, timezone)
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

	// Prepare the prompt for Claude
	claudePrompt := fmt.Sprintf(`You are a task metadata extractor. Given this task input:
%s

Extract and return ONLY valid JSON (no markdown, no explanation) with this exact structure:
{
  "cleaned_text": "task title without metadata",
  "deadline": "YYYY-MM-DD or empty string",
  "tags": ["tag1", "tag2"],
  "priority": "low/medium/high",
  "confidence": 0.0-1.0,
  "reasoning": "brief explanation"
}

Rules:
- cleaned_text: Remove all metadata (tags, dates, priority markers) from the task title
- deadline: Extract dates and convert to YYYY-MM-DD format
- tags: Extract meaningful categories/labels from the task content
- priority: Determine from urgency keywords (urgent/ASAP = high, important = medium, default = low)
- For Japanese input, extract metadata but keep cleaned_text in original language`, prompt)

	// #nosec G204 - We control all inputs to this command
	cmd := exec.CommandContext(ctx, "claude", "-p", claudePrompt)

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
	output := stdout.String()

	if err := json.Unmarshal([]byte(output), &response); err != nil {
		// Try to extract JSON from the output (Claude might include markdown)
		// Look for JSON between ```json and ```
		jsonStart := strings.Index(output, "```json")
		jsonEnd := strings.LastIndex(output, "```")
		if jsonStart != -1 && jsonEnd != -1 && jsonStart < jsonEnd {
			jsonStart += 7 // Skip "```json"
			jsonContent := strings.TrimSpace(output[jsonStart:jsonEnd])
			if err := json.Unmarshal([]byte(jsonContent), &response); err == nil {
				// Successfully parsed from markdown block
				goto parsed
			}
		}

		// Fallback to simple extraction if JSON parsing fails
		return &ExtractedMetadata{
			CleanedText: input,
		}, nil
	}

parsed:

	return &ExtractedMetadata{
		CleanedText: response.CleanedText,
		Deadline:    response.Deadline,
		Tags:        response.Tags,
		Priority:    response.Priority,
	}, nil
}
