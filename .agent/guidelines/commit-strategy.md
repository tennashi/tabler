# Commit Strategy and Granularity

This document defines the strategy for organizing and structuring commits.
For commit message format, see `commit-message.md`.

## Commit Granularity

- Keep commits atomic: one logical change per commit
- Related changes should be committed together
- Unrelated changes should be in separate commits
- Each commit should represent a complete, working state
- Each commit should pass all tests and checks (lint, format, typecheck, etc.)

## Commit Dependencies

- Consider logical dependencies between commits
- Infrastructure changes before features that use them
- Tool installations before tool configurations
- Configurations before code that depends on them
- Bug fixes that enable tests before the tests themselves
