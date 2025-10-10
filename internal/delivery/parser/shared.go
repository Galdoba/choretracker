package parser

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/urfave/cli/v3"
)

// parseContent extracts chore content information from CLI flags and constructs a ChoreContent DTO.
//
// This function iterates through provided CLI flags, identifying those that correspond to
// chore content fields (title, description, author, schedule, comment) and populates a
// ChoreContent data transfer object. It handles type conversion and empty value filtering
// to ensure only valid, non-empty values are included in the result.
//
// Behavior and side effects:
//   - Processes flags in the order provided, with later flags overwriting earlier ones
//   - Empty string values are converted to nil pointers to indicate absence of value
//   - Only string-type flags with known names are processed; others are ignored
//
// Notes:
//   - This function is package-private and intended for internal use within the parser package
//   - Performance is O(n*m) where n is flag count and m is field name count
//   - Related functions: parseID, ptrStringOf
//
// #AI_generated_comment
func parseContent(flagsProvided ...cli.Flag) dto.ChoreContent {
	c := dto.ChoreContent{}
	for _, flag := range flagsProvided {
		names := flag.Names()
		for _, name := range names {
			switch name {
			case flags.CHORE_TITLE:
				c.Title = ptrStringOf(flag.Get())
			case flags.CHORE_DESCRIPTION:
				c.Description = ptrStringOf(flag.Get())
			case flags.CHORE_AUTHOR:
				c.Author = ptrStringOf(flag.Get())
			case flags.CHORE_SCHEDULE:
				c.Schedule = ptrStringOf(flag.Get())
			case flags.CHORE_COMMENT:
				c.Comment = ptrStringOf(flag.Get())
			}
		}
	}
	return c
}

// ptrStringOf converts an interface value to a string pointer, handling empty strings appropriately.
//
// #AI_generated_comment
func ptrStringOf(val any) *string {
	switch v := val.(type) {
	case string:
		if v != "" {
			return &v
		}
	}
	return nil
}

// parseID extracts chore identification information from CLI flags and constructs a ChoreIdentity DTO.
//
// This function scans through provided CLI flags to find the chore ID field, which is used
// to uniquely identify chore records in the system. It converts the flag value to the
// appropriate type and handles zero-value filtering to ensure only valid IDs are included.
//
// Behavior and side effects:
//   - Only processes the first valid ID flag encountered; subsequent ID flags may overwrite
//   - Zero ID values are converted to nil pointers
//   - Non-int64 values for ID flags are treated as zero values
//
// Notes:
//   - This function is package-private and intended for internal use within the parser package
//   - For bulk operations, consider specialized batch parsing functions
//   - Related functions: parseContent, ptrInt64Of
//
// #AI_generated_comment
func parseID(flagsProvided ...cli.Flag) dto.ChoreIdentity {
	id := dto.ChoreIdentity{}
	for _, flag := range flagsProvided {
		names := flag.Names()
		for _, name := range names {
			switch name {
			case flags.CHORE_ID:
				id.ID = ptrInt64Of(flag.Get())
			}
		}
	}
	return id
}

// ptrInt64Of converts an interface value to an int64 pointer, handling zero values appropriately.
//
// #AI_generated_comment
func ptrInt64Of(val any) *int64 {
	switch v := val.(type) {
	case int64:
		if v != 0 {
			return &v
		}
	}
	return nil
}
