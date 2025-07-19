# Git Conventions for AI Agents

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