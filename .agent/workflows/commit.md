# Git Commit Workflow for AI Agents

This document describes the step-by-step process for making commits.

## Phase 1: Planning (role: planner)

### 1.1 Task Analysis & Commit Planning

1. **What needs to be done?**
   - Parse user request → identify subtasks
   - Check existing code if modifying (use grep/find as needed)
   - Determine change type from this table:

   | If you're... | Type | Example |
   |-------------|------|---------|
   | Adding new feature | `feat` | feat(auth): add OAuth2 login |
   | Fixing a bug | `fix` | fix(api): handle null response |
   | Refactoring | `refactor` | refactor(utils): extract validation |
   | Other | `docs`/`test`/`style`/`build`/`ci` | docs(readme): update API section |

2. **Plan commit(s) and branch name:**
   
   **First, check commit message guidelines:**
   - See `guidelines/commit.md` for format and conventions
   - Follow Conventional Commits format: `<type>(<scope>): <subject>`
   
   ```
   Branch: feat/date-filter
   
   Single commit: feat(search): add date filter
   
   OR
   
   Branch: feat/oauth-login
   Multiple commits:
   1. build(deps): add OAuth2 library
   2. feat(auth): implement OAuth2 provider
   3. test(auth): add OAuth2 tests
   ```

### 1.2 Branch Strategy

Based on your task:

| Task | Strategy |
|------|----------|
| New feature/fix | Create branch from main |
| Continue existing work | Stay on current branch |
| Experiment | Create experimental branch |

**Check current state:**
```bash
git branch --show-current
git status
```

**Rules:**
- On main? → MUST create new branch
- New branch? → ALWAYS from latest main
- NEVER commit to main directly

## Phase 2: Branch Execution (role: maintainer)

### 2.1 Creating a New Branch

When planning determines a new branch is needed:

```bash
# 1. Save any uncommitted work
git stash

# 2. Switch to main and update
git checkout main
git pull origin main

# 3. Create and switch to new branch from latest main
git checkout -b <branch-name-from-phase-1>

# 4. Restore your work if needed
git stash pop
```


## Phase 3: Commit Execution (role: maintainer)

**Loop through each commit planned in Phase 1:**

### For each planned commit:

#### Step 1: Stage files for current commit

1. **First, stage whole files that match the commit purpose**
   ```bash
   git status
   git diff
   
   # Stage files where ALL changes belong to current commit
   git add <file1> <file2>  # Files fully dedicated to this commit
   ```

2. **Then, handle mixed files using patch method**
   ```bash
   # If files contain mixed changes:
   git stash --keep-index  # Keep already staged files
   git stash show -p stash@{0} > changes.patch
   # Edit changes.patch to keep only parts for current commit
   git apply changes.patch
   git add <mixed-files>
   git stash pop  # May cause conflicts - resolve if needed
   ```

#### Step 2: Quality check staged changes
```bash
# Isolate staged changes
git stash --keep-index

# Run tests/linting/formatting on staged changes only
# (Use project-specific commands)

# If checks fail:
git reset          # Unstage
git stash pop      # Restore all changes
# Fix issues and return to Step 1
```

#### Step 3: Create commit
```bash
# If checks pass, commit with planned message
git commit -m "<message from Phase 1>"

# Restore unstaged changes for next commit
git stash pop
```

#### Step 4: Continue to next planned commit
Repeat Steps 1-3 for each commit in your Phase 1 plan.

### Commit Order Strategy

When making multiple commits, follow this order:

1. **Infrastructure first** (tools, build configs, CI/CD)
2. **Dependencies next** (libraries, configs)
3. **Implementation** (features, tests, docs)
4. **Fixes last** (linting, test failures)

## Phase 4: Post-Commit Actions (role: maintainer)

### Creating a Pull Request

1. **Ask about PR creation**
   - "Would you like to create a pull request now?"
   - If no → End workflow

2. **If yes, push changes to remote**
   ```bash
   git push -u origin $(git branch --show-current)
   ```

3. **Create PR using gh CLI**
   ```bash
   gh pr create \
     --title "<summarize the changes>" \
     --body "## Changes
   <list commits from Phase 1>
   
   ## Why
   <main reason for these changes>
   
   ## Testing
   <how to verify these changes work>"
   ```

4. **Share PR link**
   - Copy the URL that gh returns
   - The PR is now ready for review

## References

- See `guidelines/commit.md` for commit message format and conventions
- See `workflows/tdd.md` for test-driven development commit patterns