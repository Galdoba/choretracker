package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/internal/delivery"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/urfave/cli/v3"
)

func UpdateAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts, err := startTaskService(actx)
		if err != nil {
			return fmt.Errorf("failed to start service: %v", err)
		}

		r, err := delivery.ParseCliArgsUpdate(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}

		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case "":
			return fmt.Errorf("run mode not set\nuse '--run-mode' flag")
		case flags.VALUE_MODE_CLI:
			err := ts.UpdateTask(r)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
		case flags.VALUE_MODE_TUI:
			chr, err := getChoreTUI(ts.Storage)
			// chr, err := ts.ReadTask(r)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			r = injectEmptyFields(r, chr)
			r, err = delivery.EditUpdateRequest(r)
			if err != nil {
				return utils.LogError(ts.Logger, "failed to edit request", err)
			}
			if err := ts.UpdateTask(r); err != nil {
				return utils.LogError(ts.Logger, "failed to update chore", err)
			}
		default:
			return fmt.Errorf("unknown mode '%v'", mode)

		}

		return err
	}
}

func injectEmptyFields(r dto.UpdateRequest, original domain.Chore) dto.UpdateRequest {
	if r.ID == nil || *r.ID == 0 {
		r.ID = &original.ID
	}
	if r.Title == nil || *r.Title == "" {
		r.Title = &original.Title
	}
	if r.Description == nil || *r.Description == "" {
		r.Description = &original.Description
	}
	if r.Author == nil || *r.Author == "" {
		r.Author = &original.Author
	}
	if r.Schedule == nil || *r.Schedule == "" {
		r.Schedule = &original.Schedule
	}
	if r.Comment == nil || *r.Comment == "" {
		r.Comment = &original.Comment
	}
	return r
}
