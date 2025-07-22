# ADR-001: LLM Provider Selection

## Status

Proposed

## Context

The Smart Task Creation feature requires AI capabilities for:

- Natural language date and category extraction
- Task clarification dialogue
- Complex task decomposition suggestions
- Support for both Japanese and English languages

We need to start with a simple, working solution that can be enhanced later.

## Decision

Use **Claude Code subprocess invocation** as the LLM provider.

Implementation approach:

- Execute Claude Code CLI commands from the application
- Pass task input and receive structured responses
- Start simple, enhance later if needed

## Consequences

### Positive

- Simplest possible implementation
- No API costs during development
- Leverages existing Claude Code installation
- Can iterate quickly on prompts and behavior
- Easy to understand and debug

### Negative

- Requires Claude Code to be installed and accessible
- Not suitable for production deployment
- Subprocess overhead for each call

### Neutral

- Good enough for MVP and personal use
- Can be replaced with API integration later
- Defines interface that future implementations must follow

## Options Considered

### Option 1: Claude Code Subprocess (Selected)

- **Pros**: Simple, free, quick to implement
- **Cons**: Development-only, requires Claude Code
- **Evaluation**: Perfect for starting simple

### Option 2: Full API Integration

- **Pros**: Production-ready, scalable
- **Cons**: Complex, costs money, overkill for MVP
- **Evaluation**: Save for later phases

### Option 3: Hybrid Approach

- **Pros**: Flexible, future-proof
- **Cons**: Over-engineering for initial version
- **Evaluation**: Premature optimization
