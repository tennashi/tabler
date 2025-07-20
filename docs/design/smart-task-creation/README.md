# Design: Smart Task Creation

## Overview

A fast, intelligent task capture system that supports natural language input with inline shortcuts, AI-powered understanding, and task clarification through dialogue. The implementation uses Claude Code for AI features and SQLite for offline-first storage.

## Related Documents

- **PRD**: [Smart Task Creation PRD](../../prd/smart_task_creation.md)
- **ADRs**: 
  - [ADR-001: LLM Provider Selection](../../adr/001-llm-provider-selection.md)
  - [ADR-002: Offline Storage Strategy](../../adr/002-offline-storage-strategy.md)

## Implementation Phases

### Phase 1: Basic Task Creation
**Status**: Not Started

**Scope**:
- Single-line task input with Enter to save
- Inline shortcuts (@, #, !) parsing
- Basic task CRUD operations
- Simple task list display
- SQLite local storage

**Design Doc**: [phase1-basic-task-creation.md](./phase1-basic-task-creation.md)

### Phase 2: AI Enhancement
**Status**: Not Started

**Dependencies**: Phase 1 must be complete

**Scope**:
- Natural language processing for dates and categories
- Smart task decomposition suggestions
- Claude Code integration for understanding intent
- Autocomplete for frequently used tags

**Design Doc**: [phase2-ai-enhancement.md](./phase2-ai-enhancement.md)

### Phase 3: Interactive Features
**Status**: Not Started

**Dependencies**: Phase 2 must be complete

**Scope**:
- AI-powered task clarification dialogue
- Three input modes (Quick, Talk, Planning)
- Mode switching system
- Context learning from usage patterns

**Design Doc**: [phase3-interactive-features.md](./phase3-interactive-features.md)

## Architecture Overview

```
┌─────────────────────┐
│   UI (TUI/Web)      │
├─────────────────────┤
│  Task Input Parser  │ ← Handles shortcuts (@, #, !)
├─────────────────────┤
│    Task Service     │ ← Business logic
├─────────────────────┤
│ ┌─────────┐┌──────┐ │
│ │ Storage ││  AI  │ │ ← SQLite + Claude Code
│ │ (SQLite)││Engine│ │
│ └─────────┘└──────┘ │
└─────────────────────┘
```

## Cross-Phase Considerations

### Shared Components
- **Task Model**: Core data structure evolves across phases
- **Parser**: Inline shortcut parser enhanced with AI in Phase 2
- **Storage Layer**: SQLite schema migrations between phases

### Data Model Evolution
- Phase 1: Basic task fields (title, created_at, shortcuts)
- Phase 2: Add AI metadata (suggested_category, decomposition)
- Phase 3: Add conversation history and mode preferences

### Interface Design
- Parser interface must support both simple regex (Phase 1) and AI (Phase 2)
- Storage interface abstracts SQLite implementation
- AI interface allows swapping Claude Code with other providers later

## Success Metrics

Technical success measured by:
- **Performance**: Task creation <200ms, AI suggestions <2s
- **Reliability**: 99.9% success rate for basic operations
- **Offline**: 100% functionality for basic task creation without network
- **User Experience**: <5 seconds average task creation time