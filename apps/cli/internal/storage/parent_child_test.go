package storage

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/tennashi/tabler/internal/task"
)

func TestStorageParentChild(t *testing.T) {
	t.Run("CreateWithParent", func(t *testing.T) {
		t.Run("should create task with parent relationship", func(t *testing.T) {
			// Arrange
			s := setupTestStorage(t)
			
			// Create parent task first
			parent := createTestTask("Parent task")
			err := s.CreateTask(parent, nil)
			if err != nil {
				t.Fatalf("failed to create parent: %v", err)
			}
			
			// Create child task
			child := createTestTask("Child task")
			
			// Act
			err = s.CreateWithParent(child, parent.ID)
			
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			
			// Verify child was created
			retrieved, _, err := s.GetTask(child.ID)
			if err != nil {
				t.Fatalf("failed to get child task: %v", err)
			}
			if retrieved.Title != child.Title {
				t.Errorf("expected title %q, got %q", child.Title, retrieved.Title)
			}
		})
		
		t.Run("should fail with non-existent parent", func(t *testing.T) {
			// Arrange
			s := setupTestStorage(t)
			child := createTestTask("Child task")
			
			// Act
			err := s.CreateWithParent(child, "non-existent-id")
			
			// Assert
			if err == nil {
				t.Error("expected error for non-existent parent")
			}
		})
	})
	
	t.Run("GetChildren", func(t *testing.T) {
		t.Run("should retrieve all children of a parent task", func(t *testing.T) {
			// Arrange
			s := setupTestStorage(t)
			
			// Create parent
			parent := createTestTask("Parent task")
			if err := s.CreateTask(parent, nil); err != nil {
				t.Fatalf("failed to create parent: %v", err)
			}
			
			// Create children
			child1 := createTestTask("Child 1")
			child2 := createTestTask("Child 2")
			
			if err := s.CreateWithParent(child1, parent.ID); err != nil {
				t.Fatalf("failed to create child1: %v", err)
			}
			if err := s.CreateWithParent(child2, parent.ID); err != nil {
				t.Fatalf("failed to create child2: %v", err)
			}
			
			// Create unrelated task
			unrelated := createTestTask("Unrelated task")
			if err := s.CreateTask(unrelated, nil); err != nil {
				t.Fatalf("failed to create unrelated: %v", err)
			}
			
			// Act
			children, err := s.GetChildren(parent.ID)
			
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(children) != 2 {
				t.Errorf("expected 2 children, got %d", len(children))
			}
			
			// Verify children content
			childTitles := make(map[string]bool)
			for _, child := range children {
				childTitles[child.Title] = true
			}
			
			if !childTitles["Child 1"] {
				t.Error("expected to find 'Child 1'")
			}
			if !childTitles["Child 2"] {
				t.Error("expected to find 'Child 2'")
			}
		})
		
		t.Run("should return empty list for task with no children", func(t *testing.T) {
			// Arrange
			s := setupTestStorage(t)
			
			parent := createTestTask("Parent without children")
			if err := s.CreateTask(parent, nil); err != nil {
				t.Fatalf("failed to create parent: %v", err)
			}
			
			// Act
			children, err := s.GetChildren(parent.ID)
			
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if len(children) != 0 {
				t.Errorf("expected 0 children, got %d", len(children))
			}
		})
	})
	
	t.Run("GetParent", func(t *testing.T) {
		t.Run("should retrieve parent of a child task", func(t *testing.T) {
			// Arrange
			s := setupTestStorage(t)
			
			parent := createTestTask("Parent task")
			if err := s.CreateTask(parent, nil); err != nil {
				t.Fatalf("failed to create parent: %v", err)
			}
			
			child := createTestTask("Child task")
			if err := s.CreateWithParent(child, parent.ID); err != nil {
				t.Fatalf("failed to create child: %v", err)
			}
			
			// Act
			retrievedParent, err := s.GetParent(child.ID)
			
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if retrievedParent == nil {
				t.Fatal("expected parent to be non-nil")
			}
			if retrievedParent.ID != parent.ID {
				t.Errorf("expected parent ID %q, got %q", parent.ID, retrievedParent.ID)
			}
		})
		
		t.Run("should return nil for task with no parent", func(t *testing.T) {
			// Arrange
			s := setupTestStorage(t)
			
			task := createTestTask("Task without parent")
			if err := s.CreateTask(task, nil); err != nil {
				t.Fatalf("failed to create task: %v", err)
			}
			
			// Act
			parent, err := s.GetParent(task.ID)
			
			// Assert
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if parent != nil {
				t.Error("expected parent to be nil")
			}
		})
	})
}

// Helper function to create test task
func createTestTask(title string) *task.Task {
	return &task.Task{
		ID:        generateTestID(),
		Title:     title,
		Priority:  0,
		Completed: false,
	}
}

// Simple ID generator for tests
var testIDCounter int

func generateTestID() string {
	testIDCounter++
	return fmt.Sprintf("test-id-%d", testIDCounter)
}

// setupTestStorage creates a test storage instance
func setupTestStorage(t *testing.T) *Storage {
	t.Helper()
	tmpDir := t.TempDir()
	dbPath := filepath.Join(tmpDir, "test.db")
	
	storage, err := New(dbPath)
	if err != nil {
		t.Fatalf("failed to create storage: %v", err)
	}
	
	t.Cleanup(func() {
		if err := storage.Close(); err != nil {
			t.Errorf("failed to close storage: %v", err)
		}
	})
	
	if err := storage.Init(); err != nil {
		t.Fatalf("failed to init storage: %v", err)
	}
	
	return storage
}