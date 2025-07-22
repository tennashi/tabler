# Git Commit Workflow for AI Agents

This document describes the step-by-step process for making commits.

## Pre-Commit Checklist

**MANDATORY before every commit:**

1. **Run quality checks**
   ```bash
   moon run check  # Or moon run <project>:check
   ```
   This ensures:
   - All tests pass
   - Code is properly formatted
   - Linting rules are satisfied
   - Type checking passes (if applicable)

2. **Review changes**
   ```bash
   git status
   git diff
   ```

3. **Stage changes atomically**
   - Stage related changes together
   - Keep unrelated changes for separate commits

4. **Write commit message following conventions**
   - See `guidelines/commit.md` for format rules

## Commit Order Strategy

When making multiple commits, follow this order:

1. **Infrastructure first**
   - Tool installations
   - Build configurations
   - CI/CD changes

2. **Dependencies next**
   - Library additions
   - Configuration files
   - Environment setup

3. **Implementation**
   - Core functionality
   - Tests
   - Documentation

4. **Fixes last**
   - Issues found by new tools
   - Test failures
   - Linting errors

### Example Sequence

```bash
# 1. Install new tool
mise use <tool>@latest
git add .mise.toml
git commit -m "build: add <tool> for <purpose>"

# 2. Add tool configuration
git add <tool-config-file>
git commit -m "build(<tool>): configure <specific settings>"

# 3. Integrate tool
git add moon.yml
git commit -m "build(moon): add <tool> task"

# 4. Fix issues found
git add <fixed-files>
git commit -m "fix: resolve <tool> warnings"
```

## Common Scenarios

### Single Feature Implementation

```bash
# 1. Check current state
git status

# 2. Run checks
moon run check

# 3. Stage all related changes
git add <feature-files>

# 4. Commit with descriptive message
git commit -m "feat: implement <feature>"
```

### Bug Fix with Tests

```bash
# 1. Fix the bug and add test
# 2. Run tests
moon run check

# 3. Commit test first (showing the bug exists)
git add <test-file>
git commit -m "test: add test for <bug description>"

# 4. Commit the fix
git add <fix-files>
git commit -m "fix: <bug description>"
```

### Refactoring

```bash
# 1. Make refactoring changes
# 2. Ensure tests still pass
moon run check

# 3. Commit with clear scope
git add <refactored-files>
git commit -m "refactor(<scope>): <what was refactored>"
```

## References

- See `guidelines/commit.md` for commit message format and conventions
- See `workflows/tdd.md` for test-driven development commit patterns