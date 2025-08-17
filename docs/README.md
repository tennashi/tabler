# Documentation

This directory contains project documentation.

## Structure

- `prd/` - Product Requirements Documents (what to build and why)
- `design/` - Technical Design Documents (how to build it)
- `adr/` - Architecture Decision Records (technical decisions)

### PRD vs Design Doc Guidelines

**PRD (Product Requirements Document)**:

- Focus: WHY we're building it (problems, user needs, business value)
- Focus: WHAT we're building (user-facing functionality, outcomes)
- Avoid: HOW to implement (technical details, architecture, specific technologies)

**Design Doc**:

- Focus: HOW to implement the requirements from PRD
- Include: Architecture, technical choices, implementation details
- Include: Performance specs, database schemas, API designs

**Example**:

- PRD: "Users need instant feedback when creating tasks"
- Design: "Implement with <100ms response time using local SQLite cache"

## File Organization

### Simple Features

```text
prd/feature_name.md
design/feature_name.md
```

### Complex Features (multi-phase)

```text
prd/feature_name.md
design/feature_name/
├── README.md                # Overview and phases
├── phase1_core.md          # Core functionality
├── phase2_enhanced.md      # Enhanced features
└── phase3_advanced.md      # Advanced capabilities
```

## File Naming

- PRD and Design docs: Use snake_case feature names (e.g., `smart_task_creation.md`)
- ADR: Use sequential numbers (e.g., `001_frontend_framework.md`)
- Phase docs: Prefix with phase number (e.g., `phase1_basic_task_creation.md`)
