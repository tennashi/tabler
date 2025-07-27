# Design Doc: LLM-based Metadata Extraction

## Overview

This feature uses Claude Code to extract metadata from natural language task input in a single operation:
deadline dates, categories/tags, and priority levels. Instead of multiple regex patterns or separate API calls,
we leverage Claude's language understanding to extract all metadata at once.

## Background

[Link to PRD: ../../prd/smart_task_creation.md](../../prd/smart_task_creation.md)

This design implements requirements from the Smart Task Creation PRD:

- Story 3: AI-Powered Understanding (natural language dates and categories)
- Story 2: Natural Language Input with Shortcuts (enhanced with AI)
- Should Have: Natural language processing for dates and categories

## Goals

- Extract dates from natural language ("by tomorrow", "next week", "来週の金曜日")
- Infer categories/tags from task content
- Suggest priority based on keywords and context
- Single LLM call for all metadata extraction (performance)
- Support Japanese and English equally
- Sub-2s response time for AI features
- Graceful degradation when Claude unavailable

## Non-Goals

- Real-time typing suggestions (this is batch extraction)
- Learning from user corrections (covered in Context Learning feature)
- Custom extraction rules (future enhancement)
- Supporting languages beyond English and Japanese

## Design

### High-Level Architecture

````text
┌─────────────┐     ┌──────────────────┐     ┌─────────────┐
│   CLI Add   │────▶│ Metadata Service │────▶│  Storage    │
│   Command   │     │                  │     │  (SQLite)   │
└─────────────┘     └──────────────────┘     └─────────────┘
                             │
                             ▼
                    ┌──────────────────┐
                    │   Claude Code    │
                    │   Subprocess     │
                    └──────────────────┘
```text

### Detailed Design

#### Component 1: Metadata Service

**Purpose**: Orchestrates metadata extraction from task input

**Responsibilities**:

- Check for existing shortcuts (@, #, !)
- Invoke Claude for natural language processing
- Cache results for performance
- Handle errors gracefully

**Interface**:

- Input: Raw task string, current timestamp
- Output: Extracted metadata structure (cleaned text, deadline, tags, priority)

#### Component 2: Claude Client

**Purpose**: Manages communication with Claude Code subprocess

**Responsibilities**:

- Format prompts for metadata extraction
- Execute Claude subprocess with timeout
- Parse structured JSON responses
- Handle subprocess errors

**Interface**:

- Input: Task text, context (date, timezone)
- Output: Structured metadata or error

#### Component 3: Metadata Cache

**Purpose**: Reduce redundant LLM calls

**Responsibilities**:

- Cache extraction results by normalized input
- Implement LRU eviction
- Handle TTL expiration

**Interface**:

- Get/Set operations with automatic expiry
- Cache hit rate monitoring

### Data Model

```sql
-- No schema changes needed for basic metadata
-- Uses existing task table columns:
-- deadline (DATETIME)
-- tags (JSON array in metadata column)
-- priority (TEXT)

-- Cache stored in memory, not persisted
```text

### API Design

**CLI Command Enhancement**:

```bash
# Input
tabler add urgent: finish report by tomorrow #work

# Output
📋 Extracted metadata:
  📅 Deadline: Jan 16, 2024
  🏷️ Tags: work, report
  ⚡ Priority: high

✅ Task created: "finish report"
```text

**Claude Prompt Format**:

```json
{
  "task_input": "urgent: finish report by tomorrow #work",
  "current_datetime": "2024-01-15T10:00:00Z",
  "timezone": "Asia/Tokyo",
  "request": "extract_metadata"
}
```text

**Claude Response Format**:

```json
{
  "cleaned_text": "finish report",
  "deadline": "2024-01-16",
  "tags": ["work", "report"],
  "priority": "high",
  "confidence": 0.92,
  "reasoning": "urgent keyword and tomorrow deadline indicate high priority"
}
```text

**Tag Handling Enhancement**:

- Tags are stored as-is without normalization (preserving user intent)
- Claude will understand semantic relationships between tags:
  - `work`, `Work`, `仕事`, `お仕事` recognized as related
  - Typos and variations handled intelligently
- Future tag search will use Claude for fuzzy matching:
  - Search for "仕事" finds tasks tagged with "work"
  - Search for "job" finds "仕事", "work", "業務"
- Benefits over strict normalization:
  - Preserves cultural/linguistic nuances
  - Handles mixed language environments naturally
  - No information loss from aggressive normalization

### Error Handling

- **Claude Timeout**: Fall back to shortcut parsing only
- **Invalid Response**: Log error, continue with original input
- **Subprocess Failure**: Graceful degradation, task creation continues
- **Cache Errors**: Log and continue without cache

### Logging Strategy

**Applicable Use Cases**:

- [x] Performance - Track Claude response times
- [x] Error Tracking - Claude failures and fallbacks
- [x] User Behavior - Which metadata gets extracted
- [ ] Tracing - Not needed for this feature
- [ ] Security Audit - No sensitive operations
- [ ] Business Metrics - Covered by User Behavior

**Implementation Details**:

```text
Performance:
- Event: "metadata_extraction_complete"
- Fields: duration_ms, cache_hit, input_length
- Retention: 30 days
- Sampling: 100% (low volume)

Error Tracking:
- Event: "metadata_extraction_error"
- Fields: error_type, error_message, fallback_used
- Retention: 90 days

User Behavior:
- Event: "metadata_extracted"
- Fields: has_deadline, tag_count, priority_detected, language
- Retention: 90 days
```text

**Privacy Considerations**:

- Do not log task content
- Hash inputs for cache keys
- No PII in error messages

## Security Considerations

- Input sanitization before subprocess execution
- Timeout enforcement (2 seconds max)
- No shell injection via controlled subprocess args
- Claude runs with user privileges only

## Testing Strategy

- **Unit tests**: Mock Claude responses, test parsing logic
- **Integration tests**: Real Claude subprocess calls
- **Performance tests**: Ensure <2s response time
- **Error scenarios**: Timeouts, invalid responses, subprocess failures

## Migration Plan

No migration needed - this is an enhancement to existing task creation.

## Alternatives Considered

### Alternative 1: Separate LLM Calls

Make individual Claude calls for each metadata type (date, tags, priority).

**Why not chosen**:

- 3x latency (multiple round trips)
- 3x cost for Claude API
- Less context for better inference

### Alternative 2: Rule-Based Extraction

Use regex patterns and keyword matching without LLM.

**Why not chosen**:

- Limited to predefined patterns
- Poor multilingual support
- Can't understand context and nuance
````
