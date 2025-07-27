# Design Doc: Minimal Task Management

<!--
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

A simple task manager that provides the absolute minimum functionality to be useful: create tasks, list them,
and mark them as complete. This establishes the foundation for all future task management features.

## Background

[Link to PRD: ../prd/smart_task_creation.md](../prd/smart_task_creation.md)

This is the foundational implementation that enables basic task tracking. While the PRD describes advanced
features, this design focuses on delivering immediate value with the smallest possible implementation.

## Goals

- Create tasks with a title
- List all tasks with their status
- Mark tasks as complete
- Provide persistent local storage
- Establish patterns for future features
- Complete implementation in under 2 days

## Non-Goals

- Tags or categories
- Priority levels
- Deadlines or due dates
- Update/delete operations
- Filtering or search
- Any form of shortcuts or metadata

## Design

### High-Level Architecture

````text
┌─────────────────┐
│ Command Line    │ ← User interaction
│   Interface     │
└────────┬────────┘
         │
┌────────▼────────┐
│ Business Logic  │ ← Core rules and operations
│     Layer       │
└────────┬────────┘
         │
┌────────▼────────┐
│  Persistence    │ ← Data storage
│     Layer       │
└─────────────────┘
```text

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Command Line Interface

**Purpose**: Provide user interaction through terminal commands

**Responsibilities**:

- Accept user commands and arguments
- Route to appropriate operations
- Display results in readable format
- Report errors clearly

**Interface**:

- Input: Command line arguments
- Output: Formatted text output

#### Component 2: Business Logic Layer

**Purpose**: Implement core task management rules

**Responsibilities**:

- Validate task data (non-empty titles)
- Assign unique identifiers to tasks
- Track task completion state
- Coordinate between interface and storage

**Interface**:

- Input: Task operations (create, list, complete)
- Output: Operation results or errors

#### Component 3: Persistence Layer

**Purpose**: Store tasks between program executions

**Responsibilities**:

- Save tasks reliably
- Retrieve all tasks
- Update task states
- Initialize storage on first use

**Interface**:

- Input: Task data and operations
- Output: Stored tasks or confirmation

### Data Model

<!-- Show logical data model, not physical implementation -->

**Task Entity**:

- Unique identifier (system-generated)
- Title (user-provided, required)
- Completion status (boolean)
- Creation timestamp

**Storage Requirements**:

- Local to user's machine
- Persistent across restarts
- No network dependency

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**User Commands**:

1. **Create Task**
   - Input: Task description text
   - Output: Confirmation with identifier
   - Behavior: Stores new incomplete task

2. **List Tasks**
   - Input: None
   - Output: All tasks with status
   - Behavior: Shows identifier, title, completion

3. **Complete Task**
   - Input: Task identifier
   - Output: Confirmation
   - Behavior: Marks task as done

**Output Format Example**:

```text
ID    Task                    Status
---   --------------------    ------
abc   Fix login bug          [ ]
def   Review documentation   [✓]
ghi   Update dependencies    [ ]
```text

### Error Handling

- Empty task title → Reject with explanation
- Unknown task ID → Report not found
- Storage failures → User-friendly error
- Invalid commands → Show usage help

### Logging Strategy

**Applicable Use Cases**:

- [x] Error Tracking - Storage and operation failures
- [ ] Tracing - Not needed for simple operations
- [ ] User Behavior - Too early to track
- [ ] Performance - Not critical at this scale
- [ ] Security Audit - No sensitive operations
- [ ] Business Metrics - No metrics yet

**Implementation Details**:

```text
Error Tracking:
- Events: operation_failed
- Fields: operation_type, error_category
- Retention: 30 days
- Privacy: No task content in logs
```text

**Privacy Considerations**:

- Never log task content
- Only log operation metadata

## Security Considerations

- Input validation to prevent malicious data
- Storage accessible only to creating user
- No network communications
- No sensitive data handling

## Testing Strategy

- Component isolation: Test each layer independently
- Integration flow: Test complete operations
- Error scenarios: Invalid inputs, storage issues
- Data integrity: Tasks persist correctly

## Migration Plan

Not applicable - this is the initial implementation.

## Alternatives Considered

### Alternative 1: Web-based Interface

Provide browser UI instead of command line.

**Why not chosen**: Adds complexity, requires running server, harder to integrate with terminal workflow,
contradicts "minimal" goal.

### Alternative 2: Cloud Storage

Store tasks in online service.

**Why not chosen**: Requires internet, adds authentication complexity, privacy concerns, contradicts
local-first principle.

### Alternative 3: No Persistence

Keep tasks only during program execution.

**Why not chosen**: Not useful for real task tracking, users expect data to persist, would require fundamental
redesign later.
````
