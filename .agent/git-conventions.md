# Git Conventions for AI Agents

## Pre-Commit Requirements

**MANDATORY before every commit:**
```bash
moon run check  # Or moon run <project>:check
```

This ensures:
- All tests pass
- Code is properly formatted
- Linting rules are satisfied
- Type checking passes (if applicable)

**Never commit without running `moon run check` first!**

## Commit Granularity

- Keep commits atomic: one logical change per commit
- Related changes should be committed together
- Unrelated changes should be in separate commits
- Each commit should represent a complete, working state
- Each commit should pass all tests and checks (lint, format, typecheck, etc.)

## Commit Order

- Consider logical dependencies between commits
- Infrastructure changes before features that use them
- Tool installations before tool configurations
- Configurations before code that depends on them
- Bug fixes that enable tests before the tests themselves
- Example order:
  1. Install tooling (e.g., linter)
  2. Add tool configuration
  3. Add tool integration (e.g., lint tasks)
  4. Fix issues found by the tool

## Commit Messages

AI agents must follow the Conventional Commits specification:

### Format

```
<type>[optional scope][!]: <description>

[optional body]

[optional footer(s)]
```

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

```
feat: add user authentication

fix(api): handle null values in response

feat!: change database schema

feat(lang): add Japanese localization

refactor(auth)!: replace JWT with session-based auth

Refs: #456
```
