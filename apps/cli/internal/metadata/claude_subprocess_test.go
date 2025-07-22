package metadata_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/tennashi/tabler/internal/metadata"
)

func TestClaudeSubprocess(t *testing.T) {
	t.Run("executes claude subprocess", func(t *testing.T) {
		// Skip if claude is not available
		if _, err := os.Stat("/usr/local/bin/claude"); os.IsNotExist(err) {
			t.Skip("claude binary not found")
		}

		// Arrange
		client := metadata.NewClaudeClient()
		ctx := context.Background()
		currentTime := time.Date(2024, 1, 15, 10, 0, 0, 0, time.UTC)

		// Act
		result, err := client.ExecuteClaudeSubprocess(ctx, "finish report by tomorrow", currentTime, "Asia/Tokyo")
		// Assert
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if result == nil {
			t.Fatal("expected non-nil result")
		}
		// We can't assert exact values since Claude's output varies,
		// but we should have some cleaned text
		if result.CleanedText == "" {
			t.Error("expected non-empty cleaned text")
		}
	})
}
