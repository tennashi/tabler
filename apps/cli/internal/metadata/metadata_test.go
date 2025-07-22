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
