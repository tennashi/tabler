package mode

import (
	"context"
	"strings"
	"testing"

	"github.com/tennashi/tabler/internal/decomposition"
	"github.com/tennashi/tabler/internal/storage"
	"github.com/tennashi/tabler/internal/task"
)

func TestPlanningModeIntegration(t *testing.T) {
	t.Run("should integrate planning mode with decomposition", func(t *testing.T) {
		// Arrange
		// Create real storage with test database
		tmpDir := t.TempDir()
		dbPath := tmpDir + "/test.db"
		
		store, err := storage.New(dbPath)
		if err != nil {
			t.Fatalf("failed to create storage: %v", err)
		}
		defer store.Close()
		
		if err := store.Init(); err != nil {
			t.Fatalf("failed to init storage: %v", err)
		}
		
		// Create mock Claude client for now (real integration would need actual Claude)
		claude := &mockClaudeForIntegration{}
		
		// Create real decomposition components
		detector := decomposition.NewComplexityDetector()
		decomposer := decomposition.NewTaskDecomposer(claude)
		presenter := decomposition.NewInteractivePresenter()
		
		// Create handler with real components
		handler := NewPlanningHandlerWithDecomposition(
			&storageAdapter{store},
			detector,
			decomposer,
			presenter,
		)
		
		// Set input to simulate user selecting "none" for subtasks
		handler.SetInput(strings.NewReader("none\n"))
		
		// Act - test with simple task (should not decompose)
		simpleTask := "buy milk"
		result, err := handler.Process(context.Background(), simpleTask)
		
		// Assert
		if err != nil {
			t.Fatalf("unexpected error for simple task: %v", err)
		}
		if result.Title != simpleTask {
			t.Errorf("expected title %q, got %q", simpleTask, result.Title)
		}
		
		// Verify task was stored
		retrieved, _, err := store.GetTask(result.ID)
		if err != nil {
			t.Fatalf("failed to retrieve task: %v", err)
		}
		if retrieved.Title != simpleTask {
			t.Errorf("stored task has wrong title: %q", retrieved.Title)
		}
	})
}

// storageAdapter adapts storage.Storage to StorageWithDecomposition interface
type storageAdapter struct {
	*storage.Storage
}

func (s *storageAdapter) Create(t *task.Task) error {
	return s.CreateTask(t, nil)
}

func (s *storageAdapter) CreateWithParent(t *task.Task, parentID string) error {
	return s.Storage.CreateWithParent(t, parentID)
}

// mockClaudeForIntegration is a simple mock for integration testing
type mockClaudeForIntegration struct{}

func (m *mockClaudeForIntegration) Execute(_ context.Context, prompt string) (string, error) {
	// Simple mock response for integration test
	return `Here are the subtasks:

1. Book venue
2. Invite speakers  
3. Setup registration`, nil
}