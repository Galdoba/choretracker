package app

import (
	"github.com/Galdoba/choretracker/cmd/choretracker/app/actions"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/cmd/choretracker/app/subcommands"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/urfave/cli/v3"
)

func NewApp(actx *appcontext.AppContext) *cli.Command {
	cmd := cli.Command{
		Name:        constants.AppName,
		Aliases:     []string{},
		Usage:       "enchanced todo app for learning",
		Version:     constants.Version,
		Description: "this app was writen to grasp DDD and hexogonal achitecture",
		Commands: []*cli.Command{
			subcommands.AddChore(actx),
			subcommands.GetChore(actx),
			subcommands.UpdateChore(actx),
			subcommands.DeleteChore(actx),
			subcommands.Config(actx),
		},
		Flags: []cli.Flag{
			&flags.MODE,
		},
		Action:  actions.Start(actx),
		Authors: []any{constants.Author},
	}
	return &cmd
}
