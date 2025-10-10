package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/services"
	"github.com/Galdoba/choretracker/internal/delivery"
	"github.com/Galdoba/choretracker/internal/delivery/parser"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage"
	"github.com/urfave/cli/v3"
)

func AddAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		cfg := actx.Config()
		logger := actx.GetLogger()
		validator := actx.GetValidator()

		store, err := storage.NewStorage(storage.JsonStorage, cfg)
		if err != nil {
			logger.Errorf("failed to setup storage")
			return fmt.Errorf("failed to setup storage: %v", err)
		}
		ts := services.NewTaskService(store, validator, logger)

		request := parser.ParseCreateRequest(c.Flags)
		request, err = delivery.EditCreateRequest(request)
		if err != nil {
			logger.Errorf("failed to edit request")
			return fmt.Errorf("failed to edit request: %v", err)
		}

		if err := ts.CreateTask(request); err != nil {
			ts.Logger.Errorf("task failed: %v", err)
			return err
		}

		return err
	}
}
