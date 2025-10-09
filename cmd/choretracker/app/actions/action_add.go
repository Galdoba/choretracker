package actions

import (
	"context"
	"fmt"

	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/services"
	"github.com/Galdoba/choretracker/internal/delivery"
	"github.com/Galdoba/choretracker/internal/infrastructure"
	"github.com/Galdoba/choretracker/internal/infrastructure/storage"
	"github.com/urfave/cli/v3"
)

func AddAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		cfg := actx.Config()
		logger := actx.GetLogger()
		validator := infrastructure.DefaultValidator()
		store, err := storage.NewStorage(storage.JsonStorage, cfg)
		if err != nil {
			logger.Errorf("failed to setup storage")
			return fmt.Errorf("failed to setup storage: %v", err)
		}
		ts := services.NewTaskService(store, validator, logger)

		r, err := delivery.ParseCliArgsCreate(c)
		if err != nil {
			ts.Logger.Errorf("failed to parse request: %v", err)
			return fmt.Errorf("failed to parse request: %v", err)
		}

		if err := ts.CreateTask(r); err != nil {
			ts.Logger.Errorf("task failed: %v", err)
			return err
		}

		return err
	}
}
