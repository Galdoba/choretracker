package subcommands

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func DeleteChore(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:      constants.DeleteCommand,
		Usage:     "delete chore",
		UsageText: "choretracker [global options] delete [options]",
		Version:   constants.Version,
		Flags: []cli.Flag{
			&flags.ID,
		},
		Action: actions.DeleteAction(actx),
	}
	return &add
}
