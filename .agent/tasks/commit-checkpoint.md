# Commit Checkpoint Task

Execute this task when work reaches a natural stopping point and changes should be committed.

## When to Execute

- After completing a feature or bug fix
- Before switching to a different task
- When reaching a stable state in implementation
- At the end of a work session
- When explicitly requested by the user

## Prerequisites

1. Ensure you have uncommitted changes worth committing
2. Review the current branch status

## Execution Steps

1. **Check Current Status**
   ```bash
   git status
   git diff --cached
   git diff
   ```

2. **Stage Appropriate Files**
   ```bash
   git add <relevant-files>
   ```
   - Stage only files related to the current work
   - Avoid staging unrelated changes

3. **Create Commit**
   - Follow conventions in `../guidelines/commit.md`
   - Write a clear, descriptive commit message
   - Use Conventional Commits format

4. **Verify Commit**
   ```bash
   git log -1 --stat
   ```

## Important Notes

- This task does NOT push to remote unless explicitly requested
- Always follow Conventional Commits format
- Group related changes into logical commits
- If work is incomplete, consider using a WIP (Work In Progress) commit

## References

- Commit Guidelines: `../guidelines/commit.md`
- Branch Strategy: `../guidelines/branch-strategy.md`