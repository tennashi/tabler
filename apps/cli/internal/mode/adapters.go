package mode

import (
	"context"

	"github.com/tennashi/tabler/internal/decomposition"
	"github.com/tennashi/tabler/internal/storage"
	"github.com/tennashi/tabler/internal/task"
)

// StorageAdapter adapts storage.Storage to StorageWithDecomposition
type StorageAdapter struct {
	storage *storage.Storage
}

// NewStorageAdapter creates a new storage adapter
func NewStorageAdapter(storage *storage.Storage) *StorageAdapter {
	return &StorageAdapter{storage: storage}
}

// Create implements StorageWithDecomposition
func (s *StorageAdapter) Create(t *task.Task) error {
	return s.storage.CreateTask(t, []string{})
}

// CreateWithParent implements StorageWithDecomposition
func (s *StorageAdapter) CreateWithParent(t *task.Task, parentID string) error {
	// For now, just create the task without parent relationship
	// TODO: Implement parent-child relationship in storage when parentID is provided
	_ = parentID // Will be used when parent-child relationship is implemented
	return s.storage.CreateTask(t, []string{})
}

// DecomposerAdapter adapts decomposition.TaskDecomposer to Decomposer interface
type DecomposerAdapter struct {
	decomposer *decomposition.TaskDecomposer
}

// NewDecomposerAdapter creates a new decomposer adapter
func NewDecomposerAdapter(decomposer *decomposition.TaskDecomposer) *DecomposerAdapter {
	return &DecomposerAdapter{decomposer: decomposer}
}

// Decompose implements Decomposer
func (d *DecomposerAdapter) Decompose(ctx context.Context, task string) (*decomposition.DecompositionResult, error) {
	return d.decomposer.Decompose(ctx, task)
}

// PresenterAdapter adapts decomposition.InteractivePresenter to Presenter interface
type PresenterAdapter struct {
	presenter *decomposition.InteractivePresenter
}

// NewPresenterAdapter creates a new presenter adapter
func NewPresenterAdapter(presenter *decomposition.InteractivePresenter) *PresenterAdapter {
	return &PresenterAdapter{presenter: presenter}
}

// Present implements Presenter
func (p *PresenterAdapter) Present(result *decomposition.DecompositionResult) string {
	return p.presenter.Present(result)
}

// ParseSelection implements Presenter
func (p *PresenterAdapter) ParseSelection(input string, total int) ([]int, error) {
	return p.presenter.ParseSelection(input, total)
}
