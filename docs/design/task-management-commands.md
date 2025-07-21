# Design Doc: Task Management Commands

<!-- 
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

This feature completes the basic CRUD operations by adding commands to show task details, update tasks, and delete tasks. These commands provide full task lifecycle management beyond just creation and listing.

## Background

[Link to PRD: ../prd/smart_task_creation.md](../prd/smart_task_creation.md)

After implementing task creation with shortcuts, users need ways to view full details, modify existing tasks, and remove completed or unwanted tasks. This completes the foundation for task management.

## Goals

- Show complete task information including all metadata
- Update tasks with full shortcut parsing support
- Delete tasks with safety confirmation
- Maintain data consistency across operations
- Provide clear feedback for all operations

## Non-Goals

- Bulk operations (update/delete multiple)
- Undo/redo functionality
- Task archiving or soft delete
- Partial field updates
- Command aliases or shortcuts
- Interactive editing mode

## Design

### High-Level Architecture

```
┌─────────────────┐     ┌──────────────┐     ┌─────────────┐
│ Command Input   │────▶│  Operation   │────▶│ Persistence │
│ (show/up/del)   │     │  Handler     │     │    Layer    │
└─────────────────┘     └──────────────┘     └─────────────┘
                               │
                               ▼
                        ┌──────────────┐
                        │   Parser     │
                        │ (for update) │
                        └──────────────┘
```

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Show Operation

**Purpose**: Display complete task information

**Responsibilities**:
- Retrieve task by identifier
- Gather all associated metadata
- Format for human readability
- Handle non-existent tasks

**Interface**:
- Input: Task identifier
- Output: Formatted task details or not found

#### Component 2: Update Operation

**Purpose**: Modify existing task content

**Responsibilities**:
- Validate task exists
- Parse new content with shortcuts
- Replace task data completely
- Preserve system fields (ID, creation time)
- Track modification time

**Interface**:
- Input: Task identifier and new content
- Output: Success confirmation or error

#### Component 3: Delete Operation

**Purpose**: Remove tasks permanently

**Responsibilities**:
- Verify task exists
- Request user confirmation
- Remove task and relationships
- Ensure no orphaned data
- Report successful deletion

**Interface**:
- Input: Task identifier
- Output: Confirmation prompt, then result

#### Component 4: Confirmation Handler

**Purpose**: Protect against accidental deletions

**Responsibilities**:
- Display task summary before deletion
- Request explicit confirmation
- Default to safe choice (no)
- Handle user response

**Interface**:
- Input: Task to be deleted
- Output: User's confirmation decision

### Data Model

<!-- Show logical data model, not physical implementation -->

**Task Tracking Updates**:
- Add modification timestamp
- Preserve creation timestamp
- Track last update time

**Data Integrity**:
- Cascading deletion of related data
- No partial updates allowed
- Atomic operations required

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**Command Patterns**:

1. **Show Details**
   - Command: `show <id>`
   - Output: All task fields formatted
   - Error: Task not found message

2. **Update Task**
   - Command: `update <id> "new content"`
   - Behavior: Full replacement with shortcut parsing
   - Output: Confirmation message

3. **Delete Task**
   - Command: `delete <id>`
   - Behavior: Confirm then remove
   - Output: Deletion confirmation

**Show Format Example**:
```
ID: abc123
Task: Fix login bug
Status: Pending
Tags: work, urgent
Priority: High
Deadline: Tomorrow (Jan 16, 2024)
Created: Jan 15, 2024 10:30 AM
Modified: Jan 15, 2024 2:45 PM
```

**Update Behavior**:
- Complete content replacement
- All shortcuts re-parsed
- Creation time preserved
- Modification time updated

**Delete Flow**:
```
Delete task "Fix login bug"? (y/N): _
```

### Error Handling

- Unknown identifier → Clear not found message
- Empty update content → Reject with reason
- Cancelled deletion → No changes, confirm cancellation
- System errors → Rollback, user-friendly message

### Logging Strategy

**Applicable Use Cases**:
- [x] User Behavior - Command usage patterns
- [x] Error Tracking - Failed operations
- [ ] Tracing - Not needed
- [ ] Performance - Simple operations
- [x] Security Audit - Deletion tracking
- [ ] Business Metrics - Covered by user behavior

**Implementation Details**:
```
User Behavior:
- Events: command_executed
- Fields: command_type, success_status
- Retention: 90 days

Error Tracking:
- Events: operation_failed
- Fields: operation, error_type
- Retention: 30 days

Security Audit:
- Events: task_deleted
- Fields: timestamp, confirmed
- Retention: 180 days
```

**Privacy Considerations**:
- No task content in logs
- Only operation metadata tracked

## Security Considerations

- Validate all identifiers
- Require confirmation for destructive operations
- Prevent injection attacks through input validation
- Limit error detail exposure

## Testing Strategy

- Operation flows: Each command end-to-end
- Error cases: Invalid IDs, empty input
- Confirmation: Both yes and no paths
- Data integrity: No orphaned data after delete
- Update preservation: Correct fields retained

## Migration Plan

1. Add modification timestamp field
2. Default to creation time for existing tasks
3. No other structural changes needed

## Alternatives Considered

### Alternative 1: Patch Operations

Allow updating specific fields only.

**Why not chosen**: Complex interface, inconsistent with shortcut parsing approach, minimal benefit for command-line usage.

### Alternative 2: Versioned Tasks

Keep history of all changes.

**Why not chosen**: Adds storage complexity, requires version management commands, exceeds requirements for simple task manager.

### Alternative 3: Trash/Archive System

Move deleted tasks to recoverable storage.

**Why not chosen**: Adds state complexity, requires additional commands, users expect permanent deletion in simple tools.