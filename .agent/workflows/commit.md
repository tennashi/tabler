# Git Commit Workflow for AI Agents

This document describes the step-by-step process for making commits.

## Pre-Commit Checklist

**MANDATORY before every commit:**

1. **Clarify the purpose of changes**
   Before staging any files, write down:
   - What problem does this change solve?
   - Why is this change necessary?
   - What is the expected outcome?
   
   This helps ensure commit messages accurately reflect the intent.

2. **Run quality checks**
   
   Run the project's quality check command. This typically includes:
   - All tests pass
   - Code is properly formatted
   - Linting rules are satisfied
   - Type checking passes (if applicable)
   
   Check the project's README or build configuration for the specific command.

3. **Review changes**
   ```bash
   git status
   git diff
   ```

4. **Stage changes atomically**
   - Stage related changes together
   - Keep unrelated changes for separate commits

5. **Write commit message following conventions**
   - Use the purpose from step 1 to craft a clear message
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

# 3. Integrate tool with build system
git add <build-config-file>
git commit -m "build: integrate <tool> into build process"

# 4. Fix issues found
git add <fixed-files>
git commit -m "fix: resolve <tool> warnings"
```

## Partial Staging Techniques

### When files contain mixed changes

Sometimes a file contains both related and unrelated changes. Use partial staging:

```bash
# Interactive staging
git add -p <file>

# For each hunk, choose:
# y - stage this hunk
# n - do not stage this hunk
# s - split into smaller hunks
# e - manually edit the hunk
```

### Verifying staged changes only

To ensure staged changes work independently:

```bash
# 1. Stash unstaged changes
git stash --keep-index

# 2. Run tests on staged changes only
# Use the project's test/check command

# 3. If tests pass, commit
git commit -m "type: description"

# 4. Restore unstaged changes
git stash pop
```

### Mixed file example

When a file has formatting fixes mixed with feature changes:

```bash
# 1. Stage only formatting changes
git add -p file.go
# Select only formatting hunks

# 2. Commit formatting separately
git commit -m "style: format file.go"

# 3. Stage feature changes
git add -p file.go
# Select feature hunks

# 4. Commit feature
git commit -m "feat: add new functionality"
```

## Common Scenarios

### Single Feature Implementation

```bash
# 1. Check current state
git status

# 2. Run project's quality checks
# (tests, lint, format, etc.)

# 3. Stage all related changes
git add <feature-files>

# 4. Commit with descriptive message
git commit -m "feat: implement <feature>"
```

### Bug Fix with Tests

```bash
# 1. Fix the bug and add test
# 2. Run project's test suite

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
# Run project's test suite

# 3. Commit with clear scope
git add <refactored-files>
git commit -m "refactor(<scope>): <what was refactored>"
```

## References

- See `guidelines/commit.md` for commit message format and conventions
- See `workflows/tdd.md` for test-driven development commit patterns