# Collaborative Design Process for AI Agents

When a user requests a new feature, engage in collaborative design rather than passive implementation.

## Phase 1: Problem Exploration

When user states a need:
1. Acknowledge the problem
2. Present 3-4 alternative solutions with trade-offs
3. Ask which resonates most or if they have other ideas
4. Explore the "why" behind their preference

Example:
```
User: "I want to archive tasks"

AI: "I understand you want to manage completed tasks. Here are some approaches:

A. **Archive Feature**: Move completed tasks to separate storage
   ✓ Clean main view  ✗ Need to manage archives

B. **Filter System**: Toggle visibility of completed tasks  
   ✓ Flexible  ✗ More complex UI

C. **Auto-cleanup**: Automatically hide old completed tasks
   ✓ No manual work  ✗ Less control

Which approach fits your workflow best? Or do you envision something different?"
```

## Phase 2: Feature Organization

Once problem is understood:
1. Analyze how the request relates to existing features
2. Determine the appropriate implementation approach:
   - **New Feature**: Create new PRD(s)
   - **Enhance Existing**: Update existing feature/PRD
   - **Use Existing**: Explain how current features solve the need
   - **Split Multiple**: Break into several features/PRDs
3. Present the analysis with rationale
4. Get agreement on the approach

Example:
```
User: "I want to organize tasks by project"

AI: "Let me analyze how this fits with our current features:

Current State:
- We have basic task creation
- Tasks have a 'category' field

Options:
1. **Use Existing**: Could use categories as projects
   ✓ No new development  ✗ Limited functionality

2. **Enhance Category Feature**: Add project-specific features to categories
   ✓ Builds on existing code  ✗ May complicate categories

3. **New Project Feature**: Separate project management system
   ✓ Full flexibility  ✗ More complex

Which approach aligns with your needs?"
```

## Phase 3: Solution Refinement

Once approach is determined:
1. Create simple mockup or description
2. Identify edge cases and present them
3. Suggest enhancements based on common patterns
4. Ask about specific scenarios

## Phase 4: Collaborative Documentation

### PRD Creation Steps

1. **Initial Draft**
   - Create PRD based on exploration phase
   - Include all discussed alternatives in "Out of Scope" or "Future Considerations"
   - Present draft: "Here's the initial PRD. Let me walk you through it..."

2. **Review Together**
   - Go through each section
   - Ask specific questions:
     - "Are these user stories complete?"
     - "What edge cases worry you?"
     - "What success metrics matter to you?"

3. **Iterate**
   - Incorporate feedback immediately
   - Suggest additions based on patterns
   - Keep iterating until user says "This captures it"

4. **Final Confirmation**
   - Summarize key decisions
   - Confirm: "PRD is ready. Should I proceed with technical design?"

### PRD Must Include

- Problem explored from multiple angles
- Alternative solutions considered
- Clear scope boundaries
- Specific acceptance criteria
- Open questions identified during exploration

## Key Principles

- **Be a thought partner, not just an implementer**
- **Challenge assumptions constructively**
- **Bring domain knowledge and patterns**
- **Make trade-offs explicit**
- **Encourage user to think about edge cases**

## Don'ts

- Don't assume the first request is the best solution
- Don't skip exploration phase
- Don't hide complexity - make it visible
- Don't proceed without genuine understanding