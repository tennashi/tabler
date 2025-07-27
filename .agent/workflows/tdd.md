# TDD Conventions for AI Agents

This document defines how AI agents should approach Test-Driven Development in this project.

## Core Principles

Following Kent Beck's TDD and t-wada's TDD Boot Camp style:

1. **TODO List Driven Development**
2. **Red-Green-Refactor Cycle**
3. **AAA Pattern** (Arrange-Act-Assert)
4. **Tests as Specifications**

## AI Agent Development Process

### 1. Understand the Requirement

When given a feature request:

- Break it down into small, testable behaviors
- Create a TODO list of test cases
- **Clearly state the implementation approach:**
  - Bottom-up: Start from lowest level components
  - Prioritize integration: Connect components as soon as possible
  - Avoid isolated components: Each layer should be integrated immediately
- Share the TODO list with the user for confirmation

Example interaction:

````text
User: "Implement task creation with inline shortcuts"

AI Agent: "I'll break this down into testable behaviors:
- [ ] Task can be created with just a title
- [ ] Empty title should return an error
- [ ] Parse single tag with # prefix
- [ ] Parse multiple tags
- [ ] Parse priority with ! prefix
- [ ] Parse due date with @ prefix
- [ ] Combine multiple shortcuts in one input

Implementation approach: Bottom-up with early integration
- Start with the most fundamental components
- As soon as each component is ready, integrate it with others
- Prioritize end-to-end functionality over isolated perfection

Shall I start with the first test?"
```text

### 2. TDD Cycle for Each TODO Item

For each TODO item, follow this cycle:

#### 2.1 RED Phase - Write Failing Test

```go
func TestTask(t *testing.T) {
    t.Run("creation with valid title should succeed", func(t *testing.T) {
        // Arrange
        title := "Buy groceries"
        
        // Act
        task := NewTask(title)
        
        // Assert
        if task.Title != title {
            t.Errorf("expected title %q, got %q", title, task.Title)
        }
    })
}
```text

Run and show failure:

```bash
$ go test
# undefined: NewTask
```text

**⚠️ DO NOT COMMIT - Test is failing**

#### 2.2 GREEN Phase - Make Test Pass

Write minimal code:

```go
type Task struct {
    Title string
}

func NewTask(title string) *Task {
    return &Task{Title: title}
}
```text

Run and show success:

```bash
$ go test
PASS
```text

**⚠️ DO NOT COMMIT YET - Code may need refactoring**

#### 2.2.1 When Tests Fail in GREEN Phase

If a test doesn't pass as expected during the GREEN phase:

1. **Analyze why the test is failing**
   - Understand the root cause (e.g., timing issues, ordering assumptions, incorrect expectations)
   - Add debug output if needed to understand the actual behavior

2. **Consider implementation fixes FIRST**
   - Can the implementation be adjusted to meet the test's expectations?
   - Example: Add ORDER BY clause for consistent ordering
   - Example: Add secondary sort key for stable results

3. **Evaluate specification feedback**
   - Does the test expectation reflect the actual requirements?
   - Is the tested behavior actually important for the feature?

4. **Test modification is the LAST resort**
   - Only modify the test if the expectation was genuinely incorrect
   - Document why the test expectation was changed

**Example:**

```text
Problem: Test expects tasks in creation order, but they have same timestamp
Solutions in order of preference:
1. ✅ Add "ORDER BY created_at DESC, id DESC" for stable ordering
2. ✅ Consider if order matters for the feature
3. ❌ Change test to ignore order (only if order truly doesn't matter)
```text

#### 2.3 REFACTOR Phase - Improve Implementation

- Keep tests passing
- Remove duplication
- Improve naming
- Extract constants or helper functions

**Self-Review Checklist - Cohesion and Coupling:**

- [ ] **Cohesion**: Does each function have a single, clear responsibility?
- [ ] **Coupling**: Are functions passing only necessary data as parameters?
- [ ] **Magic Numbers**: Are all literals extracted as named constants?
- [ ] **Testability**: Can each function be tested independently?
- [ ] **Naming**: Do function names accurately describe their single responsibility?

Example:

```go
// Low cohesion: Parse does too many things
func Parse(input string) *Result {
    // Extracts tags, priority, deadline all in one function
}

// High cohesion: Each function has one responsibility
func extractTag(part string) (string, bool)
func extractPriority(part string) (int, bool)
func extractDeadline(part string) (*time.Time, bool)
```text

#### 2.4 TEST REFACTOR Phase - Improve Test Code

After implementation is clean, refactor the test:

**IMPORTANT: This phase is for refactoring the CURRENT test only, not for adding new test cases!**

- New test cases require starting a new RED-GREEN-REFACTOR cycle
- Only refactor if the current test has duplication or clarity issues
- If the test is already clean and simple, skip this phase

**Look for:**

- Duplicate test setup → Extract helper functions
- Similar test patterns → Convert to table-driven tests
- Complex assertions → Create custom assertion helpers
- Unclear test names → Improve readability

**Table-Driven Test Criteria:**

Use table-driven tests when:

- **Same function with different inputs**: Testing one function with various input patterns
- **Clear input-output mapping**: Input → Expected output is straightforward
- **Repeated assertion patterns**: Same checks performed across multiple tests
- **Many edge cases**: Empty strings, special characters, boundary values, combinations
- **DRY principle**: Same test code repeated 3+ times

Avoid table-driven tests when:

- **Complex setup required**: Each test needs significantly different mocks, DB setup, etc.
- **State sharing between tests**: Tests depend on results from previous tests (avoid this pattern anyway)
- **Detailed error inspection**: Need to check error types, wrapped errors, or error fields in detail
- **Async/timing critical**: Tests involving goroutines, channels, or time-sensitive operations
- **Complex test flow**: Multi-step tests with branching logic that doesn't fit well in a table

**Example refactoring:**

```go
// Before: Duplicate setup in multiple tests
func TestTask(t *testing.T) {
    t.Run("creation with valid title", func(t *testing.T) {
        task := NewTask("Buy groceries")
        if task.Title != "Buy groceries" {
            t.Errorf("wrong title")
        }
    })
    
    t.Run("creation with unicode title", func(t *testing.T) {
        task := NewTask("買い物")
        if task.Title != "買い物" {
            t.Errorf("wrong title")
        }
    })
}

// After: Table-driven test
func TestTask(t *testing.T) {
    t.Run("creation", func(t *testing.T) {
        tests := []struct {
            name  string
            title string
        }{
            {"with valid title", "Buy groceries"},
            {"with unicode title", "買い物"},
        }
        
        for _, tt := range tests {
            t.Run(tt.name, func(t *testing.T) {
                task := NewTask(tt.title)
                if task.Title != tt.title {
                    t.Errorf("expected title %q, got %q", tt.title, task.Title)
                }
            })
        }
    })
}
```text

#### 2.5 COMMIT Phase - Save Clean Code

After both implementation and test refactoring:

**IMPORTANT: Read `.agent/guidelines/commit.md` before committing!**

```bash
$ go test  # Ensure all tests pass
$ git add .
$ git commit -m "feat: implement task creation with title"
```text

**✅ NOW IT'S SAFE TO COMMIT - Code and tests are clean**

#### 2.6 Update TODO List

```text
- [x] Task can be created with just a title
- [ ] Empty title should return an error
...
```text

### 3. Commit Guidelines

**⚠️ MUST READ: `.agent/guidelines/commit.md` for commit format and rules**

**When to Commit:**

- After REFACTOR phase when both code and tests are clean
- When switching context or taking a break
- After completing a logical group of related tests

**What to Include in Commit:**

- Test file(s)
- Implementation file(s)
- Any refactoring done in both

**Commit Message Format:**
Follow the conventions in `.agent/guidelines/commit.md`. Example:

```text
feat: implement [feature name]

- Add test for [behavior]
- Implement [what was implemented]
- Refactor [what was refactored] (if applicable)
```text

### 4. Complete Cycle Before Moving On

Only proceed to the next TODO item after:

- Implementation code is clean
- Test code is clean
- All tests are passing
- Changes are committed
- TODO list is updated

**CRITICAL: Each test case is a separate cycle!**

- One test case = One complete RED-GREEN-REFACTOR-COMMIT cycle
- Do NOT write multiple test cases in RED phase
- Do NOT add new test cases during TEST REFACTOR phase
- Each new behavior/edge case starts with a new RED phase

### 5. Integration Tests Come Later

After all unit tests pass:

- Write integration tests for external dependencies
- Use real implementations (database, file system)
- Mark as integration tests with build tags or skip conditions

## Test Structure Guidelines

### Use Subtests for Organization

```go
func TestTask(t *testing.T) {
    t.Run("creation", func(t *testing.T) {
        t.Run("with valid title should succeed", func(t *testing.T) {
            // test implementation
        })
        
        t.Run("with empty title should return error", func(t *testing.T) {
            // test implementation
        })
    })
    
    t.Run("update", func(t *testing.T) {
        t.Run("title should modify the task", func(t *testing.T) {
            // test implementation
        })
    })
}
```text

### AAA Pattern

Every test must follow Arrange-Act-Assert:

```go
t.Run("with valid input should return task", func(t *testing.T) {
    // Arrange - Set up test data and dependencies
    repo := &mockRepository{}
    parser := &mockParser{}
    usecase := NewTaskUsecase(repo, parser)
    
    // Act - Execute the behavior being tested
    task, err := usecase.CreateTask(context.Background(), "Buy milk")
    
    // Assert - Verify the outcome
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if task.Title != "Buy milk" {
        t.Errorf("expected title %q, got %q", "Buy milk", task.Title)
    }
})
```text

### Test Helpers

Extract common test operations:

```go
// Test helper for creating test tasks
func newTestTask(t *testing.T, title string) *Task {
    t.Helper()
    task := NewTask(title)
    if task == nil {
        t.Fatal("task should not be nil")
    }
    return task
}

// Custom assertion
func assertTaskTitle(t *testing.T, task *Task, expected string) {
    t.Helper()
    if task.Title != expected {
        t.Errorf("expected title %q, got %q", expected, task.Title)
    }
}
```text

## Communication with User

### Before Starting

```text
"I'll implement this using TDD. Here's my plan:
[TODO list]
I'll start with the first test. OK?"
```text

### During Development

```text
"RED phase: Writing failing test...
[Show test code]
[Show test failure]

GREEN phase: Implementing minimal code...
[Show implementation]
[Show test passing]

REFACTOR phase: Improving implementation...
[Show refactored code]

TEST REFACTOR phase: Improving test code...
[Show refactored test]

Ready to commit. All tests passing.
Moving to next test..."
```text

### When Finding Test Improvements

```text
"I notice these tests have similar setup. 
Would you like me to:
1. Extract a helper function?
2. Convert to table-driven tests?
3. Keep as is for now?"
```text

## Best Practices for AI Agents

1. **Show all phases** - RED, GREEN, REFACTOR, TEST REFACTOR
2. **One test at a time** - Never write multiple tests before implementing
3. **Clean tests matter** - Tests are documentation, keep them readable
4. **Commit after refactoring** - Only commit clean code
5. **Ask when uncertain** - If requirements are unclear, ask before writing tests
6. **Update TODO lists** - Keep progress visible
7. **Explain refactoring** - Explain why you're refactoring both code and tests

## Common Pitfalls to Avoid

1. **Committing failing tests** - Never commit in RED phase
2. **Committing without refactoring** - Don't commit messy code from GREEN phase
3. **Skipping test refactoring** - Tests accumulate tech debt too
4. **Over-engineering test helpers** - Extract only when there's duplication
5. **Complex test setup** - If setup is complex, consider simpler design
6. **Unclear test names** - Tests should read like specifications
7. **Testing implementation details** - Test behavior, not internals
8. **Adding multiple tests at once** - Each test case needs its own cycle
9. **Refactoring to add new tests** - TEST REFACTOR is for improving existing test only

### Example of WRONG Approach

```text
❌ WRONG: Adding multiple test cases during TEST REFACTOR phase
- Write test for "valid task"
- Make it pass
- Refactor implementation
- TEST REFACTOR: Convert to table-driven test AND add 5 new test cases
```text

### Example of CORRECT Approach

```text
✅ CORRECT: One test case per cycle
Cycle 1:
- RED: Write test for "valid task"
- GREEN: Make it pass
- REFACTOR: Clean up implementation
- COMMIT

Cycle 2:
- RED: Write test for "empty title error"
- GREEN: Make it pass
- REFACTOR: Clean up implementation
- COMMIT

Cycle 3:
- RED: Write test for "unicode title"
- GREEN: Make it pass
- REFACTOR: Clean up implementation
- TEST REFACTOR: Now that we have 3 similar tests, refactor to table-driven
- COMMIT
```text

## Test Time Dependencies

### Prohibited Practices

- **NEVER use `time.Sleep()` in tests** - It's not a real solution and makes tests slow
- Avoid relying on execution order or timing

### Recommended Solutions

1. **Explicit time setting for deterministic tests**
   ```go
   // Use explicit timestamps in tests
   task1 := &task.Task{
       CreatedAt: time.Date(2024, 1, 1, 10, 0, 0, 0, time.UTC),
   }
   task2 := &task.Task{
       CreatedAt: time.Date(2024, 1, 1, 11, 0, 0, 0, time.UTC),
   }
````

1. **Order-independent assertions**

   ```go
   // Don't rely on order, verify by ID or content
   tasksByID := make(map[string]*Task)
   for _, task := range tasks {
       tasksByID[task.ID] = task
   }
   // Assert using the map
   ```

2. **Use interfaces for time providers**

   ```go
   type TimeProvider interface {
       Now() time.Time
   }
   // Inject mock time provider in tests
   ```

## Design Doc Alignment During Implementation

### Track Design Deviations

**During implementation:**

1. **When implementation needs differ from design**
   - Discuss the issue with team/user
   - Update Design Doc:
     - Move current approach to "Alternatives Considered" section
     - Document why it didn't work during implementation
     - Update main sections with new approach
     - Example:

       ```markdown
       ## Alternatives Considered

       ### Channel-based Concurrency

       Originally designed to use channels for concurrent updates.
       **Rejected because**: Implementation revealed deadlock risks
       when multiple goroutines access shared state.

       ## Current Design

       Using mutex locks for thread-safe access...
       ```

   - Commit with clear message:

     ```bash
     git add docs/design/<feature>.md
     git commit -m "fix(design): change concurrency model to mutex-based

     - Original channel design caused deadlocks
     - Mutex approach is simpler and more reliable
     - Moved original design to Alternatives Considered"
     ```

2. **Question unclear design points**
   - If design seems wrong or unclear, ask before proceeding
   - "The design says X, but that would cause Y issue. Should I do Z instead?"

3. **Continue implementation with updated design**
   - Implementation always follows current design doc
   - No divergence between design and implementation

### Post-Implementation Verification

**After completing all TODO items:**

1. **Verify implementation matches current design doc**
   - Design doc should already reflect all changes made during implementation
   - Any deviations should have been resolved during implementation
   - Final implementation should match design doc exactly

2. **Document lessons learned (optional)**
   - If significant insights were gained, consider:
     - Creating an ADR for architectural decisions
     - Updating project conventions
     - Sharing learnings with team

### Example Implementation Flow

````text
"During implementation, I discovered the design's approach to handling 
concurrent updates could cause race conditions.

I'll update the design doc to use mutex locks instead of channels."

[After confirmation]
"I've updated the Design Doc v1.1 with the mutex approach and committed.
Continuing implementation with the updated design..."
```text

## Remember

Clean tests are as important as clean code. They serve as living documentation and make future changes easier. Only commit when both are clean.
````
