# Project Instructions for Claude Code

@.agent/README.md

## Important Note

The main instructions and processes are defined in `.agent/` directory. This file (CLAUDE.md) and `.claude/`
directory primarily contain references to `.agent/` and Claude-specific implementations only.

## Claude-specific Files

### Sub Agent Selection Guidelines

**IMPORTANT**: When the user specifies a role (e.g., "as [role]", "role: [role]"), interpret this as an instruction to invoke the corresponding sub agent.

#### Role-based Agent Invocation

When users use role-specifying phrases:
- "as planner" → Invoke planner agent
- "role: builder" → Invoke builder agent  
- "as reviewer" → Invoke reviewer agent
- "as maintainer" → Invoke maintainer agent
- "as learner" → Invoke learner agent



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
