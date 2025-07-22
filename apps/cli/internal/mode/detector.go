package mode

import (
	"strings"
)

// ModeDetector detects appropriate mode based on input characteristics
type ModeDetector struct{}

// NewModeDetector creates a new mode detector
func NewModeDetector() *ModeDetector {
	return &ModeDetector{}
}

// DetectMode analyzes input and returns the recommended mode
func (d *ModeDetector) DetectMode(input string) Mode {
	// Very short input (<10 chars) → Quick mode
	if len(input) < 10 {
		return QuickMode
	}

	// Contains question words → Talk mode
	lowerInput := strings.ToLower(input)
	questionWords := []string{"what", "how", "why", "when", "where", "should", "could", "would"}
	for _, word := range questionWords {
		if strings.Contains(lowerInput, word) {
			return TalkMode
		}
	}

	// Contains planning keywords → Planning mode
	planningKeywords := []string{"plan", "organize", "prepare", "strategy", "roadmap", "project"}
	for _, keyword := range planningKeywords {
		if strings.Contains(lowerInput, keyword) {
			return PlanningMode
		}
	}

	// Default → Quick mode
	return QuickMode
}
