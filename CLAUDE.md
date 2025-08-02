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

#### Work Session Completion

**IMPORTANT**: When all work is completed in a session, proactively invoke the learner sub-agent to evaluate the entire conversation from start to finish. Pass the complete conversation history to the learner agent for analysis and knowledge extraction.

Example usage:
```
"All requested tasks have been completed. Let me now invoke the learner agent to analyze our work session."
→ Invoke learner agent with the entire conversation from work start to completion
```

### Improved Agent Descriptions for Claude Code

Based on official documentation best practices, use these descriptions for better automatic agent selection:

- **planner**: Use this agent at the START of any task or session to break down complex requirements, create implementation plans, or organize multi-step tasks. ALWAYS use when starting a new task. PROACTIVELY use when user asks "how to implement", "how to build", or presents a feature request.

- **builder**: Use this agent when you need to write code, create new features, implement functionality, or generate documentation. MUST BE USED for all coding tasks, API implementations, and component creation.

- **reviewer**: Use this agent PROACTIVELY after code changes to review quality, security, and best practices. Also use when user explicitly asks for code review, feedback, or quality assessment.

- **maintainer**: Use this agent when dealing with dependencies, technical debt, performance optimization, or system health. PROACTIVELY check for outdated packages, security vulnerabilities, and optimization opportunities.

- **learner**: Use this agent to extract patterns from completed work, document lessons learned, or create reusable knowledge. PROACTIVELY use after completing complex tasks to capture insights.

- **general-purpose**: Use this agent for complex research tasks, open-ended searches requiring multiple rounds of exploration, or when the task doesn't clearly fit other specialized roles. Best for investigative work where the scope is unclear.



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
