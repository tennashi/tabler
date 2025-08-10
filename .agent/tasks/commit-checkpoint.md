# Commit Checkpoint Task

Execute this task when you need to commit current changes.

## Execution Steps

### 1. Confirm Commit Intent

First, ask the user:

> "Would you like to commit the current changes?"

If no → Exit task

### 2. Assess Current Changes

```bash
# Check what has been modified
git status

# Review all changes
git diff
```

### 3. Run Quality Checks

Execute quality checks according to `../guidelines/quality-check.md`.

If checks fail:

1. Fix all issues
2. Re-run checks until they pass

### 4. Analyze and Group Changes

Analyze all changes and identify logical groups:

Example groupings:

- Feature implementation + related tests
- Bug fix + regression test
- Documentation updates
- Refactoring changes
- Configuration changes

Present the analysis to the user:

> "I've identified the following changes:
>
> - Group A: [description of changes]
> - Group B: [description of changes]
>
> Which would you like to commit? (A/B/all)"

### 5. Stage Selected Files

Based on user's choice:

```bash
# Stage selected files
git add <files-for-selected-group>
```

### 6. Create Commit

Generate commit message following `../guidelines/commit-message.md`:

```bash
# Create commit with generated message
git commit -m "<type>(<scope>): <description>"
```

Message generation rules:

- Analyze staged changes to determine type (feat/fix/docs/refactor/test/chore)
- Identify scope from modified files
- Write clear, concise description

### 7. Verify Commit

```bash
# Verify the commit was created correctly
git log -1 --stat

# Check remaining unstaged changes
git status
```

## Important Notes

- **Quality checks are mandatory** - Never skip them
- **User confirmation is required** - Both for starting and for selecting what to commit
- **Commits are local only** - This task does not push to remote
- **Follow guidelines** - Always reference the relevant guidelines for standards

## References

- Quality Check: `../guidelines/quality-check.md`
- Commit Message Format: `../guidelines/commit-message.md`
- Commit Strategy: `../guidelines/commit-strategy.md`
- Branch Strategy: `../guidelines/branch-strategy.md`
