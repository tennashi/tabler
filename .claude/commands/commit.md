# Commit Workflow

Execute the commit workflow defined in `.agent/workflows/commit.md`:

1. **Review Git Conventions**
   - Read the commit conventions from `.agent/guidelines/commit.md` (but don't display the entire file)
   - Read the commit workflow from `.agent/workflows/commit.md` for detailed steps
   - Show only a brief summary of key points:
     - Purpose clarification before staging
     - Quality checks requirement
     - Atomic commits principle
     - Commit message format

2. **Check Current Status**
   - Run `git status` to show current changes
   - Run `git diff --cached` to review staged changes

3. **Analyze and Commit**
   - Analyze the changes and suggest appropriate commit message(s)
   - If there are unrelated changes, suggest splitting into multiple commits
   - Show the proposed commit(s) with their messages
   - Execute the commit(s) immediately without asking for confirmation
   - Show the result of each commit

Note: If the user wants to modify the commit message or approach, they can interrupt or provide feedback after seeing the proposal.