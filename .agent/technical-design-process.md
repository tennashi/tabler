# Technical Design Process for AI Agents

When receiving a PRD, engage in systematic technical design to ensure robust implementation.

## Phase 1: PRD Analysis and Technical Assessment

### Initial Review
1. Read PRD thoroughly to understand requirements
2. Identify technical implications and challenges
3. Check compatibility with existing architecture
4. List potential risks and constraints

### Architecture Impact Analysis
**Assess if the feature requires:**
- New technology integrations
- Changes to existing system architecture
- Performance optimizations
- Security enhancements
- Data model changes

**Flag ADR needs when:**
- Multiple viable technical approaches exist
- Significant trade-offs need documentation
- Existing architectural decisions need revision
- Long-term maintainability is affected

Example:
```
"Based on the PRD analysis, I've identified:
- Technical challenges: Real-time sync, offline support
- Architecture impacts: Need local storage strategy
- ADR required: Yes - for offline data sync approach
- Existing ADR conflicts: ADR-003 assumes cloud-only"
```

## Phase 2: Architecture Decisions (When Needed)

### Determine ADR Necessity
Not every feature needs an ADR. Create one when:
- The decision has system-wide impact
- Multiple viable options exist with different trade-offs
- The decision is hard to reverse later
- Team members might question "why" in the future

### ADR Creation Process
1. **Identify the decision scope**
   - What exactly needs to be decided?
   - Who will be affected?

2. **Research and present options**
   ```
   "For offline storage, I see three options:
   A. IndexedDB - Good capacity, browser-only
   B. Local SQLite - Full SQL, needs native app
   C. Hybrid approach - IndexedDB + sync
   
   Let's evaluate trade-offs..."
   ```

3. **Document the decision**
   - Use the project's ADR template
   - Include context, options, decision, consequences

## Phase 3: Technical Design Documentation

### Design Document Structure
1. **Overview**
   - Feature summary from technical perspective
   - High-level approach

2. **System Design**
   - Architecture diagrams
   - Component interactions
   - Data flow

3. **Detailed Design**
   - API specifications
   - Data models
   - State management
   - Error handling strategy

4. **Security Considerations**
   - Authentication/authorization impacts
   - Data privacy concerns
   - Input validation needs

5. **Performance Considerations**
   - Expected load
   - Optimization strategies
   - Caching approach

### Example Prompts
```
"Let me create a technical design for the smart task creation feature:

1. Architecture: Client-side parser + API integration
2. Data model: Task entity with metadata fields
3. API endpoints: POST /tasks with NLP processing
4. Performance: Debounced input, client-side caching"
```

## Phase 4: Implementation Planning

### Task Breakdown
1. **Identify implementation layers**
   - Frontend components
   - Backend services
   - Database changes
   - Infrastructure needs

2. **Define task dependencies**
   - What must be built first?
   - What can be parallelized?

3. **Estimate complexity**
   - Story points or time estimates
   - Risk factors

### Testing Strategy
- Unit test requirements
- Integration test scenarios
- Performance test needs
- User acceptance criteria

## Phase 5: Review and Handoff

### Pre-implementation Checklist
- [ ] All PRD requirements mapped to technical solutions
- [ ] ADRs created for significant decisions
- [ ] Design document complete and clear
- [ ] Implementation tasks defined
- [ ] Test strategy documented
- [ ] Team questions addressed

### Handoff Communication
```
"Technical design is complete:
- Design doc: [link]
- New ADRs: [list]
- Implementation tasks created
- Estimated effort: X sprint points
- Key risks: [list]

Ready to begin implementation?"
```

## Key Principles

- **Start with understanding** - Don't jump to solutions
- **Document significant decisions** - But don't over-document
- **Consider maintainability** - Think beyond initial implementation
- **Make trade-offs explicit** - No perfect solutions
- **Collaborate with the team** - Design is a team sport

## Don'ts

- Don't create ADRs for trivial decisions
- Don't skip security and performance considerations
- Don't design in isolation - involve the team
- Don't ignore existing patterns and conventions
- Don't over-engineer - match complexity to requirements