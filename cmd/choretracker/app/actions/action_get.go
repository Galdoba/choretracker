package actions

import (
	"context"
	"fmt"
	"os"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/domain"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/internal/core/ports"
	"github.com/Galdoba/choretracker/internal/core/services"
	"github.com/Galdoba/choretracker/internal/delivery"
	"github.com/Galdoba/choretracker/internal/infrastructure"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/Galdoba/consolio/prompt"
	"github.com/urfave/cli/v3"
)

func GetAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		ts, err := startTaskService(actx)
		if err != nil {
			return fmt.Errorf("failed to start service: %v", err)
		}

		r, err := delivery.ParseCliArgsRead(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}

		got := domain.Chore{}
		mode := c.String(flags.GLOBAL_MODE)
		switch mode {
		case "":
			return fmt.Errorf("run mode not set\nuse '--run-mode' flag")
		case flags.VALUE_MODE_CLI:
			chr, err := ts.ReadTask(r)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			got = chr
		case flags.VALUE_MODE_TUI:
			chr, err := getChoreTUI(ts.Storage)
			// chr, err := ts.ReadTask(r)
			if err != nil {
				ts.Logger.Errorf("task failed: %v", err)
				return err
			}
			got = chr
			r = injectID(r, got)
			got, err = ts.ReadTask(r)
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

func startTaskService(actx *appcontext.AppContext) (*services.TaskService, error) {
	cfg := actx.Config()
	logger := actx.GetLogger()
	validator := infrastructure.DefaultValidator()
	store, err := storage.NewStorage(storage.JsonStorage, cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to setup storage: %v", err)
	}
	ts := services.NewTaskService(store, validator, logger)
	return ts, nil
}

// getChoreTUI return chore from storage in interactive mode
func getChoreTUI(store ports.Storage) (domain.Chore, error) {
	chr := domain.Chore{}
	chores, err := store.GetAll()
	if err != nil {
		return chr, fmt.Errorf("failed to get chore list from storage: %v", err)
	}
	// ids := []int64{}
	selectionPool := []*prompt.Item{}
	for _, ch := range chores {
		selectionPool = append(selectionPool, prompt.NewItem(fmt.Sprintf("id:%v (%v)", ch.ID, ch.Title), ch))
	}
	ch, err := prompt.SearchItem(prompt.WithTitle("search chores:"), prompt.FromItems(selectionPool))
	if err != nil {
		return chr, err
	}
	chr = ch.Payload().(domain.Chore)
	return chr, nil
}

func injectID(r dto.ReadRequest, original domain.Chore) dto.ReadRequest {
	if r.ID == nil || *r.ID == 0 {
		r.ID = &original.ID
	}
	return r
}
