# Branch Naming Guidelines

## Naming Convention

### Format

```text
<type>/<descriptive-name>
```

### Types

Based on Conventional Commits types:

- `feat/` - New features
- `fix/` - Bug fixes
- `docs/` - Documentation only changes
- `style/` - Code style changes (formatting, missing semicolons, etc.)
- `refactor/` - Code refactoring without changing functionality
- `test/` - Adding or updating tests
- `build/` - Changes to build system or dependencies
- `ci/` - CI/CD configuration changes
- `perf/` - Performance improvements
- `chore/` - Other changes that don't modify src or test files

### Naming Rules

1. **Use kebab-case**: `feat/user-authentication` not `feat/userAuthentication`
2. **Be descriptive but concise**: `fix/login-error` not `fix/fix-the-bug-where-users-cant-login`
3. **Include issue number if applicable**: `fix/issue-123-login-error`
4. **No personal prefixes**: Avoid `john/feature` unless specifically required

### Examples

Good branch names:

- `feat/add-payment-processing`
- `fix/memory-leak-in-parser`
- `docs/update-api-reference`
- `refactor/extract-validation-logic`
- `test/add-integration-tests`
- `ci/add-github-actions`

## References

- Follows Conventional Commits types for consistency
- See `guidelines/branch-strategy.md` for when to create branches
