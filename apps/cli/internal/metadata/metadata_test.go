package metadata_test

import (
	"context"
	"testing"

	"github.com/tennashi/tabler/internal/metadata"
)

func TestMetadataService(t *testing.T) {
	t.Run("extraction", func(t *testing.T) {
		t.Run("with empty input should return error", func(t *testing.T) {
			// Arrange
			service := metadata.NewService(nil)
			ctx := context.Background()

			// Act
			result, err := service.Extract(ctx, "")

			// Assert
			if err == nil {
				t.Fatal("expected error for empty input")
			}
			if result != nil {
				t.Error("expected nil result for empty input")
			}
		})

		t.Run("with simple text should return cleaned text", func(t *testing.T) {
			// Arrange
			service := metadata.NewService(nil)
			ctx := context.Background()
			input := "Buy groceries"

			// Act
			result, err := service.Extract(ctx, input)
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result == nil {
				t.Fatal("expected non-nil result")
			}
			if result.CleanedText != input {
				t.Errorf("expected cleaned text %q, got %q", input, result.CleanedText)
			}
		})

		t.Run("extracts deadline from 'by tomorrow'", func(t *testing.T) {
			// Arrange
			claude := &mockClaude{
				response: &metadata.ExtractedMetadata{
					CleanedText: "finish report",
					Deadline:    "2024-01-16",
					Tags:        []string{"report"},
					Priority:    "high",
				},
			}
			service := metadata.NewService(claude)
			ctx := context.Background()

			// Act
			result, err := service.Extract(ctx, "urgent: finish report by tomorrow #work")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Deadline != "2024-01-16" {
				t.Errorf("expected deadline 2024-01-16, got %q", result.Deadline)
			}
		})

		t.Run("extracts tags from task content", func(t *testing.T) {
			// Arrange
			claude := &mockClaude{
				response: &metadata.ExtractedMetadata{
					CleanedText: "prepare quarterly sales report",
					Tags:        []string{"sales", "report", "quarterly"},
				},
			}
			service := metadata.NewService(claude)
			ctx := context.Background()

			// Act
			result, err := service.Extract(ctx, "prepare quarterly sales report")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(result.Tags) != 3 {
				t.Fatalf("expected 3 tags, got %d", len(result.Tags))
			}
			expectedTags := []string{"sales", "report", "quarterly"}
			for i, tag := range expectedTags {
				if result.Tags[i] != tag {
					t.Errorf("expected tag[%d] to be %q, got %q", i, tag, result.Tags[i])
				}
			}
		})

		t.Run("extracts priority from keywords", func(t *testing.T) {
			// Arrange
			claude := &mockClaude{
				response: &metadata.ExtractedMetadata{
					CleanedText: "submit bug report",
					Priority:    "high",
				},
			}
			service := metadata.NewService(claude)
			ctx := context.Background()

			// Act
			result, err := service.Extract(ctx, "URGENT: submit bug report ASAP")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.Priority != "high" {
				t.Errorf("expected priority to be 'high', got %q", result.Priority)
			}
		})

		t.Run("handles Japanese input", func(t *testing.T) {
			// Arrange
			claude := &mockClaude{
				response: &metadata.ExtractedMetadata{
					CleanedText: "会議の議事録を作成",
					Deadline:    "2024-01-19",
					Tags:        []string{"会議", "議事録"},
					Priority:    "medium",
				},
			}
			service := metadata.NewService(claude)
			ctx := context.Background()

			// Act
			result, err := service.Extract(ctx, "来週の金曜日までに会議の議事録を作成")
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if result.CleanedText != "会議の議事録を作成" {
				t.Errorf("expected cleaned text %q, got %q", "会議の議事録を作成", result.CleanedText)
			}
			if result.Deadline != "2024-01-19" {
				t.Errorf("expected deadline 2024-01-19, got %q", result.Deadline)
			}
			if len(result.Tags) != 2 {
				t.Fatalf("expected 2 tags, got %d", len(result.Tags))
			}
			expectedTags := []string{"会議", "議事録"}
			for i, tag := range expectedTags {
				if result.Tags[i] != tag {
					t.Errorf("expected tag[%d] to be %q, got %q", i, tag, result.Tags[i])
				}
			}
		})
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
