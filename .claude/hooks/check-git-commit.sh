#!/bin/bash
# Hook to remind Claude about git conventions before commits

# Read JSON input from stdin
INPUT=$(cat)

# Extract command from tool_input.command
if command -v jq &> /dev/null; then
    COMMAND=$(echo "$INPUT" | jq -r '.tool_input.command // empty')
else
    # Fallback: extract command using grep/sed
    COMMAND=$(echo "$INPUT" | grep -o '"command":"[^"]*"' | sed 's/.*"command":"\([^"]*\)".*/\1/')
fi

# Check if command contains git add or git commit
if [[ "$COMMAND" =~ git\ (add|commit) ]]; then
    echo "📋 Git Conventions Reminder:"
    echo "================================"
    
    # Show key points from git-conventions.md
    if [[ -f ".agent/git-conventions.md" ]]; then
        echo ""
        echo "Key reminders from .agent/git-conventions.md:"
        echo ""
        
        # Extract and show commit granularity rules
        sed -n '/## Commit Granularity/,/## Commit Order/p' .agent/git-conventions.md | grep "^-" | head -5
        
        echo ""
        echo "Commit Order:"
        sed -n '/## Commit Order/,/## Commit Messages/p' .agent/git-conventions.md | grep "^-" | head -5
        
        echo ""
        echo "================================"
        echo "Full conventions at: .agent/git-conventions.md"
    else
        echo "Warning: .agent/git-conventions.md not found"
    fi
    
    echo ""
fi

# Always allow execution (exit 0)
# To block execution, use exit 1
exit 0