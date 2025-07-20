# Commit Workflow

Execute the following steps for making a git commit:

1. **Review Git Conventions**
   - Read and display the commit conventions from `.agent/git-conventions.md`
   - Emphasize these key points:
     - Commit granularity (atomic commits)
     - Commit order
     - Commit message format
     - Testing requirements

2. **Check Current Status**
   - Run `git status` to show current changes
   - Run `git diff --cached` to review staged changes

3. **Prepare Commit**
   - Suggest a commit message following the conventions
   - Ask for user confirmation

4. **Execute Commit**
   - After user approval, execute the actual commit