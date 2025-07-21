# Design Doc: Task Shortcuts

<!-- 
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

This feature adds inline shortcuts to task creation for quick entry of metadata. Users can specify tags (#), deadlines (@), and priorities (!) directly within task descriptions, making task creation more efficient and natural.

## Background

[Link to PRD: ../prd/smart_task_creation.md](../prd/smart_task_creation.md)

This implements Story 2 from the PRD: Natural Language Input with Shortcuts. Users need a fast way to add metadata without breaking their flow or using complex forms.

## Goals

- Parse shortcuts from task descriptions during creation
- Support tags with # prefix (multiple allowed)
- Support deadlines with @ prefix (basic date formats)
- Support priority with ! marks (1-3 levels)
- Remove shortcuts from final task title
- Maintain backward compatibility with plain tasks

## Non-Goals

- Complex natural language date parsing (use LLM feature for that)
- Nested or quoted shortcuts
- Shortcut escaping mechanisms
- Shortcuts in update commands
- Custom shortcut symbols
- Tag hierarchy or categories

## Design

### High-Level Architecture

```
┌─────────────────┐     ┌──────────────┐     ┌─────────────┐
│ Task Input      │────▶│ Shortcut     │────▶│ Persistence │
│ with Shortcuts  │     │ Parser       │     │ Layer       │
└─────────────────┘     └──────────────┘     └─────────────┘
                               │
                               ▼
                        ┌──────────────┐
                        │ Parse Result │
                        │ - Clean title│
                        │ - Metadata   │
                        └──────────────┘
```

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: Shortcut Parser

**Purpose**: Extract metadata from inline shortcuts

**Responsibilities**:
- Identify shortcut patterns in input text
- Extract values for each shortcut type
- Return cleaned title without shortcuts
- Preserve text that isn't shortcuts

**Interface**:
- Input: Raw task text with potential shortcuts
- Output: Cleaned title and extracted metadata

#### Component 2: Tag Processor

**Purpose**: Handle tag shortcuts (#tag)

**Responsibilities**:
- Find all hashtag patterns
- Extract tag values
- Normalize tags for consistency
- Handle edge cases (empty tags, special characters)

**Interface**:
- Input: Text containing potential tags
- Output: List of normalized tag values

#### Component 3: Date Processor

**Purpose**: Handle deadline shortcuts (@date)

**Responsibilities**:
- Recognize date patterns after @ symbol
- Support common formats (today, tomorrow, dates)
- Convert to standard date representation
- Handle invalid date inputs gracefully

**Interface**:
- Input: Date shortcut text
- Output: Standardized date or none

#### Component 4: Priority Processor

**Purpose**: Handle priority shortcuts (!)

**Responsibilities**:
- Count priority indicators
- Map to priority levels
- Handle maximum limits
- Process multiple occurrences correctly

**Interface**:
- Input: Text with priority marks
- Output: Priority level (none/low/medium/high)

### Data Model

<!-- Show logical data model, not physical implementation -->

**Extended Task Properties**:
- Existing: identifier, title, completion status, created time
- New: tags (list of strings)
- New: deadline (optional date)
- New: priority level (enumeration)

**Metadata Characteristics**:
- Tags: Multiple allowed, stored as-is without normalization
- Deadline: Single value, optional
- Priority: Single level, default none

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

**Shortcut Syntax**:

| Pattern | Purpose | Examples |
|---------|---------|----------|
| `#word` | Add tag | `#work`, `#urgent`, `#project-x` |
| `@date` | Set deadline | `@today`, `@tomorrow`, `@2024-01-15` |
| `!` | Set priority | `!` (low), `!!` (medium), `!!!` (high) |

**Processing Examples**:

```
Input: "Fix login bug #work #urgent @tomorrow !!"
Result:
  Title: "Fix login bug"
  Tags: ["work", "urgent"]
  Deadline: next day's date
  Priority: medium

Input: "Buy milk @today"
Result:
  Title: "Buy milk"
  Tags: []
  Deadline: current date
  Priority: none
```

**Date Format Support**:
- Relative: today, tomorrow
- Day names: monday, tuesday, etc.
- Standard dates: various common formats

### Error Handling

- Invalid date → Skip deadline, keep processing
- Empty tag → Ignore that tag
- Too many priority marks → Use maximum
- Malformed shortcuts → Treat as regular text

### Logging Strategy

**Applicable Use Cases**:
- [x] User Behavior - Track shortcut usage patterns
- [x] Error Tracking - Failed parsing attempts
- [ ] Tracing - Not needed
- [ ] Performance - Parsing is fast
- [ ] Security Audit - No security implications
- [ ] Business Metrics - Covered by user behavior

**Implementation Details**:
```
User Behavior:
- Events: shortcuts_used
- Fields: shortcut_types, count_per_type
- Retention: 90 days

Error Tracking:
- Events: parse_failed
- Fields: shortcut_type, error_reason
- Retention: 30 days
```

**Privacy Considerations**:
- Log shortcut usage patterns only
- Never log actual tag names or dates

## Security Considerations

- Validate all parsed values
- Prevent injection through shortcut values
- Normalize data to prevent duplicates
- Limit lengths to reasonable values

## Testing Strategy

- Parser components: Test each shortcut type
- Integration: Combined shortcuts in one input
- Edge cases: Malformed shortcuts, empty values
- Compatibility: Plain tasks still work
- Normalization: Consistent output

## Migration Plan

1. Extend task data model with new fields
2. Existing tasks get default values
3. No data migration needed
4. Full backward compatibility

## Alternatives Considered

### Alternative 1: Separate Metadata Fields

Provide distinct input fields for each metadata type.

**Why not chosen**: Disrupts flow, requires multiple inputs, less natural than inline shortcuts, poor command-line experience.

### Alternative 2: Configuration File

Define tasks in structured configuration files.

**Why not chosen**: Too heavyweight for quick capture, requires editing files, loses spontaneity of task creation.

### Alternative 3: Natural Language Only

Use only AI to extract all metadata.

**Why not chosen**: Requires network/AI availability, slower than shortcuts, less predictable, overkill for simple metadata.

### Alternative 4: Tag Normalization at Input Time

Normalize tags when parsing (e.g., lowercase, trim spaces, unify similar tags).

**Why not chosen**: 
- Complex in multilingual environments (Japanese/English mixed tags)
- Users may intentionally use different cases or variations
- Example: `Work`, `work`, `仕事`, `お仕事` have different nuances
- Better to defer to LLM-based semantic grouping in future features
- Allows fuzzy search and intelligent tag suggestions later