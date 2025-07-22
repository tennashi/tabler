package metadata

import (
	"context"
	"errors"
)

type Service struct {
	claude Claude
}

type Claude interface {
	ExtractMetadata(ctx context.Context, input string) (*ExtractedMetadata, error)
}

type ExtractedMetadata struct {
	CleanedText string
	Deadline    string
	Tags        []string
	Priority    string
}

func NewService(claude Claude) *Service {
	return &Service{claude: claude}
}

func (s *Service) Extract(ctx context.Context, input string) (*ExtractedMetadata, error) {
	if input == "" {
		return nil, errors.New("empty input")
	}

	// If we have a claude client, use it for extraction
	if s.claude != nil {
		return s.claude.ExtractMetadata(ctx, input)
	}

	// Fallback to simple extraction
	return &ExtractedMetadata{
		CleanedText: input,
	}, nil
}
