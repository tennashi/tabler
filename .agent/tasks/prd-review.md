# PRD Review Task

## Purpose

Review Product Requirements Documents to ensure they follow the Why/What principle and avoid implementation details (How).

## When to Execute

- After creating a new PRD
- Before finalizing a PRD for implementation
- When updating existing PRDs
- During PRD review cycles

## Steps

### 1. Locate PRD to Review

```bash
ls -la docs/prd/
```

Identify the PRD file that needs review.

### 2. Read and Analyze the PRD

Read the PRD thoroughly and check against the following criteria.

### 3. Check for Why/What Focus

#### Why Elements (✅ Should Include)

- [ ] **Problem Statement**: Clear description of the problem being solved
- [ ] **User Needs**: Explicit user pain points and needs
- [ ] **Business Value**: Impact and value to the business
- [ ] **Objectives**: Clear goals and desired outcomes
- [ ] **Success Metrics**: Measurable success criteria from user perspective

#### What Elements (✅ Should Include)

- [ ] **User Stories**: User-facing functionality descriptions
- [ ] **Scenarios**: Expected behaviors using Given-When-Then format
- [ ] **Requirements**: Features described from user perspective
- [ ] **Scope**: Clear boundaries of what's included/excluded

### 4. Check for How Elements (❌ Should Remove)

Look for and flag any implementation details:

#### Technical Implementation

- [ ] No specific UI components mentioned (e.g., "single-line input", "Enter key")
- [ ] No technical architecture specified (e.g., "LLM API", "SQLite", "Frontend framework")
- [ ] No internal processing details (e.g., "parsing system", "AI processes")
- [ ] No specific performance numbers (e.g., "100ms", "2 seconds")

#### Common Patterns to Replace

| ❌ Implementation Detail | ✅ User-Focused Alternative |
|-------------------------|----------------------------|
| "Single-line task input with Enter to save" | "Minimal friction task capture interface" |
| "Use LLM API service" | "Natural language understanding capability" |
| "Store in SQLite database" | "Persistent local storage" |
| "Response time < 200ms" | "Instant feedback for user actions" |
| "Frontend framework decision" | "User interface infrastructure" |
| "Authentication system" | "User identification capability" |

### 5. Verify Given-When-Then Scenarios

Each scenario should be:
- [ ] Written from user perspective (not system internals)
- [ ] Testable without knowing implementation
- [ ] Focused on observable behavior
- [ ] Free from technical specifications

### 6. Document Review Findings

Create a summary of findings:

```markdown
## PRD Review: [PRD Name]

### Why/What Coverage
- ✅ Problem clearly stated
- ✅ User needs identified
- ✅ Success metrics defined
- [Add other strengths]

### Implementation Details to Remove
- ⚠️ Line X: Contains "[specific issue]"
- ⚠️ Line Y: References "[technical detail]"
- [List all How elements found]

### Recommendations
- Replace "[implementation]" with "[user-focused alternative]"
- Move "[technical detail]" to Design Doc
- [Other specific recommendations]
```

### 7. Suggest Improvements

For each implementation detail found:
1. Propose user-focused alternative wording
2. Note which Design Doc should contain the technical detail
3. Ensure no loss of important requirements

### 8. Verify Design Doc Coverage

Check that removed implementation details are captured in corresponding Design Docs:

```bash
ls -la docs/design/
```

Ensure technical details have a home in the appropriate Design Doc.

## Expected Result

- PRD focuses exclusively on Why (problems, needs, value) and What (user-facing functionality)
- All How (implementation details) moved to Design Docs
- Clear separation of concerns between PRD and Design Doc
- PRD remains readable by non-technical stakeholders

## Success Criteria

- [ ] No technical implementation details in PRD
- [ ] All user stories written from user perspective
- [ ] Success metrics are user-observable
- [ ] Dependencies listed as capabilities, not technologies
- [ ] PRD readable by product managers, designers, and stakeholders

## Common Mistakes to Avoid

1. **Being too vague**: "Good user experience" vs "Users can complete task in one step"
2. **Hidden implementation**: "Intelligent system" might hide "uses LLM"
3. **Performance as user need**: Users want "fast" not "200ms" - be qualitative
4. **Assuming technical knowledge**: PRD should be understandable without technical background

## Reference

See `docs/prd/template.md` for PRD writing guidelines and examples.