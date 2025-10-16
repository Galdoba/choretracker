package subcommands

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func AddChore(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:      constants.AddCommand,
		Aliases:   []string{},
		Usage:     "add new chore",
		UsageText: "choretracler [global options] add [options]",
		Version:   constants.Version,
		Flags: []cli.Flag{
			&flags.TITLE,
			&flags.DESCRIPTION,
			&flags.AUTHOR,
			&flags.SHEDULE,
			&flags.COMMENT,
		},
		Action: actions.AddAction(actx),
	}
	return &add
}
