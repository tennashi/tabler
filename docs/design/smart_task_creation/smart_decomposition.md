# Design Doc: Smart Task Decomposition

<!--
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

This feature analyzes complex tasks and suggests logical subtask breakdowns using Claude Code. When users input
tasks that appear too large or vague, the system offers to decompose them into manageable, actionable subtasks
while maintaining parent-child relationships.

## Background

[Link to PRD: ../../prd/smart_task_creation.md](../../prd/smart_task_creation.md)

This design implements Story 5 from the Smart Task Creation PRD: Smart Task Decomposition. Users often create
tasks that are too broad ("plan company event") which can be overwhelming. This feature helps break them down
into concrete steps.

## Goals

- Detect complex tasks that would benefit from decomposition
- Generate logical, actionable subtasks using LLM
- Present suggestions interactively for user control
- Maintain parent-child task relationships
- Complete the entire flow in under 2 seconds
- Support both English and Japanese task decomposition

## Non-Goals

- Automatic decomposition without user consent
- Multi-level hierarchies (only single parent-child level)
- Project management features (dependencies, critical path)
- Time estimation for subtasks

## Design

### High-Level Architecture

````text
┌─────────────┐     ┌──────────────────┐     ┌─────────────┐
│   CLI Add   │────▶│ Decomposition    │────▶│  Storage    │
│   Command   │     │    Service       │     │  (SQLite)   │
└─────────────┘     └──────────────────┘     └─────────────┘
                             │
                             ▼
                    ┌──────────────────┐
                    │   Claude Code    │
                    │   Subprocess     │
                    └──────────────────┘
```text

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Complexity Detector

**Purpose**: Identify tasks that would benefit from decomposition

**Responsibilities**:

- Analyze task characteristics that indicate complexity
- Determine if decomposition should be offered
- Provide confidence score for the decision

**Interface**:

- Input: Task text from user
- Output: Decision (should decompose) with confidence score

#### Component 2: Task Decomposer

**Purpose**: Generate subtask suggestions using Claude

**Responsibilities**:

- Formulate appropriate prompts for Claude
- Request subtask decomposition
- Validate and structure the response
- Handle edge cases (too few/many suggestions)

**Interface**:

- Input: Complex task text
- Output: List of suggested subtasks with decomposition rationale

#### Component 3: Interactive Presenter

**Purpose**: Present decomposition options and handle user decisions

**Responsibilities**:

- Display suggestions in a clear, scannable format
- Offer user choices (accept all, edit, skip)
- Process user modifications to suggestions
- Coordinate final task creation

**Interface**:

- Input: Suggested subtasks from decomposer
- Output: User's final decision and task list

### Data Model

<!-- Show logical data model, not physical implementation -->

**Task Hierarchy**:

- Tasks can have a parent_id referencing another task
- Parent tasks represent the original complex task
- Child tasks are the decomposed subtasks
- Only one level of nesting supported

**Relationships**:

- One parent can have multiple children
- Children cannot have their own children
- Deleting parent should handle children appropriately

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**User Interaction Flow**:

1. User inputs: "plan company offsite for 50 people"
2. System detects complexity and offers decomposition
3. System presents 3-7 suggested subtasks
4. User can:
   - Accept all suggestions
   - Edit individual suggestions
   - Skip decomposition entirely
5. System creates tasks based on user choice

**Claude Integration**:

- Request: Task text with decomposition request
- Response: Structured list of subtasks with rationale
- Timeout: 2 seconds maximum
- Fallback: Skip decomposition if Claude unavailable

### Error Handling

- Claude timeouts: Continue without decomposition
- Invalid suggestions: Show error, allow manual entry
- Database failures: Rollback all changes, show error
- User cancellation: Exit cleanly, no tasks created

### Logging Strategy

**Applicable Use Cases**:

- [x] User Behavior - Track decomposition acceptance patterns
- [x] Performance - Measure Claude response times
- [x] Error Tracking - Monitor decomposition failures
- [ ] Tracing - Not needed for this feature
- [ ] Security Audit - No sensitive operations
- [ ] Business Metrics - Covered by User Behavior

**Implementation Details**:

```yaml
User Behavior:
- Events: decomposition_offered, decomposition_accepted, decomposition_edited
- Fields: task_complexity_score, subtask_count, user_action
- Retention: 90 days
- Privacy: Hash task content, only track patterns

Performance:
- Events: decomposition_latency
- Fields: total_time_ms, claude_time_ms, subtask_count
- Retention: 30 days

Error Tracking:
- Events: decomposition_failed
- Fields: error_type, timeout_occurred, fallback_used
- Retention: 90 days
```text

**Privacy Considerations**:

- Never log actual task content
- Use hashed identifiers for pattern analysis
- Aggregate metrics only, no individual tracking

## Security Considerations

- Input sanitization before sending to Claude
- Prevent circular parent-child relationships
- Limit maximum subtasks to prevent abuse
- Transaction integrity for multi-task creation
- No shell injection through subprocess calls

## Testing Strategy

- Unit tests: Mock Claude responses, test detection logic
- Integration tests: Full flow with real Claude subprocess
- Performance tests: Verify <2s end-to-end latency
- Edge cases: Very long tasks, minimal tasks, different languages
- User acceptance: Interactive flow testing

## Migration Plan

Database changes needed:

1. Add parent_id column to tasks table
2. Add index for efficient parent-child queries
3. No data migration needed (existing tasks remain independent)

## Alternatives Considered

### Alternative 1: Template-Based Decomposition

Use predefined templates for common task types (e.g., "plan event" template).

**Why not chosen**: Too rigid, cannot handle unique tasks, requires maintaining template library, poor
internationalization support.

### Alternative 2: Automatic Decomposition

Decompose tasks automatically without user interaction.

**Why not chosen**: Removes user agency, may create unwanted tasks, difficult to undo, goes against principle
of user control.

### Alternative 3: Multi-Level Hierarchies

Allow subtasks to have their own subtasks (unlimited nesting).

**Why not chosen**: Adds significant complexity, difficult to display in CLI, most tasks only need single
decomposition level, harder to maintain.
````
