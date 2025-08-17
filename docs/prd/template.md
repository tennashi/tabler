# PRD: [Feature Name]

<!--
PRD WRITING GUIDELINES:

PRDs should focus on WHY and WHAT, not HOW:

✅ DO (Why/What):
- WHY: Explain the problem and user needs
- WHY: Describe the business value and impact
- WHAT: Define user-facing functionality
- WHAT: Specify desired outcomes and behaviors
- WHAT: Describe success criteria

❌ DON'T (How):
- HOW: Implementation details ("single-line input", "Enter key")
- HOW: Technical architecture ("LLM API", "SQLite database")
- HOW: Internal processing ("parsing system", "AI processes")
- HOW: Specific response times ("100ms", "2 seconds")

Examples:
❌ "Single-line task input with Enter to save"
✅ "Minimal friction task capture interface"

❌ "Use LLM API service for natural language processing"
✅ "Natural language understanding capability"

❌ "Task creation completes in <200ms"
✅ "Instant feedback for all user actions"

Technical implementation details belong in the Design Doc.
-->

**Version**: 1.0

## Problem Statement

[What problem are we solving? Why does this problem exist?]

## Objectives

- [Primary objective]
- [Secondary objectives]

## User Stories

As a [type of user], I want [goal] so that [benefit].

### Story 1: [Title]

- **As a**: [User type]
- **I want**: [Action/feature]
- **So that**: [Benefit/value]

#### Scenario: [Story 1 primary scenario]

- **Given**: [Preconditions/initial state]
- **When**: [User action/event]
- **Then**: [Expected result/state change]

<!--
💡 Use Given-When-Then format for testable conditions:
- Given: Test preconditions or initial state
- When: User actions or system events
- Then: Expected results or state changes

Example:
Scenario: Successful task creation
- Given: User is on the task list page
- When: User enters "Buy milk" and presses Enter
- Then: Task "Buy milk" appears in the task list with pending status
-->

#### Scenario: [Story 1 alternative scenario]

- **Given**: [Preconditions/initial state]
- **When**: [User action/event]
- **Then**: [Expected result/state change]

### Story 2: [Title]

- **As a**: [User type]
- **I want**: [Action/feature]
- **So that**: [Benefit/value]

#### Scenario: [Story 2 primary scenario]

- **Given**: [Preconditions/initial state]
- **When**: [User action/event]
- **Then**: [Expected result/state change]

#### Scenario: [Story 2 alternative scenario]

- **Given**: [Preconditions/initial state]
- **When**: [User action/event]
- **Then**: [Expected result/state change]

## Requirements

### Functional Requirements

#### Must Have

- [Critical features that must be included]

#### Should Have

- [Important features that should be included if possible]

#### Nice to Have

- [Features that would be nice but not essential]

### Non-Functional Requirements

- **Performance**: [User-perceived speed, not specific numbers]
- **Usability**: [User experience requirements]
- **Reliability**: [Availability from user perspective]

## Success Metrics

- [How will we measure success?]
- [What are the key performance indicators?]

## Scope

### In Scope

- [What is included in this feature]

### Out of Scope

- [What is explicitly not included]

## Open Questions

- [Questions that need to be answered]

## Implementation Phases (if applicable)

- **Phase 1**: [Core MVP features]
- **Phase 2**: [Enhanced features]
- **Phase 3**: [Advanced features]

## Dependencies

- [Other features or capabilities this depends on]

<!-- Note: Avoid specifying technical implementations like specific APIs or databases -->
