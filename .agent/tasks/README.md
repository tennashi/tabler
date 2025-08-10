# Tasks

This directory contains **individual executable tasks** that AI agents must perform at specific points in their workflow.

## Purpose

Tasks are **specific actions to perform** - discrete operations that ensure consistency and safety in the development process.

## Files in this Directory

- `branch-check.md` - Verify and ensure proper branch before file operations
- `commit-checkpoint.md` - Create automatic commits at regular intervals during work

## When to Execute These Tasks

### Required Tasks

- **branch-check.md**
  - **MUST** execute BEFORE any file operations (create/edit/delete)
  - Ensures work is done on the correct branch
  - Prevents accidental changes to main/master branches

- **commit-checkpoint.md**
  - **SHOULD** execute after making significant file changes
  - Creates checkpoint commits for work in progress
  - Helps maintain granular history of changes

## Task Execution

Each task file contains:

1. Purpose and context
2. Step-by-step instructions
3. Success criteria
4. Error handling procedures
