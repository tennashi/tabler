package parser

import (
	"strings"
	"time"
)

const (
	tagPrefix = "#"
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