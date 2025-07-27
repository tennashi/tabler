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

````text
"Based on the PRD analysis, I've identified:
- Technical challenges: Real-time sync, offline support
- Architecture impacts: Need local storage strategy
- ADR required: Yes - for offline data sync approach
- Existing ADR conflicts: ADR-003 assumes cloud-only"
```text

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
````

"For offline storage, I see three options:
A. IndexedDB - Good capacity, browser-only
B. Local SQLite - Full SQL, needs native app
C. Hybrid approach - IndexedDB + sync

Let's evaluate trade-offs..."

```text

3. **Document the decision**
- Use the project's ADR template
- Include context, options, decision, consequences

## Phase 3: Technical Design Documentation

### Determining Design Document Scope

**Critical: Keep Design Docs Small and Implementable**

Each Design Doc should be:

- **5-10 implementation tasks** maximum
- **2-3 days** of implementation work
- **15-30 commits** when complete
- Delivers **one clear value** to users

If your design generates more than 10 tasks, split it into multiple phases.

**MUST: Split by Vertical Feature Slices**

Always split Design Docs by complete vertical features that deliver user value:

1. **Vertical feature slicing (REQUIRED)**
- Each Design Doc implements one complete user-facing feature
- Includes all layers: UI → Service → Storage
- Delivers working functionality end-to-end
- Example: "Basic task creation" includes parser, storage, service, and CLI

2. **Avoid horizontal slicing**
- ❌ DON'T: Separate docs for "database layer" vs "service layer"
- ❌ DON'T: Split by technical components
- ✅ DO: Each doc delivers a complete feature

3. **When vertical slice is too large (>10 tasks)**
- Split into core feature + enhancements
- Each enhancement is still a vertical slice
- Example for task creation:
  - Doc 1: Basic task creation (title only)
  - Doc 2: Task creation with shortcuts (tags, priority, deadline)
  - Doc 3: Task creation with AI assistance
- Each delivers complete functionality at its level

**Example split for Smart Task Creation:**

```text
"This PRD covers multiple feature sets. I recommend splitting the design:

Design Doc 1: Basic Task Creation (Phase 1)
- Simple input and storage
- Inline shortcuts parsing
- Basic CRUD operations

Design Doc 2: AI Enhancement (Phase 2)
- LLM integration
- Natural language processing
- Task decomposition

Design Doc 3: Interactive Features (Phase 3)
- Dialogue system
- Mode switching
- Learning/personalization

This allows incremental delivery and focused reviews."
```text

### Design Document Organization

**For single design doc:**

```text
docs/design/[feature-name].md
```text

**For multiple design docs (split by features):**

```text
docs/design/[feature-name]/
├── README.md           # Overview and feature relationships
├── [feature1-name].md  # Complete vertical feature
├── [feature2-name].md  # Another complete feature
└── [feature3-name].md  # Additional functionality
```text

**IMPORTANT: Focus on Features, Not Phases**

When creating design docs from a PRD:

- Ignore phase divisions in the PRD
- Focus on individual features that can be implemented independently
- Each design doc should represent one complete vertical feature
- Name design docs by feature, not by phase number

**README.md should include:**

- Overall architecture vision
- Phase dependencies and relationships
- Implementation order rationale
- Links to related PRD and ADRs

**Example structure:**

```text
docs/design/smart-task-creation/
├── README.md
├── natural-language-dates.md
├── llm-inference.md
├── smart-decomposition.md
├── interactive-clarification.md
├── input-modes.md
└── context-learning.md
```text

### Design Document Structure

**IMPORTANT: Always use the design doc template**

- Use `/docs/design/template.md` as the starting point
- Copy the template and fill in all sections
- Do not skip sections - mark as "Not applicable" if needed
- Follow the detail level guidance in the template

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

```text
"Let me create a technical design for the smart task creation feature:

1. Architecture: Client-side parser + API integration
2. Data model: Task entity with metadata fields
3. API endpoints: POST /tasks with NLP processing
4. Performance: Debounced input, client-side caching"
```text

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

### Task Count Check - Split if Too Large

**After creating task list, evaluate:**

```text
"I've identified 15 tasks for this design. This exceeds the recommended 5-10 tasks.

I recommend splitting this into two phases:

Phase 1: Basic Feature (7 tasks)
- Core functionality that delivers immediate value
- Can be released independently
- 2 days of work

Phase 2: Enhanced Feature (8 tasks)  
- Builds on Phase 1
- Adds advanced capabilities
- 2-3 days of work

Shall I create separate design docs for each phase?"
```text

**Red flags that indicate splitting:**

- More than 10 implementation tasks
- Estimated time > 3 days
- Mixed core and "nice-to-have" features
- Complex dependencies between task groups

**How to split:**

1. Identify the minimal valuable feature
2. Group tasks that deliver that value
3. Move enhancements to next phase
4. Each phase should stand alone

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

```text
"Technical design is complete:
- Design doc: [link]
- New ADRs: [list]
- Implementation tasks created
- Estimated effort: X sprint points
- Key risks: [list]

Ready to begin implementation?"
```text

## Phase 6: PRD Feedback During Design

### When to Update PRD

During technical design, you may discover that PRD requirements need adjustment:

1. **Technical Impossibilities**
- Requirement cannot be implemented as specified
- Platform/technology limitations discovered
- Security constraints prevent implementation

2. **Significant Complexity**
- Requirement adds disproportionate complexity
- Simpler alternative achieves same user value
- Implementation cost exceeds projected benefit

3. **Missing Requirements**
- Discovered edge cases not in PRD
- Required supporting features not specified
- Integration requirements not documented

### How to Provide Feedback

1. **Document the issue in design doc**
2. **Discuss with product owner/user**
3. **Update PRD directly:**
- Increment version number in PRD
- Make necessary changes
- Commit with clear message:
  ```bash
  git add docs/prd/<feature>.md
  git commit -m "fix(prd): update real-time sync to polling approach

  - Changed from WebSocket to 5-second polling
  - Reason: WebSocket requires new infrastructure (3x complexity)
  - Polling meets user needs with existing architecture
  - Design doc identified this during technical assessment"
  ```
4. **Continue design with updated requirements**

### Example Flow

```text
"During technical design, I discovered that the real-time sync requirement 
would require WebSocket infrastructure, increasing complexity by 3x.

A 5-second polling approach would meet the user's actual need (staying updated)
using our existing architecture.

Shall I update the PRD with this change?"

[After confirmation]
"I've updated the PRD v1.1 with the polling approach and committed the change.
Continuing with the technical design based on this update..."
```text

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
````
