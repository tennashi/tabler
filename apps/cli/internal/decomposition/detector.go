package decomposition

import (
	"strings"
)

// ComplexityDetector identifies tasks that would benefit from decomposition
type ComplexityDetector struct {
	complexVerbs []string
}

// NewComplexityDetector creates a new complexity detector
func NewComplexityDetector() *ComplexityDetector {
	return &ComplexityDetector{
		complexVerbs: []string{
			"plan", "organize", "prepare", "develop", "implement",
			"create", "build", "design", "establish", "setup",
			"research", "analyze", "review", "refactor", "migrate",
		},
	}
}

// DetectComplexity analyzes task and returns whether it's complex and why
func (d *ComplexityDetector) DetectComplexity(input string) (bool, string) {
	lowerInput := strings.ToLower(input)

	// Check for complex verbs
	for _, verb := range d.complexVerbs {
		if strings.Contains(lowerInput, verb) {
			return true, "contains complex verb: " + verb
		}
	}

	// Check task length (more than 50 chars suggests multiple actions)
	if len(input) > 50 {
		return true, "task is long and may contain multiple actions"
	}

	// Not complex
	return false, ""
}
