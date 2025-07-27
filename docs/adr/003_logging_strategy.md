# ADR-003: Logging Strategy

**Status**: Proposed

## Context

The Tabler project needs an effective logging strategy. Traditional log levels (ERROR, WARN, INFO, DEBUG) often
don't align with actual use cases, making it difficult to find necessary information.

Logs serve various purposes:

- Debugging during development
- Error tracking in production
- User behavior analysis
- Performance monitoring
- Security auditing
- Business metrics collection

Each purpose requires different information, and a single log level hierarchy cannot adequately classify them.

## Decision

**Adopt use-case-based log classification**. Instead of log levels, classify logs by their intended use with these
six categories:

1. **Tracing** (`tracing`) - Track execution flow and timing
2. **Error Tracking** (`error_tracking`) - Identify and resolve issues
3. **User Behavior** (`user_behavior`) - Analyze for UX improvements
4. **Performance** (`performance`) - Monitor resource usage and efficiency
5. **Security Audit** (`security_audit`) - Compliance and audit trails
6. **Business Metrics** (`business_metrics`) - KPIs and business analytics

Each log entry must include:

- `use_case`: One of the six categories above
- `timestamp`: ISO 8601 format
- `trace_id`: ID to correlate related operations

## Consequences

### Positive

- **Clear purpose**: Each log's intended use is explicit, making information easier to find
- **Appropriate fields**: Can define necessary fields per use case
- **Efficient filtering**: Simple filtering by use case
- **Optimized retention**: Can set appropriate retention periods per use case
- **Business value**: Logs directly align with business value

### Negative

- **Learning curve**: Developers need to learn the new classification method
- **Tool compatibility**: Some existing log analysis tools expect traditional log levels
- **Migration cost**: Need to migrate existing log outputs (no impact as implementation hasn't started)

### Neutral

- **Debug information**: Can include additional information in tracing logs for development
- **Severity**: Can add `severity` or `priority` attributes when needed
- **Sampling**: Can apply sampling strategies for high-volume use cases

## Options Considered

### Option 1: Traditional Log Levels

Use standard ERROR, WARN, INFO, DEBUG levels.

- ✅ Compatible with many tools
- ✅ Familiar to developers
- ❌ Doesn't match actual use cases
- ❌ Different purposes mixed at same level

### Option 2: Log Levels + Categories Hybrid

Add categories on top of traditional log levels.

- ✅ Maintains compatibility with existing tools
- ✅ Allows more detailed classification
- ❌ Complex with dual classification
- ❌ Ambiguous combinations of levels and categories

### Option 3: Use-Case-Based Classification (Selected)

Classify only by use case, no log levels.

- ✅ Simple and clear
- ✅ Fields optimized for use case
- ✅ Directly tied to business value
- ❌ Compatibility challenges with some tools
- ❌ Requires learning new approach
