# Agent Roles

This directory contains role definitions for specialized AI agents. Each role follows a consistent structure that emphasizes continuous improvement through observation and measurement.

## Role Philosophy

Every agent role participates in a complete development lifecycle:

```
Plan → Build → Review → Maintain → Learn
  ↑                                    ↓
  └────────────────────────────────────┘
```

## The Five Core Roles

### 1. **Planner** - The Strategist
Transforms requirements into actionable plans
- Breaks down complex requests
- Identifies dependencies
- Creates clear task definitions
- Optimizes execution order

### 2. **Builder** - The Creator  
Constructs high-quality deliverables
- Writes clean code
- Creates documentation
- Follows patterns
- Builds for reusability

### 3. **Reviewer** - The Guardian
Ensures quality and correctness
- Verifies requirements are met
- Identifies issues early
- Provides constructive feedback
- Validates security and performance

### 4. **Maintainer** - The Keeper
Preserves system health
- Manages dependencies
- Reduces technical debt
- Optimizes performance
- Ensures reliability

### 5. **Learner** - The Sage
Extracts and shares knowledge
- Identifies patterns
- Documents lessons
- Prevents repeated mistakes
- Improves team effectiveness

## How Roles Work Together

```
User Request
    ↓
[Planner] → Creates structured plan
    ↓
[Builder] → Implements solution
    ↓
[Reviewer] → Validates quality
    ↓
[Maintainer] → Ensures long-term health
    ↓
[Learner] → Extracts patterns for improvement
    ↓
Better processes for next iteration
```

## Role Structure

Each role document contains:

1. **Purpose**: One-line summary of the role
2. **Responsibility Scope**: What the role owns (and doesn't own)
3. **Observable Metrics**: Primary and secondary indicators
4. **Improvement Cycle**: The 5-step process (Observe → Analyze → Plan → Execute → Verify)
5. **Decision Framework**: How to interpret metrics and make decisions
6. **Interactions**: Dependencies and collaborations with other roles
7. **Anti-patterns**: Common pitfalls to avoid
8. **Examples**: Real scenarios showing the role in action

## Key Principles

### Metrics as Health Indicators
- Metrics reveal symptoms, not goals
- Understanding context is crucial
- Root cause analysis over quick fixes

### Continuous Improvement
- Small, incremental changes
- Measure impact of changes
- Learn from both successes and failures

### Clear Boundaries
- Each role has specific responsibilities
- Overlap is minimal and intentional
- Collaboration points are well-defined

## Available Roles

- `planner.md` - Strategic planning and task breakdown
- `builder.md` - Implementation and construction
- `reviewer.md` - Quality assurance and validation
- `maintainer.md` - System health and sustainability
- `learner.md` - Knowledge extraction and sharing
- `template.md` - Template for creating new roles