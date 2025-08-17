# Smart Task Creation - Design Documents

## Overview

This directory contains technical design documents for the Smart Task Creation feature set, implementing the
vision from the [Smart Task Creation PRD](../../prd/smart_task_creation.md).

## Related Documents

- **PRD**: [Smart Task Creation PRD](../../prd/smart_task_creation.md)
- **ADRs**:
  - [ADR-001: LLM Provider Selection](../../adr/001-llm-provider-selection.md) - Use Claude Code subprocess
  - [ADR-002: Offline Storage Strategy](../../adr/002-offline-storage-strategy.md) - SQLite for local storage
  - [ADR-003: Logging Strategy](../../adr/003-logging-strategy.md) - Use case based logging

## Implementation Status

### Completed Features ✅

1. **[Minimal Task Management](./minimal_task_management.md)**
   - Basic task creation, listing, and completion
   - Foundation for task management system

2. **[Task Shortcuts](./task_shortcuts.md)**
   - Inline shortcuts (@, #, !) for quick metadata entry
   - Natural language patterns for dates and priorities

3. **[Task Management Commands](./task_management_commands.md)**
   - Full CRUD operations for tasks
   - Show, update, and delete functionality

### Designed Features 📋

1. **[LLM-based Metadata Extraction](./llm_metadata_extraction.md)**
   - Natural language date parsing
   - Category and priority inference
   - Single Claude call for all metadata

2. **[Smart Task Decomposition](./smart_decomposition.md)**
   - Detect complex tasks
   - Suggest logical subtasks
   - Parent-child relationships

3. **[Interactive Task Clarification](./interactive_clarification.md)**
   - AI-powered dialogue for vague inputs
   - Context-aware questioning
   - Brief, focused conversations

4. **[Input Mode System](./input_modes.md)**
   - Three modes: Quick, Talk, Planning
   - Manual control and auto-detection
   - Optimized for different workflows

5. **[Context Learning Foundation](./context_learning.md)**
   - Local-only pattern tracking
   - Personalized suggestions
   - Complete privacy control

## Architecture Overview

````text
┌─────────────────────────────────────────────────────┐
│                    CLI Interface                     │
│  (Quick Mode)    (Talk Mode)    (Planning Mode)    │
└─────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────┐
│                  Core Services                       │
│                                                      │
│  ┌─────────────┐  ┌──────────────┐  ┌────────────┐│
│  │  Metadata    │  │   Dialogue   │  │    Task    ││
│  │  Extractor   │  │   Manager    │  │ Decomposer ││
│  └─────────────┘  └──────────────┘  └────────────┘│
│                                                      │
│  ┌─────────────┐  ┌──────────────┐  ┌────────────┐│
│  │   Pattern    │  │    Mode      │  │   Claude   ││
│  │   Tracker    │  │   Detector   │  │   Client   ││
│  └─────────────┘  └──────────────┘  └────────────┘│
└─────────────────────────────────────────────────────┘
                           │
                           ▼
┌─────────────────────────────────────────────────────┐
│                 Storage Layer                        │
│                  (SQLite DB)                         │
│                                                      │
│  - Tasks Table (with parent_id)                     │
│  - User Patterns Table                              │
│  - Metadata Cache (in-memory)                       │
└─────────────────────────────────────────────────────┘
```text

## Cross-Feature Considerations

### Shared Components

- **Claude Client**: Used by all AI-powered features
- **SQLite Database**: Stores tasks and patterns
- **Mode Manager**: Coordinates between different input modes

### Data Flow

1. User input → Mode detection/selection
2. Mode handler → Appropriate feature services
3. Claude integration for AI features
4. Final task creation in storage

### Performance Requirements

- Task creation and persistence: <100ms
- Basic operations (list, complete): <200ms
- Shortcut parsing: <50ms for inline patterns
- Pattern matching and autocomplete: <50ms
- AI-enhanced operations: <2s (with 2s timeout)
- Mode detection: <50ms for automatic selection

### Privacy Principles

- All data stored locally
- No cloud sync by default
- User controls all learning
- Easy data export/deletion

## Implementation Guidelines

1. Start with simplest working implementation
2. Use Claude Code subprocess (ADR-001)
3. Follow logging strategy (ADR-003)
4. Test each feature thoroughly
5. Maintain backward compatibility

### Technical Implementation Details

#### User Interface
- Primary input: Single-line text field with Enter key submission
- Automatic field clearing after successful save
- Visual mode indicators near input field
- Full keyboard navigation support

#### Backend Architecture
- SQLite database with parent_id for task relationships
- Index on created_at, category, and status fields
- LLM integration via Claude API subprocess
- Graceful degradation when AI unavailable

#### Security & Privacy
- Input validation (max 500 characters)
- Parameterized queries for SQL injection prevention
- Local-only data storage by default
- No telemetry without explicit consent

## Next Steps

1. Implement LLM-based Metadata Extraction
2. Add Smart Task Decomposition
3. Build Interactive Clarification
4. Create Input Mode System
5. Add Context Learning
6. Gather user feedback
7. Iterate based on usage
````
