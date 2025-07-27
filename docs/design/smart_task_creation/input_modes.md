# Design Doc: Input Mode System

<!--
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

This feature implements three distinct input modes (Quick, Talk, Planning) that users can control via command
prefixes or automatic detection. Each mode optimizes for different task creation scenarios, from rapid capture
to comprehensive planning.

## Background

[Link to PRD: ../../prd/smart_task_creation.md](../../prd/smart_task_creation.md)

This design implements Story 6 from the Smart Task Creation PRD: Mode Control. Users work differently in
different contexts - sometimes needing quick capture, sometimes wanting dialogue, sometimes requiring detailed
planning. This system provides the right tool for each situation.

## Goals

- Provide three distinct modes: Quick, Talk, and Planning
- Enable manual mode selection via command prefixes (/q, /t, /p)
- Implement intelligent automatic mode detection
- Show clear mode indicators to users
- Maintain mode context during task creation
- Ensure seamless switching between modes

## Non-Goals

- Persistent mode settings across sessions
- Custom mode creation
- Mode-specific UI themes
- Learning mode preferences (covered in context learning)
- Sub-modes or mode combinations

## Design

### High-Level Architecture

````text
┌─────────────┐     ┌──────────────────┐     ┌─────────────┐
│   CLI Add   │────▶│   Mode Manager   │────▶│ Mode Handler│
│   Command   │     │                  │     │  (Q/T/P)    │
└─────────────┘     └──────────────────┘     └─────────────┘
                             │                        │
                             ▼                        ▼
                    ┌──────────────────┐     ┌─────────────┐
                    │ Mode Detector    │     │  Storage    │
                    │ (Claude-based)   │     │  (SQLite)   │
                    └──────────────────┘     └─────────────┘
```text

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Mode Manager

**Purpose**: Central coordinator for mode selection and routing

**Responsibilities**:

- Parse command input for mode prefixes
- Route to appropriate mode handler
- Maintain current mode state
- Coordinate mode transitions

**Interface**:

- Input: Raw command arguments
- Output: Processed task via appropriate mode

#### Component 2: Mode Detector

**Purpose**: Intelligently determine appropriate mode when not specified

**Responsibilities**:

- Analyze input characteristics
- Apply detection heuristics
- Use Claude for complex detection
- Return mode recommendation with confidence

**Interface**:

- Input: Task text without mode prefix
- Output: Recommended mode with confidence score

#### Component 3: Quick Mode Handler

**Purpose**: Optimize for speed and minimal interaction

**Responsibilities**:

- Process task with minimal parsing
- Apply basic shortcuts (@, #, !)
- Skip all optional enhancements
- Create task immediately

**Interface**:

- Input: Task text
- Output: Created task (no interaction)

#### Component 4: Talk Mode Handler

**Purpose**: Enable conversational task refinement

**Responsibilities**:

- Initiate clarification dialogue
- Process conversational exchanges
- Extract task details from dialogue
- Handle skip requests

**Interface**:

- Input: Initial task text
- Output: Refined task after dialogue

#### Component 5: Planning Mode Handler

**Purpose**: Facilitate comprehensive task breakdown

**Responsibilities**:

- Analyze task complexity
- Generate decomposition suggestions
- Handle user choices on subtasks
- Create task hierarchies

**Interface**:

- Input: Complex task description
- Output: Parent task with subtasks

### Data Model

<!-- Show logical data model, not physical implementation -->

**Mode State** (session only):

- Current mode (quick/talk/planning)
- Mode override flag (manual vs auto)
- Mode confidence (for auto detection)

**No persistent storage** for mode preferences in this phase

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**Command Patterns**:

```text
# Explicit mode selection
tabler add /quick buy milk
tabler add /q buy milk

tabler add /talk prepare for meeting
tabler add /t prepare for meeting

tabler add /plan organize conference
tabler add /p organize conference

# Automatic mode detection
tabler add buy milk                    # → Quick mode
tabler add discuss project status      # → Talk mode  
tabler add plan company retreat        # → Planning mode
```text

**Mode Indicators**:

- Quick: Minimal output, just confirmation
- Talk: Shows dialogue prompts with 🤔 emoji
- Planning: Shows decomposition options with 📋 emoji

**Detection Heuristics**:

1. Very short input (<10 chars) → Quick
2. Contains question words → Talk
3. Contains planning keywords → Planning
4. High complexity score → Planning
5. Default → Quick

### Error Handling

- Invalid mode prefix: Show available modes
- Mode detection failure: Default to Quick
- Claude unavailable: Use simple heuristics
- Mode handler errors: Fallback to basic creation

### Logging Strategy

**Applicable Use Cases**:

- [x] User Behavior - Track mode usage patterns
- [x] Performance - Mode detection latency
- [ ] Error Tracking - Minimal errors expected
- [ ] Tracing - Not needed
- [ ] Security Audit - No security implications
- [x] Business Metrics - Mode effectiveness

**Implementation Details**:

```yaml
User Behavior:
- Events: mode_selected, mode_detected, mode_overridden
- Fields: mode_type, selection_method, task_length
- Retention: 90 days

Performance:
- Events: mode_detection_time
- Fields: detection_ms, used_claude, confidence
- Retention: 30 days

Business Metrics:
- Events: mode_completion_rate
- Fields: mode_type, completed_successfully
- Retention: 90 days
```text

**Privacy Considerations**:

- Log mode usage patterns only
- No task content in logs
- Aggregate statistics only

## Security Considerations

- Validate mode prefixes to prevent injection
- No mode-specific security risks
- Standard input sanitization applies
- No special permissions per mode

## Testing Strategy

- Unit tests: Mode detection logic, prefix parsing
- Integration tests: Full flow for each mode
- User studies: Mode selection accuracy
- Performance tests: Detection speed
- Edge cases: Ambiguous inputs, mixed signals

## Migration Plan

No migration needed - modes are additive to existing functionality.

## Alternatives Considered

### Alternative 1: Persistent Mode Setting

Keep users in their selected mode until changed.

**Why not chosen**: Confusing state, users forget current mode, requires mode status display, adds complexity
without clear benefit.

### Alternative 2: Learning-Based Mode Selection

Learn user's preferred modes for different contexts.

**Why not chosen**: Requires usage history, adds complexity, covered in separate context learning feature, may
be unpredictable.

### Alternative 3: Combined Modes

Allow mixing modes (e.g., quick + planning).

**Why not chosen**: Confusing UX, unclear behavior, modes designed to be distinct approaches, would complicate
implementation significantly.
````
