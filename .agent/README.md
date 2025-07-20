# AI Agent Conventions

This directory contains conventions and rules for AI agents working on this project.

## IMPORTANT: When to Read These Files

AI agents MUST read the relevant files in this directory:

- **ALWAYS** at the start of a session:
  - Read `/README.md` for project overview
  - Read this entire `.agent/README.md` file
  - Read `project-conventions.md` for basic project rules
- **AS NEEDED** based on the task:
  - **BEFORE** making git commits - read `git-conventions.md`
  - **WHEN** designing new features - read `collaborative-design-process.md`
  - **WHEN** creating technical designs - read `technical-design-process.md` and `design-review-checklist.md`
  - **WHEN** installing tools - read `tool-installation-conventions.md`

## Directory Contents

- `git-conventions.md` - Conventional Commits format rules
- `project-conventions.md` - Language requirements for all files
- `collaborative-design-process.md` - Product design and PRD creation process
- `technical-design-process.md` - Engineering design and implementation process
- `design-review-checklist.md` - Guidelines for appropriate abstraction in design docs
- `tool-installation-conventions.md` - How to install tools (using mise)

## For AI Agent Developers

If you're integrating an AI agent with this project, ensure your agent:

1. Automatically discovers and reads the `.agent/` directory
2. Applies these conventions during its work
3. Re-reads files when they're updated
