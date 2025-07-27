# Monorepo Management (moon)

All build, test, and quality check tasks must be executed through **moon**.

## Mandatory Rules

1. **No direct commands** - Never use `npm run`, `go test`, etc.
2. **Use moon run** - All tasks must be executed via moon
3. **Verify tasks** - Check available tasks with `moon query tasks`

## Task Execution

### Basic Commands

```bash
# Run a task
moon run <task-name>

# Run task for specific project
moon run <project>:<task-name>

# Run task for all projects
moon run :<task-name>
```

### Task Discovery

```bash
# List all available tasks
moon query tasks

# List tasks for specific project
moon query tasks --project <project-name>

# Show task details
moon query tasks <task-name> --json
```

## Task Definition Conventions

### Task Definition Locations

1. **Global tasks**: `.moon/tasks.yml`
2. **Language-specific tasks**: `.moon/tasks/<language>.yml` (e.g., `go.yml`, `typescript.yml`)
3. **Tag-based tasks**: `.moon/tasks/tag-<tag>.yml`

### Project-Level Task Customization

Projects can customize workspace tasks in their `moon.yml`:

```yaml
tasks:
  # Extend inherited task
  test:
    args: "--race" # For Go: add race condition detection
    options:
      mergeArgs: "append" # Add to existing args

  # Add environment variables
  build:
    env:
      CGO_ENABLED: "0"
    options:
      mergeEnv: "append"
```

### Merge Strategies

- `append`: Add local values after global (default)
- `prepend`: Add local values before global
- `replace`: Replace entirely
- `preserve`: Keep global values

### Best Practices

- Define reusable tasks at workspace level
- Inherit and customize for project-specific needs
- Document task purpose in descriptions
- Use consistent naming conventions

## Cache Management

```bash
# Clear all caches
moon clean

# Clear specific project cache
moon clean <project>

# Run without cache
moon run <task> --no-cache
```

## Troubleshooting

### Common Issues

1. **Task not found**
   ```bash
   moon query tasks | grep <partial-name>
   ```

2. **Dependency conflicts**
   ```bash
   moon query projects --affected
   ```

3. **Cache issues**
   ```bash
   moon clean && moon run <task>
   ```

## Prohibited Practices

- Direct tool execution (`go test`, `npm run`, etc.)
- Building/testing without moon
- Running tasks without verification
- Manually editing moon cache files
