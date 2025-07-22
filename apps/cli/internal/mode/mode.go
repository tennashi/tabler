package mode

// Mode represents the input mode for task creation
type Mode string

const (
	QuickMode    Mode = "quick"
	TalkMode     Mode = "talk"
	PlanningMode Mode = "planning"
)

// ParseModePrefix extracts mode prefix from input and returns the remaining task description
func ParseModePrefix(input string) (Mode, string, bool) {
	// Check shortcuts first (more common)
	if len(input) > 2 && input[2] == ' ' {
		switch input[:2] {
		case "/q":
			return QuickMode, input[3:], true
		case "/t":
			return TalkMode, input[3:], true
		case "/p":
			return PlanningMode, input[3:], true
		}
	}

	// Check full mode names
	if len(input) > 5 && input[5] == ' ' {
		switch input[:5] {
		case "/talk":
			return TalkMode, input[6:], true
		case "/plan":
			return PlanningMode, input[6:], true
		}
	}
	if len(input) > 6 && input[6] == ' ' && input[:6] == "/quick" {
		return QuickMode, input[7:], true
	}

	// No prefix found
	return QuickMode, input, false
}
