# Tool Installation Conventions for AI Agents

This project uses **mise** for tool version management. All tools should be installed through mise.

## Decision Guidelines - Always Choose These

### Installing Tools
**Always use**: `mise use <tool>@latest`
- This installs AND adds to `.mise.toml` in one step
- Don't use `mise install` followed by manual config editing
- Don't use `mise local` (it's an alias for `use`)

### Running Commands
**Always use**: `mise exec -- <command>`
- Consistent across all environments
- Works in scripts and CI
- Don't rely on shell activation or PATH modifications

### Backend Selection
**Check in this order**:
1. Standard mise registry: `mise use <tool>@latest`
2. If not found, try aqua: `mise use aqua:<org>/<tool>@latest`
3. If still not found, use official installer and document in README

### Configuration
**Always commit**: `.mise.toml` file
- This ensures reproducible environments
- Don't use global configurations for project tools

## Common Patterns

### Adding a new tool
```bash
# First, check if it exists
mise registry | grep <toolname>

# If found:
mise use <tool>@latest

# If not found, try aqua:
mise use aqua:<org>/<tool>@latest

# Always verify:
mise exec -- <tool> --version
```

### Don't Do These
- ❌ Don't edit `.mise.toml` manually then run `mise install`
- ❌ Don't use `eval "$(mise activate)"` in scripts
- ❌ Don't assume tools are in PATH
- ❌ Don't mix `mise global` with project tools
- ❌ Don't use different installation methods for the same tool

Remember: Consistency over flexibility!