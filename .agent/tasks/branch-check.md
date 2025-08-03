# Branch Check Task

## Purpose

Ensure working on appropriate branch before file operations

## When to Execute

- Before file creation
- Before file editing
- Before file deletion

## Steps

### 1. Check Current Branch

```bash
git branch --show-current
```

### 2. Apply Branch Strategy

Use decision criteria from `guidelines/branch-strategy.md` to determine if a new branch is needed.

### 3. If New Branch Required

1. Check for uncommitted changes:
   ```bash
   git status --porcelain
   ```

2. If changes exist, stash them:
   ```bash
   git stash push -m "WIP: moving to new branch"
   ```

3. Switch to main and update:
   ```bash
   git checkout main
   git pull origin main
   ```

4. Create new branch following `guidelines/branch-naming.md`:
   ```bash
   git checkout -b <type>/<descriptive-name>
   ```

5. If changes were stashed, restore them:
   ```bash
   git stash pop
   ```

### 4. Verify Ready to Work

```bash
git status
```

## Expected Result

- Working on appropriate branch for the task
- Not making direct changes to protected branches
