package actions

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/Galdoba/choretracker/cmd/choretracker/app/flags"
	"github.com/Galdoba/choretracker/internal/appcontext"
	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/internal/utils"
	"github.com/urfave/cli/v3"
)

func ServeAction(actx *appcontext.AppContext) cli.ActionFunc {
	return func(ctx context.Context, c *cli.Command) error {
		port := c.Int(flags.SERVER_PORT)
		ts := actx.GetService()
		logger := actx.GetLogger()

		// Хэндлер для получения всех chores
		http.HandleFunc("/chores", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			chores, err := ts.Storage.GetAll()
			if err != nil {
				logger.Errorf("Failed to get chores: %v", err)
				http.Error(w, fmt.Sprintf("Failed to get chores: %v", err), http.StatusInternalServerError)
				return
			}

			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			for _, chore := range chores {
				fmt.Fprintf(w, "%s\n\n", chore.String())
			}
		})

		// Хэндлер для получения конкретной chore по ID
		http.HandleFunc("/chores/", func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet {
				http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
				return
			}

			// Извлекаем ID из URL
			idStr := r.URL.Path[len("/chores/"):]
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				http.Error(w, "Invalid chore ID", http.StatusBadRequest)
				return
			}

			req := dto.ToServiceRequest{
				Action: dto.Read,
				Identity: dto.ChoreIdentity{
					ID: &id,
				},
			}

			chore, err := ts.ServeRequest(req)
			if err != nil {
				logger.Errorf("Failed to get chore %d: %v", id, err)
				http.Error(w, fmt.Sprintf("Chore not found: %v", err), http.StatusNotFound)
				return
			}

			w.Header().Set("Content-Type", "text/plain; charset=utf-8")
			fmt.Fprint(w, chore.String())
		})

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
