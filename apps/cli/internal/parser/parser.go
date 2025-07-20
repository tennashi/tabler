package parser

import "time"

type ParseResult struct {
	Title    string
	Tags     []string
	Priority int
	Deadline *time.Time
}

func Parse(input string) *ParseResult {
	return &ParseResult{
		Title: input,
		Tags:  []string{},
	}
}