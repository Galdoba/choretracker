package parser

import (
	"testing"

	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/urfave/cli/v3"
)

func TestParseCreateRequest(t *testing.T) {
	tests := []struct {
		name     string
		flags    []cli.Flag
		expected dto.CreateRequest
	}{
		{
			name:     "empty flags",
			flags:    []cli.Flag{},
			expected: dto.CreateRequest{},
		},
		{
			name: "with content flags",
			flags: []cli.Flag{
				&cli.StringFlag{Name: "title", Value: "Test Chore"},
				&cli.StringFlag{Name: "description", Value: "Test Description"},
			},
			expected: dto.CreateRequest{
				ChoreContent: dto.ChoreContent{
					Title:       stringPtr("Test Chore"),
					Description: stringPtr("Test Description"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseCreateRequest(tt.flags)
			if !compareCreateRequest(result, tt.expected) {
				t.Errorf("ParseCreateRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseReadRequest(t *testing.T) {
	tests := []struct {
		name     string
		flags    []cli.Flag
		expected dto.ReadRequest
	}{
		{
			name:     "empty flags",
			flags:    []cli.Flag{},
			expected: dto.ReadRequest{},
		},
		{
			name: "with ID flag",
			flags: []cli.Flag{
				&cli.Int64Flag{Name: "id", Value: 123},
			},
			expected: dto.ReadRequest{
				ChoreIdentity: dto.ChoreIdentity{
					ID: int64Ptr(123),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseReadRequest(tt.flags)
			if !compareReadRequest(result, tt.expected) {
				t.Errorf("ParseReadRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseUpdateRequest(t *testing.T) {
	tests := []struct {
		name     string
		flags    []cli.Flag
		expected dto.UpdateRequest
	}{
		{
			name:     "empty flags",
			flags:    []cli.Flag{},
			expected: dto.UpdateRequest{},
		},
		{
			name: "with ID and content flags",
			flags: []cli.Flag{
				&cli.Int64Flag{Name: "id", Value: 123},
				&cli.StringFlag{Name: "title", Value: "Updated Title"},
				&cli.StringFlag{Name: "description", Value: "Updated Description"},
			},
			expected: dto.UpdateRequest{
				ChoreIdentity: dto.ChoreIdentity{
					ID: int64Ptr(123),
				},
				ChoreContent: dto.ChoreContent{
					Title:       stringPtr("Updated Title"),
					Description: stringPtr("Updated Description"),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseUpdateRequest(tt.flags)
			if !compareUpdateRequest(result, tt.expected) {
				t.Errorf("ParseUpdateRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseDeleteRequest(t *testing.T) {
	tests := []struct {
		name     string
		flags    []cli.Flag
		expected dto.DeleteRequest
	}{
		{
			name:     "empty flags",
			flags:    []cli.Flag{},
			expected: dto.DeleteRequest{},
		},
		{
			name: "with ID flag",
			flags: []cli.Flag{
				&cli.Int64Flag{Name: "id", Value: 123},
			},
			expected: dto.DeleteRequest{
				ChoreIdentity: dto.ChoreIdentity{
					ID: int64Ptr(123),
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ParseDeleteRequest(tt.flags)
			if !compareDeleteRequest(result, tt.expected) {
				t.Errorf("ParseDeleteRequest() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Helper functions for comparison
func compareCreateRequest(a, b dto.CreateRequest) bool {
	return compareChoreContent(a.ChoreContent, b.ChoreContent)
}

func compareReadRequest(a, b dto.ReadRequest) bool {
	return compareChoreIdentity(a.ChoreIdentity, b.ChoreIdentity)
}

func compareUpdateRequest(a, b dto.UpdateRequest) bool {
	return compareChoreIdentity(a.ChoreIdentity, b.ChoreIdentity) &&
		compareChoreContent(a.ChoreContent, b.ChoreContent)
}

func compareDeleteRequest(a, b dto.DeleteRequest) bool {
	return compareChoreIdentity(a.ChoreIdentity, b.ChoreIdentity)
}

func compareChoreContent(a, b dto.ChoreContent) bool {
	return compareStringPtrs(a.Title, b.Title) &&
		compareStringPtrs(a.Description, b.Description) &&
		compareStringPtrs(a.Author, b.Author) &&
		compareStringPtrs(a.Schedule, b.Schedule) &&
		compareStringPtrs(a.Comment, b.Comment)
}

func compareChoreIdentity(a, b dto.ChoreIdentity) bool {
	return compareInt64Ptrs(a.ID, b.ID)
}

func compareStringPtrs(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func compareInt64Ptrs(a, b *int64) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}

func stringPtr(s string) *string {
	return &s
}

func int64Ptr(i int64) *int64 {
	return &i
}