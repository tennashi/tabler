# Contributing to Tabler

Thank you for your interest in contributing to Tabler! This guide will help you get started with development.

## Development Setup

### Prerequisites

We use [mise](https://mise.jdx.dev/) for tool management. Install it first:

```bash
# macOS
brew install mise

# or see https://mise.jdx.dev/getting-started.html for other platforms
```

### Getting Started

```bash
# Clone the repository
git clone https://github.com/tennashi/tabler.git
cd tabler

# Install all required tools
mise install

# Check that everything is working
moon check
```

## Development Workflow

### Before You Start

Read the relevant documentation:
- Product requirements in [`docs/prd/`](docs/prd/)
- Design documents in [`docs/design/`](docs/design/)
- Architecture decisions in [`docs/adr/`](docs/adr/)

### Making Changes

1. Create a feature branch
2. Make your changes
3. Run checks before committing:
   ```bash
   moon check  # Runs all linting and formatting
   ```
4. Write clear commit messages
5. Submit a pull request

### Code Style

- We use `dprint` for code formatting
- We use `markdownlint` for markdown files
- All code and documentation must be in English

## AI-Assisted Development

If you're using AI tools (like Claude Code) for development:
- AI agents should follow conventions in [`.agent/`](.agent/)
- These conventions are specifically for AI tools, not human developers

## Project Structure

- `docs/` - Project documentation
- `.agent/` - AI agent conventions (for AI tools only)
- `.moon/` - Build tool configuration
- `.claude/` - Claude-specific configurations

## Getting Help

- Check existing issues and PRs
- Read the documentation in `docs/`
- Ask questions in issues

## License

By contributing, you agree that your contributions will be licensed under the same license as the project.