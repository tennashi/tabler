package decomposition

import (
	"fmt"
	"strconv"
	"strings"
)

// InteractivePresenter formats decomposition results for interactive display
type InteractivePresenter struct{}

// NewInteractivePresenter creates a new interactive presenter
func NewInteractivePresenter() *InteractivePresenter {
	return &InteractivePresenter{}
}

// Present formats the decomposition result for display
func (p *InteractivePresenter) Present(result *DecompositionResult) string {
	var b strings.Builder

	// Header
	b.WriteString("Task decomposition for: ")
	b.WriteString(result.OriginalTask)
	b.WriteString("\n\n")

	// Subtasks
	for i, subtask := range result.Subtasks {
		b.WriteString(fmt.Sprintf("[%d] %s\n", i+1, subtask))
	}

	// Instructions
	b.WriteString("\nSelect subtasks to create (e.g., '1,3-5' or 'all' or 'none'): ")

	return b.String()
}

// ParseSelection parses user input for subtask selection
func (p *InteractivePresenter) ParseSelection(input string, total int) ([]int, error) {
	input = strings.TrimSpace(input)

	// Handle special cases
	if input == "all" {
		return p.selectAll(total), nil
	}

	if input == "none" {
		return []int{}, nil
	}

	// Parse comma-separated values
	return p.parseCommaSeparated(input, total)
}

// selectAll returns all indices from 1 to total
func (p *InteractivePresenter) selectAll(total int) []int {
	selections := make([]int, total)
	for i := 0; i < total; i++ {
		selections[i] = i + 1
	}
	return selections
}

// parseCommaSeparated parses comma-separated selection values
func (p *InteractivePresenter) parseCommaSeparated(input string, total int) ([]int, error) {
	var selections []int
	parts := strings.Split(input, ",")

	for _, part := range parts {
		part = strings.TrimSpace(part)

		if strings.Contains(part, "-") {
			// Parse range
			nums, err := p.parseRange(part, total)
			if err != nil {
				return nil, err
			}
			selections = append(selections, nums...)
		} else {
			// Parse single number
			num, err := p.parseSingleNumber(part, total)
			if err != nil {
				return nil, err
			}
			selections = append(selections, num)
		}
	}

	return selections, nil
}

// parseRange parses a range like "1-3"
func (p *InteractivePresenter) parseRange(part string, total int) ([]int, error) {
	rangeParts := strings.Split(part, "-")
	if len(rangeParts) != 2 {
		return nil, fmt.Errorf("invalid range format: %s", part)
	}

	start, err := strconv.Atoi(strings.TrimSpace(rangeParts[0]))
	if err != nil {
		return nil, fmt.Errorf("invalid number in range: %s", rangeParts[0])
	}

	end, err := strconv.Atoi(strings.TrimSpace(rangeParts[1]))
	if err != nil {
		return nil, fmt.Errorf("invalid number in range: %s", rangeParts[1])
	}

	// Validate range
	if err := p.validateRange(start, end, total); err != nil {
		return nil, err
	}

	// Build range
	var nums []int
	for i := start; i <= end; i++ {
		nums = append(nums, i)
	}
	return nums, nil
}

// parseSingleNumber parses a single number
func (p *InteractivePresenter) parseSingleNumber(part string, total int) (int, error) {
	num, err := strconv.Atoi(part)
	if err != nil {
		return 0, fmt.Errorf("invalid number: %s", part)
	}

	if num < 1 || num > total {
		return 0, fmt.Errorf("number out of bounds: %d", num)
	}

	return num, nil
}

// validateRange validates range boundaries
func (p *InteractivePresenter) validateRange(start, end, total int) error {
	if start < 1 || start > total || end < 1 || end > total {
		return fmt.Errorf("range out of bounds: %d-%d", start, end)
	}

	if start > end {
		return fmt.Errorf("invalid range: start > end")
	}

	return nil
}