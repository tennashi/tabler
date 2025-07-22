package metadata_test

import (
	"testing"

	"github.com/tennashi/tabler/internal/metadata"
)

func TestMetadataCache(t *testing.T) {
	t.Run("stores and retrieves results", func(t *testing.T) {
		// Arrange
		cache := metadata.NewCache()
		key := "test input"
		expected := &metadata.ExtractedMetadata{
			CleanedText: "test",
			Tags:        []string{"work"},
			Priority:    "high",
		}

		// Act
		cache.Set(key, expected)
		result, found := cache.Get(key)

		// Assert
		if !found {
			t.Fatal("expected to find cached value")
		}
		if result.CleanedText != expected.CleanedText {
			t.Errorf("expected cleaned text %q, got %q", expected.CleanedText, result.CleanedText)
		}
		if len(result.Tags) != len(expected.Tags) || result.Tags[0] != expected.Tags[0] {
			t.Errorf("expected tags %v, got %v", expected.Tags, result.Tags)
		}
		if result.Priority != expected.Priority {
			t.Errorf("expected priority %q, got %q", expected.Priority, result.Priority)
		}
	})
}
