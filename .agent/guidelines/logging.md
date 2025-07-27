# Logging Conventions for AI Agents

## Overview

This document defines logging conventions based on use cases. Logs are classified solely by their intended use,
without traditional log levels. When implementing logging in design documents, specify which use cases apply and how
logs will be structured for those specific purposes.

## Logging Use Cases and What to Record

### 1. Tracing

**Purpose**: Track execution flow and operation dependencies across the system

**What to Record**:

- Trace ID (unique per request/operation)
- Span ID and parent span ID
- Operation name
- Start and end timestamps
- Duration
- Status (success/failure)
- Input/output summary (sanitized)

**Example**:

````json
{
  "use_case": "tracing",
  "trace_id": "abc-123-def",
  "span_id": "span-456",
  "parent_span_id": "span-123",
  "operation": "parse_natural_language",
  "start_time": "2024-01-20T10:30:45.123Z",
  "duration_ms": 45,
  "status": "success",
  "input_summary": "task creation request"
}
```text

### 2. Error Tracking

**Purpose**: Identify, diagnose, and resolve production issues

**What to Record**:

- Error type and message
- Stack trace
- Trace ID (to correlate with request flow)
- User context
- System state
- Suggested remediation

**Example**:

```json
{
  "use_case": "error_tracking",
  "trace_id": "abc-123-def",
  "error_type": "ParseError",
  "message": "Unable to parse date from input",
  "stack_trace": "...",
  "user_input": "remind me yesterday",
  "suggestion": "Date must be in the future"
}
```text

### 3. User Behavior

**Purpose**: Understand how users interact with the system to improve UX

**What to Record**:

- User actions
- Feature usage
- Success/failure patterns
- Time to complete tasks
- User preferences used

**Example**:

```json
{
  "use_case": "user_behavior",
  "trace_id": "abc-123-def",
  "action": "create_task",
  "input_method": "natural_language",
  "success": true,
  "time_to_complete_ms": 200,
  "features_used": ["reminder", "due_date"]
}
```text

### 4. Performance Monitoring

**Purpose**: Identify bottlenecks and optimize system performance

**What to Record**:

- Operation metrics
- Resource usage
- Queue depths
- Cache performance
- External service latencies

**Example**:

```json
{
  "use_case": "performance",
  "trace_id": "abc-123-def",
  "operation": "list_tasks",
  "total_time_ms": 150,
  "db_query_time_ms": 100,
  "cache_hit": false,
  "result_count": 42,
  "memory_used_mb": 45
}
```text

### 5. Security Audit

**Purpose**: Maintain compliance and provide security audit trails

**What to Record**:

- Who did what when
- Access attempts
- Data modifications
- Permission checks
- Administrative actions

**Example**:

```json
{
  "use_case": "security_audit",
  "trace_id": "abc-123-def",
  "event": "data_export",
  "user_id": "admin_456",
  "ip_address": "192.168.1.100",
  "exported_records": 1500,
  "timestamp": "2024-01-20T10:30:45Z"
}
```text

### 6. Business Metrics

**Purpose**: Track business KPIs and product analytics

**What to Record**:

- Feature adoption
- User engagement
- Conversion events
- Revenue-related actions
- Growth metrics

**Example**:

```json
{
  "use_case": "business_metrics",
  "trace_id": "abc-123-def",
  "event": "feature_adoption",
  "feature": "team_collaboration",
  "user_segment": "enterprise",
  "action": "first_use"
}
```text

## Implementation Guidelines

### Mandatory Fields

Every log entry must include:

- `use_case`: One of the six defined use cases
- `timestamp`: ISO 8601 format
- `trace_id`: For correlating related operations

### Design Document Template

When creating a design document, include:

```markdown
## Logging Strategy

This feature implements logging for:

**Tracing**:

- All major operations with timing
- Parent-child relationships for sub-operations
- Retention: 7 days

**Error Tracking**:

- All errors with sanitized context
- Stack traces for unexpected errors
- Retention: 30 days

**Performance Monitoring**:

- Operations taking >100ms
- Resource usage for memory-intensive operations
- Retention: 7 days
```text

## Use Case Selection Guide

| If you need to...        | Use this case      |
| ------------------------ | ------------------ |
| Follow execution flow    | `tracing`          |
| Debug production issues  | `error_tracking`   |
| Understand user patterns | `user_behavior`    |
| Find slow operations     | `performance`      |
| Audit who did what       | `security_audit`   |
| Track business KPIs      | `business_metrics` |

## Privacy and Security

### Never Log

- Passwords, tokens, API keys
- Full credit card numbers
- Unencrypted PII
- Private message contents

### Always Sanitize

- User input (mask sensitive patterns)
- File paths (use hashes)
- Email addresses (partially mask)
- IP addresses (optional geo-aggregation)

## Technical Notes

### Correlation with Trace IDs

All use cases should include `trace_id` to enable cross-use-case correlation:

```bash
# Find all logs for a specific request
jq 'select(.trace_id == "abc-123-def")' logs.json | jq -s 'sort_by(.timestamp)'
```text

### Development vs Production

In development environments, tracing logs may include additional debug information:

```json
{
  "use_case": "tracing",
  "trace_id": "abc-123-def",
  "operation": "parse_task",
  "debug_info": {  // Only in development
    "raw_input": "buy milk tomorrow",
    "parser_state": {...}
  }
}
```text

### Sampling Strategies

For high-volume use cases:

```json
{
  "use_case": "tracing",
  "sampled": true,
  "sample_rate": 0.1,  // 10% sampling
  ...
}
```text

## Testing Requirements

1. Verify `use_case` field is valid
2. Ensure `trace_id` is present
3. Check timestamp format
4. Validate no sensitive data exposed
5. Test correlation across use cases
6. Verify sampling works correctly
````
