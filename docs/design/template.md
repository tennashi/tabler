# Design Doc: [Feature Name]

<!--
DETAIL LEVEL GUIDANCE:
- Focus on WHAT and WHY, not HOW (implementation details)
- Describe component responsibilities and interfaces, not code
- Use diagrams for architecture, not class definitions
- Keep language/framework agnostic where possible
- Target audience: developers who will implement this design
-->

## Overview

[Brief summary of what this design implements - 2-3 sentences max]

## Background

[Link to PRD: ../prd/feature_name.md]

[Any additional technical context needed - 1-2 paragraphs max]

## Goals

- [Technical goals derived from PRD]
- [Performance requirements]
- [Scalability requirements]

## Non-Goals

- [What this design explicitly does not cover]

## Design

### High-Level Architecture

[Architecture diagram or ASCII art]

[Description of major components and their interactions]

### Detailed Design

<!-- For each component: describe WHAT it does, not HOW it's coded -->

#### Component 1: [Name]

**Purpose**: [One sentence description]

**Responsibilities**:

- [What this component is responsible for]
- [Keep at conceptual level]

**Interface**:

- Input: [What data/requests it receives]
- Output: [What data/responses it provides]
- [No code signatures - describe conceptually]

### Data Model

<!-- Show logical data model, not physical implementation -->

[Describe entities, relationships, and key fields]
[Use simple diagrams or bullet points]
[Avoid SQL DDL unless critical to design]

### API Design

<!-- Describe API behavior and contracts, not exact schemas -->

[Describe operations available]
[Show example interactions]
[Focus on behavior, not exact field names]

### Error Handling

[How errors are handled and reported]

### Logging Strategy

Define which use cases apply to this feature and what specific data to log:

**Applicable Use Cases**:

- [ ] Tracing - Execution flow and timing
- [ ] Error Tracking - Problem identification and diagnosis
- [ ] User Behavior - UX improvement insights
- [ ] Performance - Resource usage and efficiency
- [ ] Security Audit - Compliance and audit trails
- [ ] Business Metrics - KPIs and analytics

**Implementation Details**:

```text
[For each selected use case, specify:]
- What specific events/data to log
- Required fields beyond standard (use_case, timestamp, trace_id)
- Retention period for this use case
- Any sampling strategy if high volume
```

**Privacy Considerations**:
[Identify any PII or sensitive data that needs sanitization]

## Security Considerations

[Authentication, authorization, data protection]

## Testing Strategy

- Unit tests: [Coverage goals and approach]
- Integration tests: [Key scenarios]
- Performance tests: [Benchmarks and targets]

## Migration Plan

[If applicable, how to migrate from current state]

## Alternatives Considered

### Alternative 1: [Name]

[Description]

[Why not chosen]

### Alternative 2: [Name]

[Description]

[Why not chosen]
