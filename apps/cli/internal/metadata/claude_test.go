package metadata_test

import (
	"context"
	"testing"
	"time"

	"github.com/tennashi/tabler/internal/metadata"
)

func TestClaudeClient(t *testing.T) {
	t.Run("formats prompt correctly", func(t *testing.T) {
		// Arrange
		client := metadata.NewClaudeClient()
		input := "urgent: finish report by tomorrow #work"
		currentTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)

		// Act
		prompt := client.FormatPrompt(input, currentTime, "Asia/Tokyo")

		// Assert
		expected := `{
  "task_input": "urgent: finish report by tomorrow #work",
  "current_datetime": "2024-01-15T10:00:00Z",
  "timezone": "Asia/Tokyo",
  "request": "extract_metadata"
}`
		if prompt != expected {
			t.Errorf("expected prompt:\n%s\ngot:\n%s", expected, prompt)
		}
	})

	t.Run("handles subprocess timeout", func(t *testing.T) {
		// Arrange
		client := metadata.NewClaudeClient()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
		defer cancel()

		// Act
		_, err := client.ExtractMetadata(ctx, "test input")

		// Assert
		if err == nil {
			t.Fatal("expected timeout error")
		}
		if err.Error() != "context deadline exceeded" {
			t.Errorf("expected context deadline error, got: %v", err)
		}
	})
}
