package clarification

import (
	"strings"
)

// VaguenessDetector identifies when task input needs clarification
type VaguenessDetector struct {
	genericWords     []string
	vagueVerbs       []string
	minWordCount     int
	clarityThreshold float64
}

// NewVaguenessDetector creates a new vagueness detector
func NewVaguenessDetector() *VaguenessDetector {
	return &VaguenessDetector{
		genericWords: []string{
			"thing", "things", "stuff", "it", "that", "this",
			"something", "everything", "anything",
		},
		vagueVerbs: []string{
			"do", "handle", "deal", "work", "fix", "check",
			"look", "see", "get", "prepare",
		},
		minWordCount:     3,
		clarityThreshold: 0.4,
	}
}

// DetectVagueness analyzes task input and returns if it needs clarification
func (d *VaguenessDetector) DetectVagueness(input string) (bool, float64) {
	lowerInput := strings.ToLower(strings.TrimSpace(input))
	words := strings.Fields(lowerInput)

	// Calculate vagueness score (0.0 = clear, 1.0 = very vague)
	score := 0.0

	// Factor 1: Length
	score += d.scoreLengthFactor(len(words))

	// Factor 2: Generic words
	score += d.scoreGenericWords(words)

	// Factor 3: Vague verbs
	score += d.scoreVagueVerbs(words)

	// Factor 4: Question marks
	score += d.scoreQuestionMarks(input)

	// Factor 5: Lacks specifics
	score += d.scoreSpecificity(input, lowerInput, words)

	// Factor 6: Missing context
	score += d.scoreMissingContext(lowerInput)

	// Cap score at 1.0
	if score > 1.0 {
		score = 1.0
	}

	// Determine if clarification is needed
	isVague := score >= d.clarityThreshold

	return isVague, score
}

// scoreLengthFactor scores based on task length
func (d *VaguenessDetector) scoreLengthFactor(wordCount int) float64 {
	switch wordCount {
	case 1:
		return 0.6
	case 2:
		return 0.3
	default:
		return 0.0
	}
}

// scoreGenericWords scores based on generic word usage
func (d *VaguenessDetector) scoreGenericWords(words []string) float64 {
	genericCount := 0
	for _, word := range words {
		for _, generic := range d.genericWords {
			if word == generic {
				genericCount++
				break
			}
		}
	}
	if genericCount > 0 {
		return float64(genericCount) * 0.2
	}
	return 0.0
}

// scoreVagueVerbs scores based on vague verb usage
func (d *VaguenessDetector) scoreVagueVerbs(words []string) float64 {
	for _, word := range words {
		for _, verb := range d.vagueVerbs {
			if word == verb {
				return 0.3
			}
		}
	}
	return 0.0
}

// scoreQuestionMarks scores based on presence of questions
func (d *VaguenessDetector) scoreQuestionMarks(input string) float64 {
	if strings.Contains(input, "?") {
		return 0.4
	}
	return 0.0
}

// scoreSpecificity scores based on lack of specific details
func (d *VaguenessDetector) scoreSpecificity(input, lowerInput string, words []string) float64 {
	// Check for numbers
	for _, char := range input {
		if char >= '0' && char <= '9' {
			return 0.0
		}
	}

	// Check for dates
	if d.containsDateWords(lowerInput) {
		return 0.0
	}

	// Check for proper nouns
	if d.containsProperNouns(input) {
		return 0.0
	}

	// No specifics found and task is not too short
	if len(words) > 2 {
		return 0.2
	}
	return 0.0
}

// containsDateWords checks if input contains date-related words
func (d *VaguenessDetector) containsDateWords(lowerInput string) bool {
	dateWords := []string{
		"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday",
		"today", "tomorrow", "week", "month",
		"january", "february", "march", "april", "may", "june",
		"july", "august", "september", "october", "november", "december",
	}
	for _, dateWord := range dateWords {
		if strings.Contains(lowerInput, dateWord) {
			return true
		}
	}
	return false
}

// containsProperNouns checks for capitalized words (simple heuristic)
func (d *VaguenessDetector) containsProperNouns(input string) bool {
	words := strings.Fields(input)
	for i := 1; i < len(words); i++ {
		if len(words[i]) > 0 && words[i][0] >= 'A' && words[i][0] <= 'Z' {
			return true
		}
	}
	return false
}

// scoreMissingContext scores based on incomplete phrases
func (d *VaguenessDetector) scoreMissingContext(lowerInput string) float64 {
	contextPhrases := []string{
		"send the", "prepare the", "finish the", "complete the",
		"review the", "update the", "fix the",
	}

	for _, phrase := range contextPhrases {
		if idx := strings.Index(lowerInput, phrase); idx != -1 {
			// Check what comes after the phrase
			afterPhrase := lowerInput[idx+len(phrase):]
			afterWords := strings.Fields(afterPhrase)
			if len(afterWords) <= 1 {
				return 0.3
			}
		}
	}
	return 0.0
}
