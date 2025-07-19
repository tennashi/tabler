# Documentation

This directory contains project documentation.

## Structure

- `prd/` - Product Requirements Documents (what to build and why)
- `design/` - Technical Design Documents (how to build it)
- `adr/` - Architecture Decision Records (technical decisions)

## File Organization

### Simple Features
```
prd/feature_name.md
design/feature_name.md
```

### Complex Features (multi-phase)
```
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