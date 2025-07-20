# ADR-002: Offline Storage Strategy

## Status

Proposed

## Context

The Smart Task Creation feature requires:
- Basic task creation must work offline (per PRD)
- Tasks created offline need to persist until online
- AI features (Claude Code) won't work offline, but basic creation should

We need a storage strategy that handles offline-first task creation with eventual sync.

## Decision

Use **Local SQLite database** for offline storage.

Implementation approach:
- SQLite for local task storage (works offline)
- Tasks created offline get queued for AI processing
- When online + Claude Code available, process queued tasks
- Simple file-based storage, no complex sync

## Consequences

### Positive
- True offline capability for basic task creation
- No data loss when offline
- SQLite is simple, reliable, well-understood
- Single file storage, easy backup
- Can query and filter tasks locally

### Negative
- Need SQLite driver/library
- Database migrations for schema changes
- No real-time sync across devices

### Neutral
- Good enough for personal task management
- Can add sync layer later if needed
- Standard SQL knowledge applies

## Options Considered

### Option 1: SQLite Database (Selected)
- **Pros**: Offline-first, reliable, queryable, single file
- **Cons**: Needs migration strategy, no built-in sync
- **Evaluation**: Best for offline-first requirement

### Option 2: JSON File Storage
- **Pros**: Simple, no dependencies
- **Cons**: No queries, concurrent access issues, can corrupt
- **Evaluation**: Too basic for task management needs

### Option 3: IndexedDB (Browser)
- **Pros**: Native browser support, good for web app
- **Cons**: Browser-only, complex API, size limits
- **Evaluation**: Limits deployment options

### Option 4: Cloud-First with Cache
- **Pros**: Real-time sync, backup included
- **Cons**: Doesn't meet offline requirement
- **Evaluation**: Conflicts with offline-first requirement