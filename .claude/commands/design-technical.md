# Technical Design Command

Follow the technical design process defined in `.agent/workflows/technical-design.md` to create a technical design based on a PRD.

## Process Overview

1. **Analyze the PRD** to understand technical implications
2. **Identify if ADRs are needed** for architectural decisions
3. **Create technical design documentation**
4. **Plan implementation** with task breakdown
5. **Prepare for handoff** with checklist

## Usage

### Without Arguments

When run without arguments (`/design-technical`):

1. Search for all PRDs in `/docs/prd/` directory
2. Check which PRDs don't have corresponding design docs in `/docs/design/`
3. List unprocessed PRDs and prompt user to select one
4. Create technical design for the selected PRD

### With PRD Argument

When user provides a PRD or asks for technical design:

1. Read the PRD thoroughly
2. Analyze technical impacts and architecture compatibility
3. Determine if any Architecture Decision Records (ADRs) are needed
4. Create comprehensive technical design documentation
5. Break down implementation into tasks
6. Provide clear handoff communication

## Key Principles

- Start with understanding the requirements
- Only create ADRs for significant architectural decisions
- Consider security, performance, and maintainability
- Make trade-offs explicit
- Design collaboratively

## Remember

- Check existing ADRs for conflicts
- Use project conventions and patterns
- Don't over-engineer solutions
- Focus on maintainability

@.agent/workflows/technical-design.md
