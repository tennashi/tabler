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

1. Determine branch name following `guidelines/branch-naming.md`
2. Create and switch to new branch:

   ```bash
   git checkout -b <type>/<descriptive-name>
   ```

### 4. Verify Ready to Work

```bash
git status
```

## Expected Result

- Working on appropriate branch for the task
- Not making direct changes to protected branches
