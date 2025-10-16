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

func AddAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts, err := startTaskService(actx)
		if err != nil {
			return fmt.Errorf("failed to start service: %v", err)
		}

		r, err := parser.ParseCliArgsCreate(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}
		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case flags.VALUE_MODE_CLI:
			//do nothing: expected to have data from flags --title and --shedule
		case flags.VALUE_MODE_TUI:
			r, err = tui.EditCreateRequest(r)
			if err != nil {
				return utils.LogError(ts.Logger, "failed to edit request", err)
			}
		case flags.VALUE_MODE_SERVER:
			return fmt.Errorf("server mode not implemented")
		}

		if err := ts.CreateTask(r); err != nil {
			ts.Logger.Errorf("task failed: %v", err)
			return err
		}
		return err
	}
}
