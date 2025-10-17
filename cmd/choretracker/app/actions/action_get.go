package actions

import (
	"context"
	"fmt"
	"os"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/delivery/parser"
	"github.com/Galdoba/choretracker/internal/delivery/tui"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/urfave/cli/v3"
)

func GetAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts := actx.GetService()

		req, err := parser.ParseCliCommand(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}

		got := &domain.Chore{}
		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case "":
			return fmt.Errorf("run mode not set\nuse '--run-mode' flag")
		case flags.VALUE_MODE_CLI:
			chr, err := ts.ServeRequest(req)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			got = chr
		case flags.VALUE_MODE_TUI:
			got, err := tui.SelectChore(ts)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			if id, ok := got.GetID(); ok {
				req.InjectID(id)
			}
			got, err = ts.ServeRequest(req)
			if err != nil {
				return utils.LogError(ts.Logger, "read chore failed", err)
			}
		default:
			return fmt.Errorf("unknown mode '%v'", mode)

		}

		fmt.Fprintf(os.Stdout, "%v", got.String())

		return err
	}
}
