package actions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/delivery/handlers"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/urfave/cli/v3"
)

func ServeAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		port := c.Int(flags.SERVER_PORT)
		ts := actx.GetService()
		logger := actx.GetLogger()

		// Хэндлер для получения всех chores
		http.HandleFunc("/chores", handlers.GetAll(ts))

		http.HandleFunc("/chores", handlers.Create(ts))

		// Хэндлер для получения конкретной chore по ID
		http.HandleFunc("/chores/", handlers.Get(ts))

		// Корневой хэндлер
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprint(w, "ChoreTracker Server\n\nAvailable endpoints:\n- GET /chores - list all chores\n- GET /chores/{id} - get specific chore\n")
		})

		serverAddr := fmt.Sprintf(":%d", port)
		logger.Infof("Starting server on http://localhost%s", serverAddr)

		if err := http.ListenAndServe(serverAddr, nil); err != nil {
			return utils.LogError(logger, "Server failed", err)
		}

		return nil
	}
}
