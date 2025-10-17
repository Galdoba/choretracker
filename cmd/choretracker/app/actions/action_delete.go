package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/delivery/parser"
	"github.com/Galdoba/choretracker/internal/delivery/tui"
	"github.com/urfave/cli/v3"
)

func DeleteAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts := actx.GetService()
		req, err := parser.ParseCliCommand(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}

		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case "":
			return fmt.Errorf("run mode not set\nuse '--run-mode' flag")
		case flags.VALUE_MODE_CLI:
			_, err := ts.ServeRequest(req)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
		case flags.VALUE_MODE_TUI:
			chr, err := tui.SelectChore(ts)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			if id, ok := chr.GetID(); ok {
				req.InjectID(id)
			}
			_, err = ts.ServeRequest(req)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
		default:
			return fmt.Errorf("unknown mode '%v'", mode)

		}

		return err
	}
}
