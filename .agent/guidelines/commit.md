# Git Conventions for AI Agents

This document defines the conventions and rules for git commits. For step-by-step commit workflows, see `workflows/commit.md`.

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

For specific ordering strategies and examples, see `workflows/commit.md`

## Commit Messages

AI agents must follow the Conventional Commits specification:

### Format

````text
<type>[optional scope][!]: <description>

[optional body]

[optional footer(s)]
```text

### Types

- `feat`: New feature (correlates with MINOR in SemVer)
- `fix`: Bug fix (correlates with PATCH in SemVer)
- `docs`: Documentation only changes
- `style`: Changes that don't affect code meaning (formatting, semicolons, etc.)
- `refactor`: Code changes that neither fix bugs nor add features
- `test`: Adding or modifying tests
- `chore`: Changes to build process or auxiliary tools
- `perf`: Performance improvements
- `ci`: Changes to CI configuration
- `build`: Changes affecting build system or dependencies

### Additional Rules

1. **Scope** (optional): Provide additional context in parentheses
   - Example: `feat(auth): add OAuth2 support`

2. **Breaking Changes**: Use `!` after type/scope
   - Example: `feat!: change API response format`
   - Example with scope: `feat(api)!: change response structure`

3. **Body** (optional): Detailed explanation after blank line
   - Use when the description alone isn't sufficient

4. **Footer** (optional): References and metadata
   - Issue references: `Refs: #123`

### Examples

```text
feat: add user authentication

fix(api): handle null values in response

feat!: change database schema

feat(lang): add Japanese localization

refactor(auth)!: replace JWT with session-based auth

Refs: #456
```text
````
