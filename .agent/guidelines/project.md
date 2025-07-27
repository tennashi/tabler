# Project Conventions for AI Agents

## Language

- All code, documentation, commit messages, and files must be written in English

## Build System

This project uses **moon** as the build system and task runner.

### Running Tasks

**Always use moon for running tasks:**

````bash
# List all available tasks
moon query tasks

# Common task (usually available)
moon run check    # Runs all quality checks before commit

# Run specific tasks
moon run <task-name>
```text

**Important**:

- Always check available tasks with `moon query tasks` before running
- If `check` task exists, run it before committing
- Use `moon run <task-name>` to execute tasks, not direct tool invocations

**Do NOT use:**

- Direct tool invocations (e.g., `go test`, `npm test`)
- Make or other build systems
- Shell scripts for common tasks

### Task Dependencies

Moon handles task dependencies automatically. When you run a task, moon will:

- Execute dependent tasks first
- Cache results for efficiency
- Run tasks in parallel when possible
````
