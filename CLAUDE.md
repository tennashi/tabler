# Project Instructions for Claude Code

@.agent/README.md

## Important Note

The main instructions and processes are defined in `.agent/` directory. This file (CLAUDE.md) and `.claude/`
directory primarily contain references to `.agent/` and Claude-specific implementations only.

## Claude-specific Files

### Subagents

Location: `/.claude/agents/`

- Execute tasks defined in `.agent/tasks/`
- Before file operations: branch-checker subagent

### Slash Commands

Location: `/.claude/commands/`

- These files reference the actual processes defined in `.agent/`
- Contain only Claude-specific command handling logic

### Temporary Workspace

Location: `/.agent/tmp/`

- Git-ignored directory for Claude's temporary work
- Use freely for:
  - Session notes and context
  - Work-in-progress analysis
  - Temporary calculations or data
  - Any files you need during development
- Everything in this directory is ephemeral and not tracked by git
