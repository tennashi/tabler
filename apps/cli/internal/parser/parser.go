package parser

import (
	"strings"
	"time"
)

const (
	tagPrefix      = "#"
	priorityMarker = "!"
	deadlinePrefix = "@"
)

type ParseResult struct {
	Title    string
	Tags     []string
	Priority int
	Deadline *time.Time
}

func Parse(input string) *ParseResult {
	result := &ParseResult{
		Tags: []string{},
	}

	parts := strings.Split(input, " ")
	titleParts := []string{}

	for _, part := range parts {
		if tag, isTag := extractTag(part); isTag {
			result.Tags = append(result.Tags, tag)
		} else if priority, isPriority := extractPriority(part); isPriority {
			result.Priority = priority
		} else if deadline, isDeadline := extractDeadline(part); isDeadline {
			result.Deadline = deadline
		} else {
			titleParts = append(titleParts, part)
		}
	}

	result.Title = strings.Join(titleParts, " ")

	return result
}

func extractTag(part string) (string, bool) {
	if strings.HasPrefix(part, tagPrefix) && len(part) > len(tagPrefix) {
		return part[len(tagPrefix):], true
	}
	return "", false
}

func extractPriority(part string) (int, bool) {
	if len(part) == 0 || part[0] != priorityMarker[0] {
		return 0, false
	}

	count := 0
	for _, ch := range part {
		if ch == rune(priorityMarker[0]) {
			count++
		} else {
			return 0, false
		}
	}

	if count >= 1 && count <= 3 {
		return count, true
	}
	return 0, false
}

func extractDeadline(part string) (*time.Time, bool) {
	if !strings.HasPrefix(part, deadlinePrefix) || len(part) <= len(deadlinePrefix) {
		return nil, false
	}

	dateStr := part[len(deadlinePrefix):]
	return parseDeadlineString(dateStr)
}

var weekdayMap = map[string]time.Weekday{
	"mon":       time.Monday,
	"monday":    time.Monday,
	"tue":       time.Tuesday,
	"tuesday":   time.Tuesday,
	"wed":       time.Wednesday,
	"wednesday": time.Wednesday,
	"thu":       time.Thursday,
	"thursday":  time.Thursday,
	"fri":       time.Friday,
	"friday":    time.Friday,
	"sat":       time.Saturday,
	"saturday":  time.Saturday,
	"sun":       time.Sunday,
	"sunday":    time.Sunday,
}

func parseDeadlineString(dateStr string) (*time.Time, bool) {
	switch dateStr {
	case "today":
		return todayDeadline(), true
	case "tomorrow":
		return tomorrowDeadline(), true
	default:
		if weekday, ok := weekdayMap[dateStr]; ok {
			return weekdayDeadline(weekday), true
		}
		// Try to parse as YYYY-MM-DD format
		if deadline, ok := parseSpecificDate(dateStr); ok {
			return deadline, true
		}
		return nil, false
	}
}

func todayDeadline() *time.Time {
	return dateAtStartOfDay(time.Now())
}

func tomorrowDeadline() *time.Time {
	return dateAtStartOfDay(time.Now().AddDate(0, 0, 1))
}

func weekdayDeadline(weekday time.Weekday) *time.Time {
	now := time.Now()
	daysUntilWeekday := (int(weekday) - int(now.Weekday()) + 7) % 7
	if daysUntilWeekday == 0 {
		daysUntilWeekday = 7
	}
	return dateAtStartOfDay(now.AddDate(0, 0, daysUntilWeekday))
}

func dateAtStartOfDay(date time.Time) *time.Time {
	deadline := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return &deadline
}

func parseSpecificDate(dateStr string) (*time.Time, bool) {
	parsedTime, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return nil, false
	}
	return dateAtStartOfDay(parsedTime), true
}
