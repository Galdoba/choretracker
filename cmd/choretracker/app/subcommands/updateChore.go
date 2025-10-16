package subcommands

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func UpdateChore(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:      constants.UpdateCommand,
		Usage:     "update chore data",
		UsageText: "choretracker [global options] update [options]",
		Version:   constants.Version,
		Flags: []cli.Flag{
			&flags.ID,
			&flags.TITLE,
			&flags.DESCRIPTION,
			&flags.AUTHOR,
			&flags.SHEDULE,
			&flags.COMMENT,
		},
		Action: actions.UpdateAction(actx),
	}
	return &add
}
