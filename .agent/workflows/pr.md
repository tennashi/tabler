# PR Merge Automation Workflow

This workflow guides the AI agent through creating a PR, waiting for CI, and merging.

## Step 1: Analyze Changes

Execute:

```bash
git diff --stat main...HEAD
git log --oneline main...HEAD
```

Then:

- Extract type from branch name (e.g., `feat/` → `feat:`)
- Write meaningful title based on actual changes
- Fill PR body following `.github/pull_request_template.md`

## Step 2: Create PR

Execute:

```bash
git push -u origin $(git branch --show-current)
gh pr create --title "YOUR_GENERATED_TITLE" --body "YOUR_GENERATED_BODY"
```

## Step 3: Wait for CI

Execute:

```bash
gh pr checks --watch --fail-fast
```

If CI fails, try to fix automatically:

- Run linter/formatter if available in the project
- Commit any fixes and push
- Return to Step 3 to retry

If CI passes, continue to Step 4.

## Step 4: Merge PR

Execute:

```bash
gh pr merge --merge --delete-branch
```

## Step 5: Return to Main

Execute:

```bash
git checkout main
git pull origin main
```
