# AI Agent Working Instructions

This is your primary instruction manual for working on this project.

## Getting Started - Do This First

After reading this document, you MUST immediately:

1. Read this file completely to understand how to work
2. Check the README in each subdirectory for detailed information
3. Read `guidelines/project.md` for project-specific rules and standards
4. Apply the relevant guidelines to your current work
5. Execute required tasks at their designated checkpoints

## How You Should Work

### File Operation Protocol

When modifying ANY files in this project:

1. **First** - Execute the `branch-check` task to ensure you're on the correct branch
2. **Then** - Make your file changes (create, edit, or delete)
3. **Finally** - Execute the `commit-checkpoint` task to create proper commits

This sequence is mandatory for every file operation.

### PRD Creation Protocol

When creating or updating Product Requirements Documents:

1. **Create/Update** - Write or modify the PRD in `docs/prd/`
2. **Review** - Execute the `prd-review` task to ensure Why/What focus
3. **Fix** - Address any implementation details (How) found in the review
4. **Verify** - Re-run `prd-review` until no issues are found

This ensures PRDs maintain proper focus on problems and user needs, not implementation.

### Understanding the Structure

Each directory serves a specific role in guiding your work:

- **`tasks/`** - Single actions you must execute at specific checkpoints
- **`guidelines/`** - Rules and standards that govern how you work
- **`workflows/`** - Step-by-step procedures for completing complex operations
- **`tmp/`** - Your personal workspace for temporary files (git-ignored)

### When to Use What

- **Execute tasks** at required checkpoints:
  - `branch-check` before any file operations
  - `prd-review` after creating/updating PRDs
  - `commit-checkpoint` after significant changes
- **Follow guidelines** continuously throughout your work
- **Use workflows** when tackling multi-step operations
- **Work in tmp/** for any temporary or experimental content

## Directory Reference

```plaintext
.agent/
├── guidelines/     # What rules to follow
├── tasks/         # What to execute and when
├── workflows/     # How to complete complex operations
└── tmp/          # Where to put temporary work
```
