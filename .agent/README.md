# AI Agent Conventions

This directory contains conventions and rules for AI agents working on this project.

## IMPORTANT: When to Read These Files

AI agents MUST read the relevant files in this directory:

- **ALWAYS** at the start of a session:
  - Read `/README.md` for project overview
  - Read this entire `.agent/README.md` file
  - Read `guidelines/project.md` for basic project rules
- **AS NEEDED** based on the task:
  - **BEFORE** making git commits - read `guidelines/commit.md` and `workflows/commit.md`
  - **WHEN** designing new features - read `workflows/collaborative-design.md`
  - **WHEN** creating technical designs - read `workflows/technical-design.md` and `workflows/design-review-checklist.md`
  - **WHEN** installing tools - read `guidelines/tool-management.md` and `workflows/tool-installation.md`
  - **WHEN** implementing logging - read `guidelines/logging.md`
  - **WHEN** implementing features using TDD - read `workflows/tdd.md`

## Directory Structure

### Guidelines (guidelines/)
**Conventions and rules** - What to do and why

- `commit.md` - Git commit conventions (Conventional Commits)
- `project.md` - Language requirements for all files
- `logging.md` - Use case-based logging guidelines
- `tool-management.md` - How to install and manage tools (using mise)

### Workflows (workflows/)
**Processes and procedures** - How to execute tasks

- `tdd.md` - Test-Driven Development process for AI agents
- `commit.md` - Step-by-step git commit procedures
- `tool-installation.md` - Installing and managing tools with mise
- `collaborative-design.md` - Product design and PRD creation process
- `technical-design.md` - Engineering design and implementation process
- `design-review-checklist.md` - Guidelines for appropriate abstraction in design docs

### Temporary Files (tmp/)
**Working directory** - For AI agent temporary work

- Git-ignored directory for temporary files
- Use for drafts, debugging files, session notes
- Files persist between sessions but are not tracked by git

## For AI Agent Developers

If you're integrating an AI agent with this project, ensure your agent:

1. Automatically discovers and reads the `.agent/` directory
2. Applies these conventions during its work
3. Re-reads files when they're updated
