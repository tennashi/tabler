# Design Doc: Interactive Task Clarification

**Status**: Implemented ✅

<!-- 
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

This feature provides an AI-powered dialogue system that helps users clarify vague or incomplete tasks through contextual questions. When users input ambiguous tasks, the system engages in a brief conversation to extract actionable details before creating the task.

## Background

[Link to PRD: ../../prd/smart_task_creation.md](../../prd/smart_task_creation.md)

This design implements Story 4 from the Smart Task Creation PRD: Task Clarification Through Dialogue. Users often have fuzzy ideas that need refinement. Rather than creating vague tasks, this feature helps transform thoughts into clear action items through conversation.

## Goals

- Detect when task input is vague or needs clarification
- Generate contextual, helpful questions using Claude
- Maintain conversation state during dialogue
- Allow users to skip dialogue at any point
- Transform vague inputs into clear, actionable tasks
- Keep dialogues brief and focused (2-3 exchanges max)

## Non-Goals

- Long-form conversational UI
- General chatbot functionality
- Learning from past dialogues (covered in context learning)
- Voice-based interaction
- Multi-task planning sessions

## Design

### High-Level Architecture

```
┌─────────────┐     ┌──────────────────┐     ┌─────────────┐
│   CLI Add   │────▶│   Dialogue       │────▶│  Storage    │
│  (Talk Mode)│     │   Manager        │     │  (SQLite)   │
└─────────────┘     └──────────────────┘     └─────────────┘
                             │
                             ▼
                    ┌──────────────────┐
                    │   Claude Code    │
                    │   Subprocess     │
                    └──────────────────┘
```

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Vagueness Detector

**Purpose**: Identify when task input needs clarification

**Responsibilities**:
- Analyze task input for ambiguity indicators
- Determine if dialogue would be helpful
- Trigger appropriate interaction mode

**Interface**:
- Input: Raw task text
- Output: Clarity score and suggested dialogue trigger

#### Component 2: Dialogue Manager

**Purpose**: Orchestrate the clarification conversation

**Responsibilities**:
- Maintain conversation context
- Generate appropriate questions via Claude
- Process user responses
- Determine when enough information is gathered
- Enforce dialogue limits (2-3 exchanges)

**Interface**:
- Input: Initial vague task and user responses
- Output: Clarified task details or user skip signal

#### Component 3: Question Generator

**Purpose**: Create contextual questions using Claude

**Responsibilities**:
- Analyze what information is missing
- Generate natural, helpful questions
- Adapt based on previous responses
- Provide examples in questions when helpful

**Interface**:
- Input: Current task understanding and dialogue history
- Output: Next question or completion signal

#### Component 4: Response Processor

**Purpose**: Extract information from user answers

**Responsibilities**:
- Parse user responses for task details
- Update task understanding
- Detect when user wants to skip
- Merge clarifications into final task

**Interface**:
- Input: User response and current context
- Output: Updated task information or skip signal

### Data Model

<!-- Show logical data model, not physical implementation -->

**Session State** (temporary, not persisted):
- Original input
- Current understanding
- Dialogue history (questions and answers)
- Extracted details (deadline, tags, etc.)

**No permanent storage** for dialogues - only final task stored

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**User Interaction Example**:
```
$ tabler add /talk need to prepare for the thing

🤔 I'd like to help clarify this task. What thing do you need to prepare for?
> the team presentation next week

📅 When next week is the presentation, and what's the main topic?
> Thursday, about Q4 planning

✅ Got it! Creating task: "Prepare Q4 planning presentation for team meeting"
📅 Deadline: Thursday, Jan 25
🏷️ Tags: presentation, planning, team
```

**Skip Options**:
- User can type "skip" at any prompt
- Empty response continues with current understanding
- Ctrl+C cancels entire task creation

**Claude Integration**:
- Request: Current task understanding + dialogue history
- Response: Next clarifying question or "sufficient info" signal
- Context window: Include full dialogue for coherence

### Error Handling

- Claude timeout: Create task with current understanding
- Invalid responses: Re-ask with clarification
- User frustration detection: Offer to skip after 3 exchanges
- Cancellation: Exit cleanly, no task created

### Logging Strategy

**Applicable Use Cases**:
- [x] User Behavior - Track dialogue effectiveness
- [x] Performance - Measure dialogue completion time
- [ ] Error Tracking - Minimal errors expected
- [ ] Tracing - Not needed
- [ ] Security Audit - No sensitive operations
- [x] Business Metrics - Dialogue success rate

**Implementation Details**:
```
User Behavior:
- Events: dialogue_started, dialogue_completed, dialogue_skipped
- Fields: exchange_count, clarification_score, final_task_length
- Retention: 90 days
- Privacy: No actual dialogue content logged

Performance:
- Events: dialogue_response_time
- Fields: response_time_ms, exchange_number
- Retention: 30 days

Business Metrics:
- Events: dialogue_success
- Fields: initial_vagueness, final_clarity, user_satisfied
- Retention: 90 days
```

**Privacy Considerations**:
- Never log conversation content
- Only track metadata and patterns
- Hash any identifiers

## Security Considerations

- Sanitize all inputs before sending to Claude
- No persistent storage of conversations
- Timeout enforcement on all operations
- Rate limiting to prevent abuse
- No execution of user inputs

## Testing Strategy

- Unit tests: Mock dialogues, vagueness detection
- Integration tests: Full dialogue flows with Claude
- User studies: Test question quality and helpfulness
- Edge cases: Very vague inputs, non-cooperative users
- Performance: Ensure responsive conversation flow

## Migration Plan

No migration needed - this is a new interaction mode that doesn't affect existing data.

## Alternatives Considered

### Alternative 1: Form-Based Clarification

Show a form with fields for missing information.

**Why not chosen**: Less natural, requires users to understand task structure, poor mobile experience, doesn't handle unique cases well.

### Alternative 2: Suggestion-Based System

Provide multiple interpretations for user to choose.

**Why not chosen**: Limits possibilities, can miss user's actual intent, becomes overwhelming with many options, doesn't gather new information.

### Alternative 3: Always-On Dialogue

Make dialogue mandatory for all tasks.

**Why not chosen**: Annoying for clear tasks, slows down power users, reduces efficiency, goes against "quick capture" goal.