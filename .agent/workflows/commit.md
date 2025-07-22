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

## Handling Mixed Changes

### When files contain multiple unrelated changes

Since AI agents cannot use interactive staging (`git add -p`), use the stash-based approach:

#### Step 1: Stage files with single-purpose changes

```bash
# First, stage files that contain only changes for current purpose
git add <file-with-single-purpose>
```

#### Step 2: Handle mixed files using stash

For files containing mixed changes:

```bash
# 1. Temporarily save ALL changes
git stash

# 2. Restore only the stashed file you need
git checkout stash@{0} -- <mixed-file>

# 3. Manually edit the file to keep only relevant changes
# Remove unrelated changes from the file

# 4. Stage the cleaned file
git add <mixed-file>

# 5. Restore remaining unstaged changes
git stash pop
# Resolve any conflicts if they occur
```

### Verifying staged changes work independently

**Important**: Ensure staged changes can stand alone:

```bash
# 1. Stash unstaged changes (keep staged changes)
git stash --keep-index

# 2. Run tests on staged changes only
# Use the project's test/check command
# Build should pass with only staged changes

# 3. If tests pass, commit
git commit -m "type: description"

# 4. Restore unstaged changes
git stash pop
```

### Example: Separating formatting from features

When formatting and feature changes are mixed:

```bash
# 1. Stash all changes
git stash

# 2. Apply and edit for formatting only
git stash show -p > all-changes.patch
# Manually edit all-changes.patch to keep only formatting
git apply all-changes.patch
git add <files>
git commit -m "style: format code"

# 3. Restore and handle feature changes
git stash pop
git add <feature-files>
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