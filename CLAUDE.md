# Project Instructions for Claude Code

@.agent/README.md

## Important Note

The main instructions and processes are defined in `.agent/` directory. This file (CLAUDE.md) and `.claude/`
directory primarily contain references to `.agent/` and Claude-specific implementations only.

## Claude-specific Files

### Sub Agent Selection Guidelines

**IMPORTANT**: When receiving any task, always delegate it to an appropriate sub agent.

#### Agent Selection Criteria

Each sub agent has specific use cases defined in `.claude/agents/`:
- Review the `description` field in each agent file for selection criteria
- These descriptions specify when to use each agent and include proactive triggers
- Always check `.claude/agents/` for the most up-to-date agent capabilities


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
