# Branch Strategy Guidelines

## When to Create a New Branch

### ALWAYS Create a New Branch When:

1. **On main/master branch**
   - Never make direct changes to main
   - Always create a feature branch first

2. **Starting new work**
   - New feature implementation
   - Bug fix
   - Documentation update
   - Any code changes

3. **Work doesn't match current branch**
   - Current branch is `fix/login-error` but you need to add a new feature
   - Current branch is `docs/readme` but you need to fix a bug

### Continue on Current Branch When:

1. **Work directly relates to branch purpose**
   - On `feat/user-auth` and adding more auth functionality
   - On `fix/memory-leak` and fixing related issues

2. **Making related changes**
   - Fixing typos in code you just wrote
   - Adding tests for feature you just implemented
   - Updating docs for your current changes

## Branch Creation Decision Flow

```
Current Branch?
├── main/master → ALWAYS create new branch
└── feature branch
    └── Does work match branch purpose?
        ├── Yes → Continue on current branch
        └── No → Create new branch
```

## Protected Branches

Never commit directly to:
- `main` (or `master`)
- `develop` (if exists)
- `release/*` branches
- Any branch marked as protected in repository settings

## Branch Lifecycle

1. **Create**: Branch from latest main
2. **Work**: Make focused changes
3. **Merge**: Via pull request
4. **Delete**: After successful merge

## References
- See `guidelines/branch-naming.md` for naming conventions
- See `workflows/commit.md` for commit workflow