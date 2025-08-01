# Design Review Checklist for AI Agents

When creating or reviewing design documents, ensure appropriate abstraction levels:

## What Belongs Where (role: reviewer)

### Architecture Decision Records (ADRs) (role: reviewer)

Location: `/docs/adr/`

- Technology and library selections
- Significant technical decisions with rationale
- Trade-offs and alternatives considered
- Project-wide technical choices
- One decision per ADR

### Feature Design Documents (role: reviewer)

Location: `/docs/design/[feature-name]/`

- Feature-specific architecture and components
- Data models and interfaces
- Business logic and algorithms
- Implementation considerations (concepts, not specific libraries)
- Testing strategies

### Project Conventions (role: reviewer)

Location: `/.agent/`

- Development processes and workflows
- Coding standards and conventions
- Git commit conventions
- Language requirements

## Red Flags to Avoid (role: reviewer)

❌ **Don't put in feature design docs:**

- Specific library names (use concepts: "CLI framework" not "cobra")
- Project-wide directory structure
- Language-specific package names
- Development environment setup

✅ **Do put in feature design docs:**

- Component responsibilities and interactions
- Data flow and processing logic
- Interface contracts
- Performance and security considerations
- Implementation notes (conceptual level)

## Examples (role: reviewer)

**Too Specific (Avoid):**

````text
"Use github.com/spf13/cobra for CLI commands"
"Install with go get github.com/mattn/go-sqlite3"
```text

**Appropriate Abstraction:**

```text
"Use a CLI framework for command handling"
"SQLite driver with foreign key support enabled"
```text

## Review Questions (role: reviewer)

Before committing a design document, ask:

1. Could this design be implemented in a different language?
2. Are library choices documented in ADRs instead?
3. Is project structure documented separately?
4. Does it focus on WHAT and WHY, not HOW?

Remember: Design docs should survive technology changes!
````
