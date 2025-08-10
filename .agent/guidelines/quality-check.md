# Quality Check Guidelines

This document defines the mandatory quality checks that must pass before any commit.

## Core Requirement

**All commits MUST pass quality checks. No exceptions.**

## Primary Check Command

```bash
moon run :check
```

This command is mandatory. It typically includes:

- Format checking
- Linting
- Testing

## Handling Check Failures

### Mandatory Fix Policy

When checks fail:

1. **Identify the failure** - Read error messages carefully
2. **Fix all issues** - No commits with failing checks
3. **Re-run checks** - Ensure all checks pass

### Auto-fixable Issues

Some issues can be automatically fixed:

```bash
# Format issues
moon run :format

# Then re-run checks
moon run :check
```

## No-Skip Policy

Quality checks cannot be skipped for any reason:

- ❌ "It's just documentation" - Documentation has quality standards too
- ❌ "It's urgent" - Urgency doesn't justify broken code
- ❌ "It was already broken" - Fix it or don't commit
- ❌ "The test is flaky" - Fix the test or the code

## References

- Commit Strategy: `commit-strategy.md`
- Project Standards: `project.md`
