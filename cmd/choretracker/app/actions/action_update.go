package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/delivery/parser"
	"github.com/Galdoba/choretracker/internal/delivery/tui"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/urfave/cli/v3"
)

func UpdateAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts := actx.GetService()
		r, err := parser.ParseCliCommand(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}
		// got := &domain.Chore{}
		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case "":
			return fmt.Errorf("run mode not set\nuse '--run-mode' flag")
		case flags.VALUE_MODE_CLI:
			_, err = ts.ServeRequest(r)
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
				r.InjectID(id)
				r.InjectContent(chr.Content())
			}

			if err := tui.EditRequest(&r); err != nil {
				return utils.LogError(ts.Logger, "failed to edit request", err)
			}
			if _, err := ts.ServeRequest(r); err != nil {
				// got = updated
				return utils.LogError(ts.Logger, "failed to update chore", err)
			}
		default:
			return fmt.Errorf("unknown mode '%v'", mode)
		}

		return err
	}
}
