package subcommands

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/constants"
	"github.com/Galdoba/choretracker/internal/models/chore"
	"github.com/urfave/cli/v3"
)

func AddChore(actx *appcontext.AppContext) *cli.Command {
	add := cli.Command{
		Name:                            constants.AddCommand,
		Aliases:                         []string{},
		Usage:                           "add new chore",
		UsageText:                       "",
		ArgsUsage:                       "",
		Version:                         constants.Version,
		Description:                     "",
		DefaultCommand:                  "",
		Category:                        "",
		Commands:                        []*cli.Command{},
		Flags:                           []cli.Flag{},
		HideHelp:                        false,
		HideHelpCommand:                 false,
		HideVersion:                     false,
		EnableShellCompletion:           false,
		ShellCompletionCommandName:      "",
		ShellComplete:                   nil,
		ConfigureShellCompletionCommand: nil,
		Before:                          nil,
		After:                           nil,
		Action: func(ctx context.Context, c *cli.Command) error {
			ch, err := chore.NewChore()
			if err != nil {
				return fmt.Errorf("failed to create chore")
			}
			return ch.Edit()
			// return nil
		},
		CommandNotFound:          nil,
		OnUsageError:             nil,
		InvalidFlagAccessHandler: nil,
		Hidden:                   false,
		Authors:                  []any{},
		Copyright:                "",
		Reader:                   nil,
		Writer:                   nil,
		ErrWriter:                nil,
		ExitErrHandler:           nil,
		Metadata:                 map[string]interface{}{},
		ExtraInfo: func() map[string]string {
			panic("TODO")
		},
		CustomRootCommandHelpTemplate: "",
		SliceFlagSeparator:            "",
		DisableSliceFlagSeparator:     false,
		UseShortOptionHandling:        false,
		Suggest:                       false,
		AllowExtFlags:                 false,
		SkipFlagParsing:               false,
		CustomHelpTemplate:            "",
		PrefixMatchCommands:           false,
		SuggestCommandFunc:            nil,
		MutuallyExclusiveFlags:        []cli.MutuallyExclusiveFlags{},
		Arguments:                     []cli.Argument{},
		ReadArgsFromStdin:             false,
	}
	return &add
}
