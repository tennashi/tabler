# Claude Code Instructions

## Primary Reference

@.agent/README.md

All project instructions are defined in `.agent/` directory. This file only contains Claude Code-specific mappings.

## Critical Translation Rule for Claude Code

When you see instructions to "Execute the `X` task":

1. **ALWAYS interpret this as**: Launch the `X-er` subagent
2. **NEVER attempt to**: Read or execute task files directly
3. **The mapping is**: `X` task → `X-er` subagent (append "-er" to task name)

Example translations:

- "Execute the `branch-check` task" → Launch `branch-checker` subagent
- "Execute the `commit-checkpoint` task" → Launch `commit-checkpointer` subagent

## Claude Implementation Structure

### Subagents (`.claude/agents/`)

Each subagent wraps a corresponding task from `.agent/tasks/`:

- Subagent naming: task name + "-er" suffix
- Subagents provide Claude-specific execution context
- Always use subagents, never execute tasks directly

### Commands (`.claude/commands/`)

Slash commands wrap workflows from `.agent/workflows/`:

- Commands provide the Claude Code interface
- They reference the actual workflow definitions

## Workspace

Use `.agent/tmp/` for all temporary work (git-ignored)
