# AI Agent Conventions

Central hub for AI agent conventions, guidelines, and workflows for this project.

## Quick Start

AI agents working on this project should:

1. **Start here** - Read this README for overview
2. **Check subdirectories** - Each directory has its own README with detailed information
3. **Follow the rules** - Apply conventions consistently throughout work

## Directory Structure

```plaintext
.agent/
├── guidelines/     # Conventions and rules (what to do and why)
├── tasks/         # Executable tasks (specific actions to perform)
├── workflows/     # Multi-step processes (how to accomplish goals)
└── tmp/          # Git-ignored workspace for temporary files
```

## When to Use Each Directory

### 📋 Tasks (`tasks/`)

**Required checkpoints** in your workflow

- Execute `branch-check.md` BEFORE any file operations
- Execute `commit-checkpoint.md` after significant changes

### 📐 Guidelines (`guidelines/`)

**Standards to follow** throughout development

- Project conventions, commit standards, branching strategies
- Tool management and logging practices

### 🔄 Workflows (`workflows/`)

**Step-by-step processes** for complex operations

- TDD implementation, design documentation
- Commit procedures, tool installation

### 💾 Temporary (`tmp/`)

**Working space** for ephemeral content

- Session notes, drafts, debugging files
- Git-ignored, persistent between sessions

## Essential Reading Order

1. **This file** (`.agent/README.md`) - Overview
2. **Subdirectory READMEs** - Detailed information for each area
3. **`guidelines/project.md`** - Core project requirements
4. **Task-specific files** - As needed based on current work

## For AI Agent Developers

When integrating an AI agent with this project:

1. Configure automatic discovery of `.agent/` directory
2. Read relevant files based on task context
3. Apply conventions consistently
4. Re-read files when they're updated
