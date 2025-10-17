package subcommands

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func Serve(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:      constants.ServeCommand,
		Usage:     "start HTTP server",
		UsageText: "choretracker serve [options]",
		Version:   constants.Version,
		Flags: []cli.Flag{
			&flags.PORT,
		},
		Action: actions.ServeAction(actx),
	}
	return &add
}
