# Design Doc: Phase 1 - Basic Task Creation (CLI)

## Overview

Implement a CLI-focused task creation tool with inline shortcut parsing and local SQLite storage. Simple, fast,
and works offline.

## Background

[PRD: Smart Task Creation](../../prd/smart_task_creation.md) - Phase 1 Implementation

Focus on CLI-only implementation for rapid development and personal use.

## Goals

- Single command to create tasks: `tabler add "task description"`
- Inline shortcuts (@, #, !) parsing
- SQLite local storage
- Basic task listing and filtering
- Fast and simple

## Non-Goals

- Web UI or TUI
- AI features
- Multi-device sync
- Complex task management

## Design

### High-Level Architecture

```text
┌─────────────────┐
│   CLI Command   │ ← tabler add/list/show/delete
└────────┬────────┘
         │
┌────────▼────────┐
│ Shortcut Parser │ ← Extract @, #, !
└────────┬────────┘
         │
┌────────▼────────┐
│  Task Service   │ ← Business logic
└────────┬────────┘
         │
┌────────▼────────┐
│ SQLite Storage  │ ← ~/.tabler/tasks.db
└─────────────────┘
```

### CLI Commands

```bash
# Create task with shortcuts
tabler add "Fix bug in login #work @tomorrow !!"

# List tasks
tabler list
tabler list --tag work
tabler list --today
tabler list --overdue

# Show task details
tabler show <id>

# Update task
tabler update <id> "new description"

# Delete task
tabler delete <id>

# Mark complete
tabler done <id>
```

### Detailed Design

#### Component 1: CLI Interface

**Implementation**: Using cobra or similar CLI framework

```go
type CLI struct {
    taskService TaskService
}

func (c *CLI) AddCommand() *cobra.Command {
    return &cobra.Command{
        Use:   "add [task description]",
        Short: "Add a new task",
        Args:  cobra.ExactArgs(1),
        Run:   c.addTask,
    }
}
```

#### Component 2: Shortcut Parser

**Patterns**:

- `@today`, `@tomorrow`, `@mon`, `@2024-01-15` → deadline
- `#tag` → tags (multiple allowed)
- `!`, `!!`, `!!!` → priority (1-3)

**Examples**:

```text
Input: "Fix login bug #work #urgent @tomorrow !!"
Output: 
  Title: "Fix login bug"
  Tags: ["work", "urgent"]
  Deadline: 2024-01-16
  Priority: 2
```

#### Component 3: Storage

**Database Location**: `~/.tabler/tasks.db`

**Schema**:

```sql
CREATE TABLE tasks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    deadline INTEGER,
    priority INTEGER DEFAULT 0,
    completed INTEGER DEFAULT 0,
    created_at INTEGER NOT NULL,
    updated_at INTEGER NOT NULL
);

CREATE TABLE task_tags (
    task_id TEXT NOT NULL,
    tag TEXT NOT NULL,
    PRIMARY KEY (task_id, tag),
    FOREIGN KEY (task_id) REFERENCES tasks(id)
);
```

### Output Format

```bash
$ tabler list
ID    TASK                           TAGS         DUE        PRI
---   ----------------------------   ----------   --------   ---
a3f   Fix login bug                  work,urgent  tomorrow   !!
b2d   Review PR #123                 work         today      !
c1e   Buy groceries                  personal     -          -

$ tabler show a3f
Task: Fix login bug
Tags: #work #urgent
Due: Tomorrow (2024-01-16)
Priority: High (!!)
Created: 2024-01-15 10:30:00
```

### Configuration

Config file at `~/.tabler/config.toml`:

```toml
# Date formats
date_format = "2006-01-02"
time_format = "15:04"

# Display
color = true
timezone = "local"
```

## Testing Strategy

- **Unit tests**: Parser patterns, date parsing
- **Integration tests**: CLI commands end-to-end
- **Manual tests**: Various shortcut combinations

## Implementation Notes

- Task IDs should be sortable and time-ordered (e.g., ULID)
- Database foreign key constraints must be enabled
- Date parsing should respect user's timezone
- Use database transactions for data consistency
- Table output should align columns properly
- Color support should be optional/configurable

## Future Considerations

- Phase 2: Add Claude Code integration
- Phase 3: Interactive mode
- Export/import functionality
- Task templates
