# Design Doc: Context Learning Foundation

<!-- 
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

This feature tracks user patterns in task creation to provide personalized suggestions and improve the task creation experience over time. The system learns from usage patterns while maintaining complete user privacy through local-only storage.

## Background

[Link to PRD: ../../prd/smart_task_creation.md](../../prd/smart_task_creation.md)

This implements the "Should Have" requirements for context learning from usage patterns, autocomplete for frequently used tags/categories, and personalized suggestions. All learning happens locally to protect user privacy.

## Goals

- Track task creation patterns (tags, priorities, time patterns)
- Provide autocomplete for frequently used tags
- Suggest relevant metadata based on context
- Improve mode detection accuracy over time
- Maintain complete user privacy (local-only)
- Allow users to control and clear their data

## Non-Goals

- Cross-device synchronization
- Sharing patterns between users  
- Complex ML models (keep it simple)
- Predicting task content
- Learning from task completion patterns

## Design

### High-Level Architecture

```
┌─────────────┐     ┌──────────────────┐     ┌─────────────┐
│   CLI Add   │────▶│ Context Learning │────▶│   Pattern   │
│   Command   │     │     Service      │     │   Storage   │
└─────────────┘     └──────────────────┘     │  (SQLite)   │
                             │                └─────────────┘
                             ▼
                    ┌──────────────────┐
                    │   Suggestion     │
                    │   Generator      │
                    └──────────────────┘
```

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Pattern Tracker

**Purpose**: Record usage patterns from task creation

**Responsibilities**:
- Extract patterns from created tasks
- Record frequency and recency of patterns
- Update pattern database
- Respect privacy settings

**Interface**:
- Input: Created task with metadata
- Output: Extracted patterns stored in database

#### Component 2: Pattern Storage

**Purpose**: Persist learned patterns locally

**Responsibilities**:
- Store pattern data in SQLite
- Implement retention policies
- Handle data expiration
- Support export/delete operations

**Interface**:
- Input: Pattern records
- Output: Stored patterns with metadata

#### Component 3: Suggestion Generator

**Purpose**: Generate contextual suggestions from patterns

**Responsibilities**:
- Query relevant patterns
- Rank suggestions by relevance
- Consider time context
- Format suggestions for display

**Interface**:
- Input: Current context (time, partial input)
- Output: Ranked suggestions list

#### Component 4: Autocomplete Provider

**Purpose**: Provide real-time completions during input

**Responsibilities**:
- Monitor user input
- Match against frequent patterns
- Provide fast completions
- Handle partial matches

**Interface**:
- Input: Partial tag or category
- Output: Completion suggestions

#### Component 5: Privacy Manager

**Purpose**: Give users control over their data

**Responsibilities**:
- Enable/disable learning
- Clear stored patterns
- Export user data
- Implement retention limits

**Interface**:
- Input: User privacy commands
- Output: Privacy action confirmation

### Data Model

<!-- Show logical data model, not physical implementation -->

**Pattern Types**:
- Tag frequency (tag → usage count)
- Time patterns (time of day → task types)
- Priority patterns (keywords → priority)
- Mode preferences (context → preferred mode)

**Pattern Record Structure**:
- Pattern type
- Pattern key
- Pattern value
- Frequency count
- Last used timestamp
- First seen timestamp

**Privacy Controls**:
- Learning enabled flag
- Retention period setting
- Excluded pattern list

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**Autocomplete Example**:
```
$ tabler add finish report #wo[TAB]
                            ^^^
Suggestions: #work (used 45 times)
            #workshop (used 3 times)
```

**Time-based Suggestions**:
```
$ tabler add [morning]
Suggested tags: #morning-routine, #exercise (based on past patterns)
```

**Privacy Commands**:
```
$ tabler privacy status
Learning: enabled
Patterns stored: 127
Retention: 90 days

$ tabler privacy clear
Clear all learned patterns? [y/N]

$ tabler privacy export
Exported patterns to: tabler-patterns-2024-01-15.json
```

### Error Handling

- Storage failures: Continue without learning
- Corrupt patterns: Skip and log
- Privacy conflicts: Respect user preference
- Performance issues: Disable temporarily

### Logging Strategy

**Applicable Use Cases**:
- [x] User Behavior - Track learning effectiveness
- [x] Performance - Pattern matching speed
- [ ] Error Tracking - Minimal errors expected
- [ ] Tracing - Not needed
- [ ] Security Audit - No security implications
- [x] Business Metrics - Feature adoption

**Implementation Details**:
```
User Behavior:
- Events: pattern_learned, suggestion_accepted, learning_disabled
- Fields: pattern_type, suggestion_rank, acceptance_rate
- Retention: 90 days
- Privacy: No actual patterns logged

Performance:
- Events: autocomplete_latency
- Fields: pattern_count, match_time_ms
- Retention: 30 days

Business Metrics:
- Events: learning_engagement
- Fields: patterns_stored, suggestions_used
- Retention: 90 days
```

**Privacy Considerations**:
- Never log actual task content
- Only track pattern types and counts
- Aggregate metrics only
- Hash any identifiers

## Security Considerations

- All data stored locally only
- No network transmission of patterns
- User owns all their data
- Clear data removal options
- No pattern sharing between users

## Testing Strategy

- Unit tests: Pattern extraction, suggestion ranking
- Integration tests: Full learning cycle
- Privacy tests: Data isolation, clearing
- Performance tests: Autocomplete speed
- Long-term tests: Pattern effectiveness over time

## Migration Plan

1. Add pattern tables to existing SQLite database
2. Start collecting patterns for new tasks
3. No migration of historical data needed
4. Gradual improvement as patterns accumulate

## Alternatives Considered

### Alternative 1: Cloud-Based Learning

Store patterns in cloud for cross-device sync.

**Why not chosen**: Privacy concerns, requires user accounts, adds infrastructure complexity, contradicts local-first principle.

### Alternative 2: Pre-Built Pattern Library

Ship with common patterns pre-loaded.

**Why not chosen**: Not personalized, may not match user needs, bloats installation, hard to maintain globally.

### Alternative 3: Complex ML Models

Use sophisticated ML for pattern recognition.

**Why not chosen**: Overkill for this use case, harder to explain to users, increased complexity, larger resource usage.