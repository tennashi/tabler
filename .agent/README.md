# AI Agent Conventions

This directory contains conventions and rules for AI agents working on this project.

## IMPORTANT: When to Read These Files

AI agents MUST read the relevant files in this directory:

- **ALWAYS** at the start of a session - read all files
- **BEFORE** making git commits - read `git-conventions.md`
- **BEFORE** creating or modifying files - read `project-conventions.md`

## Directory Contents

- `git-conventions.md` - Conventional Commits format rules
- `project-conventions.md` - Language requirements for all files

## For AI Agent Developers

If you're integrating an AI agent with this project, ensure your agent:
1. Automatically discovers and reads the `.agent/` directory
2. Applies these conventions during its work
3. Re-reads files when they're updated