package parser

import (
	"strings"
	"time"
)

const (
	tagPrefix      = "#"
	priorityMarker = "!"
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