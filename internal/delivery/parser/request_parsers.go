package parser

import (
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/urfave/cli/v3"
)

// These functions processes a collection of CLI flags and extracts chore content information
// to construct specific type of Request data transfer objects.

// ParseCreateRequest transforms CLI flags into a CreateRequest DTO for chore creation.
//
// #AI_generated_comment
func ParseCreateRequest(flags []cli.Flag) dto.CreateRequest {
	return dto.CreateRequest{
		ChoreContent: parseContent(flags...),
	}
}

// ParseReadRequest transforms CLI flags into a ReadRequest DTO for chore retrieval.
//
// #AI_generated_comment
func ParseReadRequest(flags []cli.Flag) dto.ReadRequest {
	return dto.ReadRequest{
		ChoreIdentity: parseID(flags...),
	}
}

// ParseUpdateRequest transforms CLI flags into an UpdateRequest DTO for chore modification.
//
// #AI_generated_comment
func ParseUpdateRequest(flags []cli.Flag) dto.UpdateRequest {
	return dto.UpdateRequest{
		ChoreIdentity: parseID(flags...),
		ChoreContent:  parseContent(flags...),
	}
}

// ParseDeleteRequest transforms CLI flags into a DeleteRequest DTO for chore removal.
//
// #AI_generated_comment
func ParseDeleteRequest(flags []cli.Flag) dto.DeleteRequest {
	return dto.DeleteRequest{
		ChoreIdentity: parseID(flags...),
	}
}
