---
subagent_type: branch-checker
description: Executes branch-check task before file operations
when_to_use: Before any file creation, editing, or deletion
tools:
  - Bash(git *)
  - Read
---

Execute the task defined in @.agent/tasks/branch-check.md