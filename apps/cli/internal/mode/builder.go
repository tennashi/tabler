package mode

import (
	"github.com/tennashi/tabler/internal/clarification"
	"github.com/tennashi/tabler/internal/claude"
	"github.com/tennashi/tabler/internal/decomposition"
	"github.com/tennashi/tabler/internal/storage"
)

// ManagerBuilder helps build a ModeManager with optional features
type ManagerBuilder struct {
	useClarification bool
	useDecomposition bool
	claudeClient     *claude.Client
	storage          *storage.Storage
}

// NewManagerBuilder creates a new builder
func NewManagerBuilder() *ManagerBuilder {
	return &ManagerBuilder{}
}

// WithClarification enables dialogue-based clarification for Talk mode
func (b *ManagerBuilder) WithClarification() *ManagerBuilder {
	b.useClarification = true
	if b.claudeClient == nil {
		b.claudeClient = claude.NewClient()
	}
	return b
}

// WithDecomposition enables task decomposition for Planning mode
func (b *ManagerBuilder) WithDecomposition(storage *storage.Storage) *ManagerBuilder {
	b.useDecomposition = true
	b.storage = storage
	if b.claudeClient == nil {
		b.claudeClient = claude.NewClient()
	}
	return b
}

// Build creates the ModeManager with configured features
func (b *ManagerBuilder) Build() *ModeManager {
	manager := &ModeManager{
		handlers: make(map[Mode]ModeHandler),
		detector: NewModeDetector(),
	}

	// Always register Quick handler
	manager.RegisterHandler(QuickMode, NewQuickHandler())

	// Register Talk handler with or without clarification
	if b.useClarification && b.claudeClient != nil {
		// Create clarification components
		vaguenessDetector := clarification.NewVaguenessDetector()
		questionGen := clarification.NewQuestionGenerator(b.claudeClient)
		responseProcessor := clarification.NewResponseProcessor()
		dialogueManager := clarification.NewDialogueManager(vaguenessDetector, questionGen, responseProcessor)

		// Use enhanced Talk handler
		manager.RegisterHandler(TalkMode, NewTalkHandlerWithClarification(dialogueManager))
	} else {
		// Use basic Talk handler
		manager.RegisterHandler(TalkMode, NewTalkHandler())
	}

	// Register Planning handler with or without decomposition
	if b.useDecomposition && b.storage != nil && b.claudeClient != nil {
		// Create decomposition components
		complexityDetector := decomposition.NewComplexityDetector()
		decomposer := decomposition.NewTaskDecomposer(b.claudeClient)
		presenter := decomposition.NewInteractivePresenter()

		// Create adapters
		storageAdapter := NewStorageAdapter(b.storage)
		decomposerAdapter := NewDecomposerAdapter(decomposer)
		presenterAdapter := NewPresenterAdapter(presenter)

		// Use enhanced Planning handler
		manager.RegisterHandler(PlanningMode, NewPlanningHandlerWithDecomposition(
			storageAdapter,
			complexityDetector,
			decomposerAdapter,
			presenterAdapter,
		))
	} else {
		// Use basic Planning handler
		manager.RegisterHandler(PlanningMode, NewPlanningHandler())
	}

	return manager
}
