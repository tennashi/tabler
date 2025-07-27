# Tool Installation Workflow for AI Agents

This document describes the step-by-step process for installing and managing tools using mise.

## Installing a New Tool

### Step 1: Check if tool exists in mise registry

````bash
# Search for the tool
mise registry | grep <toolname>

# Or check exact name
mise registry | grep "^<toolname>$"
```text

### Step 2: Install the tool

If found in standard registry:

```bash
mise use <tool>@latest
```text

If not found, try aqua backend:

```bash
mise use aqua:<org>/<tool>@latest
```text

### Step 3: Verify installation

```bash
# Verify the tool is available
mise exec -- <tool> --version

# Check it's in .mise.toml
cat .mise.toml
```text

### Step 4: Commit the changes

```bash
git add .mise.toml
git commit -m "build: add <tool> for <purpose>"
```text

## Common Tool Installation Patterns

### Installing a specific version

```bash
# Use specific version
mise use <tool>@<version>

# Examples:
mise use node@20.11.0
mise use go@1.21.5
```text

### Installing from aqua backend

```bash
# When tool is not in main registry
mise use aqua:golangci/golangci-lint@latest
mise use aqua:kubernetes-sigs/kind@latest
```text

### Installing multiple tools at once

```bash
# Install several tools
mise use node@latest python@latest go@latest

# All get added to .mise.toml in one operation
```text

## Running Commands with Tools

### Always use mise exec

```bash
# Correct way - works everywhere
mise exec -- <command>

# Examples:
mise exec -- node --version
mise exec -- npm install
mise exec -- go test ./...
```text

### Running complex commands

```bash
# Use quotes for complex commands
mise exec -- bash -c "npm install && npm test"

# Or use moon tasks which handle this automatically
moon run test
```text

## Troubleshooting

### Tool not found in registry

1. Check exact spelling: `mise registry | less`
2. Try aqua backend: `mise use aqua:<org>/<tool>@latest`
3. If still not found, document manual installation in README

### Version conflicts

```bash
# Check current versions
mise list

# Update to latest
mise use <tool>@latest

# Pin to specific version if needed
mise use <tool>@<specific-version>
```text

### PATH issues

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
