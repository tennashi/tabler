package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/tennashi/tabler/internal/metadata"
	"github.com/tennashi/tabler/internal/service"
)

func TestTaskServiceWithLLM(t *testing.T) {
	t.Run("creates task with LLM metadata extraction", func(t *testing.T) {
		// Arrange
		dataDir := t.TempDir()
		claude := &mockClaude{
			response: &metadata.ExtractedMetadata{
				CleanedText: "finish quarterly report",
				Deadline:    "2024-01-16",
				Tags:        []string{"report", "quarterly"},
				Priority:    "high",
			},
		}
		metadataService := metadata.NewService(claude)

		taskService, err := service.NewTaskServiceWithMetadata(dataDir, metadataService)
		if err != nil {
			t.Fatalf("failed to create service: %v", err)
		}
		defer func() {
			_ = taskService.Close()
		}()

		// Act
		taskID, err := taskService.CreateTaskFromInput("urgent: finish quarterly report by tomorrow")
		// Assert
		if err != nil {
			t.Fatalf("failed to create task: %v", err)
		}

		// Verify task was created with extracted metadata
		task, tags, err := taskService.GetTask(taskID)
		if err != nil {
			t.Fatalf("failed to get task: %v", err)
		}

		if task.Title != "finish quarterly report" {
			t.Errorf("expected title %q, got %q", "finish quarterly report", task.Title)
		}

		expectedDeadline := time.Date(2024, 1, 16, 0, 0, 0, 0, time.UTC)
		if !task.Deadline.Equal(expectedDeadline) {
			t.Errorf("expected deadline %v, got %v", expectedDeadline, task.Deadline)
		}

		if len(tags) != 2 {
			t.Errorf("expected 2 tags, got %d", len(tags))
		}

		if task.Priority != 3 {
			t.Errorf("expected priority 3 (high), got %d", task.Priority)
		}
	})
}

type mockClaude struct {
	response *metadata.ExtractedMetadata
	err      error
}

func (m *mockClaude) ExtractMetadata(_ context.Context, _ string) (*metadata.ExtractedMetadata, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.response, nil
}
