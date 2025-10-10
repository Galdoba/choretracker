package parser

import (
	"testing"

	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/urfave/cli/v3"
)

func TestParseContent(t *testing.T) {
	tests := []struct {
		name     string
		flags    []cli.Flag
		expected dto.ChoreContent
	}{
		{
			name:     "empty flags",
			flags:    []cli.Flag{},
			expected: dto.ChoreContent{},
		},
		{
			name: "all content fields",
			flags: []cli.Flag{
				&cli.StringFlag{Name: "title", Value: "Test Title"},
				&cli.StringFlag{Name: "description", Value: "Test Description"},
				&cli.StringFlag{Name: "author", Value: "Test Author"},
				&cli.StringFlag{Name: "schedule", Value: "daily"},
				&cli.StringFlag{Name: "comment", Value: "Test Comment"},
			},
			expected: dto.ChoreContent{
				Title:       stringPtr("Test Title"),
				Description: stringPtr("Test Description"),
				Author:      stringPtr("Test Author"),
				Schedule:    stringPtr("daily"),
				Comment:     stringPtr("Test Comment"),
			},
		},
		{
			name: "mixed flags with unknown",
			flags: []cli.Flag{
				&cli.StringFlag{Name: "title", Value: "Test Title"},
				&cli.StringFlag{Name: "unknown", Value: "Unknown Value"},
				&cli.StringFlag{Name: "description", Value: "Test Description"},
			},
			expected: dto.ChoreContent{
				Title:       stringPtr("Test Title"),
				Description: stringPtr("Test Description"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseContent(tt.flags...)
			if !compareChoreContent(result, tt.expected) {
				t.Errorf("parseContent() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseID(t *testing.T) {
	tests := []struct {
		name     string
		flags    []cli.Flag
		expected dto.ChoreIdentity
	}{
		{
			name:     "empty flags",
			flags:    []cli.Flag{},
			expected: dto.ChoreIdentity{},
		},
		{
			name: "with ID flag",
			flags: []cli.Flag{
				&cli.Int64Flag{Name: "id", Value: 123},
			},
			expected: dto.ChoreIdentity{
				ID: int64Ptr(123),
			},
		},
		{
			name: "with zero ID",
			flags: []cli.Flag{
				&cli.Int64Flag{Name: "id", Value: 0},
			},
			expected: dto.ChoreIdentity{},
		},
		{
			name: "mixed flags",
			flags: []cli.Flag{
				&cli.Int64Flag{Name: "id", Value: 456},
				&cli.StringFlag{Name: "title", Value: "Test Title"},
			},
			expected: dto.ChoreIdentity{
				ID: int64Ptr(456),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseID(tt.flags...)
			if !compareChoreIdentity(result, tt.expected) {
				t.Errorf("parseID() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPtrStringOf(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected *string
	}{
		{
			name:     "non-empty string",
			input:    "test string",
			expected: stringPtr("test string"),
		},
		{
			name:     "empty string",
			input:    "",
			expected: nil,
		},
		{
			name:     "non-string type",
			input:    123,
			expected: nil,
		},
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ptrStringOf(tt.input)
			if !compareStringPtrs(result, tt.expected) {
				t.Errorf("ptrStringOf(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestPtrInt64Of(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected *int64
	}{
		{
			name:     "non-zero int64",
			input:    int64(42),
			expected: int64Ptr(42),
		},
		{
			name:     "zero int64",
			input:    int64(0),
			expected: nil,
		},
		{
			name:     "non-int64 type",
			input:    "not int64",
			expected: nil,
		},
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ptrInt64Of(tt.input)
			if !compareInt64Ptrs(result, tt.expected) {
				t.Errorf("ptrInt64Of(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}