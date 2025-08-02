# Tool Installation Workflow for AI Agents

This document describes the step-by-step process for installing and managing tools using mise.

## Installing a New Tool (role: maintainer)

### Step 1: Check if tool exists in mise registry (role: maintainer)

````bash
# Search for the tool
mise registry | grep <toolname>

# Or check exact name
mise registry | grep "^<toolname>$"
```text

### Step 2: Install the tool (role: maintainer)

If found in standard registry:

```bash
mise use <tool>@latest
```text

If not found, try aqua backend:

```bash
mise use aqua:<org>/<tool>@latest
```text

### Step 3: Verify installation (role: maintainer)

```bash
# Verify the tool is available
mise exec -- <tool> --version

# Check it's in .mise.toml
cat .mise.toml
```text

### Step 4: Commit the changes (role: maintainer)

```bash
git add .mise.toml
git commit -m "build: add <tool> for <purpose>"
```text

## Common Tool Installation Patterns (role: maintainer)

### Installing a specific version (role: maintainer)

```bash
# Use specific version
mise use <tool>@<version>

# Examples:
mise use node@20.11.0
mise use go@1.21.5
```text

### Installing from aqua backend (role: maintainer)

```bash
# When tool is not in main registry
mise use aqua:golangci/golangci-lint@latest
mise use aqua:kubernetes-sigs/kind@latest
```text

### Installing multiple tools at once (role: maintainer)

```bash
# Install several tools
mise use node@latest python@latest go@latest

# All get added to .mise.toml in one operation
```text

## Running Commands with Tools (role: maintainer)

### Always use mise exec (role: maintainer)

```bash
# Correct way - works everywhere
mise exec -- <command>

# Examples:
mise exec -- node --version
mise exec -- npm install
mise exec -- go test ./...
```text

### Running complex commands (role: maintainer)

```bash
# Use quotes for complex commands
mise exec -- bash -c "npm install && npm test"

# Or use moon tasks which handle this automatically
moon run test
```text

## Troubleshooting (role: maintainer)

### Tool not found in registry (role: maintainer)

1. Check exact spelling: `mise registry | less`
2. Try aqua backend: `mise use aqua:<org>/<tool>@latest`
3. If still not found, document manual installation in README

### Version conflicts (role: maintainer)

```bash
# Check current versions
mise list

# Update to latest
mise use <tool>@latest

# Pin to specific version if needed
mise use <tool>@<specific-version>
```text

### PATH issues (role: maintainer)

```bash
# mise exec ensures correct PATH
mise exec -- which <tool>

# Don't rely on shell activation
# ❌ eval "$(mise activate)"
# ✅ mise exec -- <command>
```text

## References

- See `guidelines/tool-management.md` for tool selection rules
- See `workflows/commit.md` for committing tool changes
````
