package claude

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os/exec"
	"strings"
)

// Client provides general Claude AI execution capabilities
type Client struct{}

// NewClient creates a new Claude client
func NewClient() *Client {
	return &Client{}
}

// Execute runs a prompt through Claude and returns the response
func (c *Client) Execute(ctx context.Context, prompt string) (string, error) {
	// Create the claude command
	cmd := exec.CommandContext(ctx, "claude", prompt)

	// Capture output
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	if err := cmd.Run(); err != nil {
		// Check if the error is because claude command doesn't exist
		if strings.Contains(stderr.String(), "command not found") ||
			strings.Contains(err.Error(), "executable file not found") {
			return "", fmt.Errorf("claude CLI not found: %w", err)
		}
		return "", fmt.Errorf("claude execution failed: %s", stderr.String())
	}

	// Return the response
	response := strings.TrimSpace(stdout.String())
	return response, nil
}

// ExecuteJSON runs a prompt and expects a JSON response
func (c *Client) ExecuteJSON(ctx context.Context, prompt string, result interface{}) error {
	response, err := c.Execute(ctx, prompt)
	if err != nil {
		return err
	}

	// Try to parse as JSON
	if err := json.Unmarshal([]byte(response), result); err != nil {
		// If not valid JSON, return the raw response error
		return fmt.Errorf("invalid JSON response: %s", response)
	}

	return nil
}
