package subcommands

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func Config(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:      constants.ConfigCommand,
		Usage:     "print current config file",
		UsageText: "choretracker config",
		Version:   constants.Version,
		Action:    actions.ConfigAction(actx),
	}
	return &add
}
