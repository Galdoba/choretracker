package subcommands

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func GetChore(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:      constants.GetCommand,
		Usage:     "get chore data",
		UsageText: "choretracler [global options] get [options]",
		Version:   constants.Version,
		Flags: []cli.Flag{
			&flags.ID,
		},
		Action: actions.GetAction(actx),
	}
	return &add
}
