package metadata

import (
	"context"
	"encoding/json"
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
