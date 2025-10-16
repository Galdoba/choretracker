package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/delivery"
	"github.com/urfave/cli/v3"
)

func DeleteAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts, err := startTaskService(actx)
		if err != nil {
			return fmt.Errorf("failed to start service: %v", err)
		}

		r, err := delivery.ParseCliArgsDelete(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}

		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case "":
			return fmt.Errorf("run mode not set\nuse '--run-mode' flag")
		case flags.VALUE_MODE_CLI:
			err := ts.DeleteTask(r)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
		case flags.VALUE_MODE_TUI:
			chr, err := getChoreTUI(ts.Storage)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			if r.ID == nil || *r.ID == 0 {
				r.ID = &chr.ID
			}
			err = ts.DeleteTask(r)
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
