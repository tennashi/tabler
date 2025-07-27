# Git Commit Workflow for AI Agents

This document describes the step-by-step process for making commits.

## Phase 1: Planning (role: planner)

### 1.1 Analyze Task and Determine Change Type

**ALWAYS start here** - Understand what needs to be done before any implementation:

#### Step 1: Parse the Request
First, clearly understand what the user is asking for:
- What specific functionality or fix is requested?
- Are there multiple sub-tasks that need separate commits?
- Any constraints or preferences mentioned?

#### Step 2: Investigate Current State (if modifying existing code)
Use these commands as needed based on your task:

```bash
# Find relevant files by extension
find . -type f -name "*.<ext>" | grep -v node_modules | grep -i <keyword>
# Common extensions: ts, js, tsx, jsx, py, go, rs, java, etc.

# Search for specific text in codebase
grep -r "<search-term>" --include="*.<ext>" --exclude-dir=node_modules

# Check for existing tests
find . -name "*.test.*" -o -name "*.spec.*" -o -name "*_test.*" | grep -i <feature>

# Understand project structure
tree -L 2 -d src/  # or use ls -la if tree is not available

# Check recent changes in the area
git log --oneline -10 -- <path-or-file>
```

**Note**: Adapt these commands to your project's language and structure.

#### Step 3: Determine Commit Type
Based on your understanding, identify the primary change type:

| If you're... | Use type | Example |
|-------------|----------|---------|
| Adding new user-facing functionality | `feat` | New API endpoint, UI component |
| Fixing broken functionality | `fix` | Correcting calculation errors, null handling |
| Restructuring code without changing behavior | `refactor` | Extract function, rename variables |
| Improving performance | `perf` | Optimize algorithm, add caching |
| Adding/updating tests only | `test` | New test cases, test fixes |
| Updating documentation only | `docs` | README updates, JSDoc comments |
| Code formatting only | `style` | Prettier, ESLint fixes |
| Changing build/dependencies | `build` | Webpack config, package updates |
| Updating CI/CD | `ci` | GitHub Actions, deployment scripts |

#### Step 4: Define Success Criteria
Before proceeding, be clear about:
- What files need to be created/modified
- What the end result should look like
- How you'll verify it works (tests, manual verification)
- Whether this could break existing functionality

This analysis ensures:
- Appropriate branch strategy (next section)
- Clear, accurate commit messages
- Focused, atomic changes
- Proper risk assessment

### 1.2 Determine Branch Strategy

Based on the task analysis and change type identified above, decide your branch strategy:

| Purpose | Branch Strategy |
|---------|----------------|
| New feature/task | Create new branch from main |
| Bug fix for production | Create hotfix branch from main |
| Continuing existing work | Use current feature branch |
| Experimentation | Create experimental branch |
| Documentation only | Can use current branch or create docs branch |

### 1.3 Check Current State

```bash
# Understand where you are
git branch --show-current
git status
```

### 1.4 Branch Decision Rules

Based on your purpose and current state, determine the appropriate branch strategy:

#### Key Rules:

- **If on main/master branch** → ALWAYS create a new branch (NO EXCEPTIONS)
- **New branches** → ALWAYS create from latest main
- **Never commit directly to main/master**

#### When to create a new branch:

- **On main branch** → Always
- **Starting a new task or feature**
- **Experimenting with approaches** that might not work out
- **Working on someone else's branch** without explicit permission
- **The current branch is a release or develop branch**

#### Work directly on current branch when:

- **Already on your own feature branch** for the SAME task
- **Continuing work** you just started in the same session
- **Making fixes to work** you just committed
- **Explicitly told to work on current branch**

## Phase 2: Branch Execution (role: maintainer)

### 2.1 Creating a New Branch

When the planning phase determines a new branch is needed:

```bash
# 1. Save any uncommitted work
git stash

# 2. Switch to main and update
git checkout main
git pull origin main

# 3. Create and switch to new branch from latest main
git checkout -b <descriptive-branch-name>

# 4. Restore your work if needed
git stash pop
```

### 2.2 Branch Naming Conventions

Choose names that reflect the purpose from Phase 1:

```bash
# Task/Issue based (preferred)
git checkout -b issue-123-user-authentication
git checkout -b task/add-payment-processing

# Type-based with description
git checkout -b feat/user-profile-page
git checkout -b fix/memory-leak-in-parser
git checkout -b refactor/payment-service

# Personal work branches
git checkout -b yourname/experiment-with-new-api
```

## Phase 3: Implementation (role: builder)

Execute the changes according to the purpose defined in Phase 1.

## Phase 4: Pre-Commit Quality Checks (role: maintainer)

**MANDATORY before every commit:**

1. **Run quality checks**

   Run the project's quality check command. This typically includes:
   - All tests pass
   - Code is properly formatted
   - Linting rules are satisfied
   - Type checking passes (if applicable)

   Check the project's README or build configuration for the specific command.

2. **Review changes**
   ```bash
   git status
   git diff
   ```

3. **Stage changes atomically**
   - Stage related changes together
   - Keep unrelated changes for separate commits

4. **Write commit message following conventions**
   - Use the purpose from Phase 1 to craft a clear message
   - See `guidelines/commit.md` for format rules

## Phase 5: Commit Execution (role: maintainer)

### Commit Order Strategy

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

## Handling Mixed Changes (role: maintainer)

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

## Common Scenarios (role: maintainer)

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

## Phase 6: Post-Commit Actions (role: maintainer)

### After committing changes on a feature branch:

1. **Ask about PR creation**
   - "Would you like to create a pull request now?"
   - If **no** → End workflow (changes remain local)
   - If **yes** → Continue to step 2

2. **Create PR using gh CLI**
   ```bash
   # Create PR with --head flag (automatically pushes)
   gh pr create \
     --head $(git branch --show-current) \
     --title "Brief description of changes" \
     --body "## What
   - Summary of changes

   ## Why  
   - Reason for changes

   ## Testing
   - How to test"
   ```

3. **Share PR link**
   - Copy and share the PR URL that gh returns
   - The PR is now ready for review

### Note about gh pr create

- Always use `--head $(git branch --show-current)` flag
- This automatically pushes unpushed commits
- In non-interactive mode, `--title` and `--body` are required
- No need to manually push before creating PR

## References

- See `guidelines/commit.md` for commit message format and conventions
- See `workflows/tdd.md` for test-driven development commit patterns