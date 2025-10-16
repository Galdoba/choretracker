package parser

import (
	"fmt"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/urfave/cli/v3"
)

func ParseCliArgsCreate(c *cli.Command) (dto.CreateRequest, error) {
	r := dto.CreateRequest{}
	switch c.Name {
	case constants.AddCommand:
		r.ChoreContent = parseContent(c)
	default:
		return r, fmt.Errorf("command '%v' does not use CreateRequest", c.Name)
	}
	return r, nil
}

func ParseCliArgsRead(c *cli.Command) (dto.ReadRequest, error) {
	r := dto.ReadRequest{}
	switch c.Name {
	case constants.GetCommand:
		r.ChoreIdentity = parseID(c)
	default:
		return r, fmt.Errorf("command '%v' does not use ReadRequest", c.Name)
	}
	return r, nil
}

func ParseCliArgsUpdate(c *cli.Command) (dto.UpdateRequest, error) {
	r := dto.UpdateRequest{}
	switch c.Name {
	case constants.UpdateCommand:
		r.ChoreIdentity = parseID(c)
		r.ChoreContent = parseContent(c)
	default:
		return r, fmt.Errorf("command '%v' does not use UpdateRequest", c.Name)
	}
	return r, nil
}

func ParseCliArgsDelete(c *cli.Command) (dto.DeleteRequest, error) {
	r := dto.DeleteRequest{}
	switch c.Name {
	case constants.DeleteCommand:
		r.ChoreIdentity = parseID(c)
	default:
		return r, fmt.Errorf("command '%v' does not use DeleteRequest", c.Name)
	}
	return r, nil
}

func parseContent(c *cli.Command) dto.ChoreContent {
	cntnt := dto.ChoreContent{}
	for i, val := range []string{
		c.String(flags.CHORE_TITLE),
		c.String(flags.CHORE_DESCRIPTION),
		c.String(flags.CHORE_AUTHOR),
		c.String(flags.CHORE_SCHEDULE),
		c.String(flags.CHORE_COMMENT),
	} {
		switch i {
		case 0:
			cntnt.Title = &val
		case 1:
			cntnt.Description = &val
		case 2:
			cntnt.Author = &val
		case 3:
			cntnt.Schedule = &val
		case 4:
			cntnt.Comment = &val
		}
	}
	return cntnt
}

func parseID(c *cli.Command) dto.ChoreIdentity {
	id := dto.ChoreIdentity{}
	if val := c.Int64(flags.CHORE_ID); val != 0 {
		id.ID = &val
	}
	return id
}
