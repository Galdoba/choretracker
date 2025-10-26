package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Galdoba/choretracker/internal/core/dto"
	"github.com/Galdoba/choretracker/internal/core/services"
)

// GetAll - получает список всех Chore
func GetAll(ts *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		chores, err := ts.Storage.GetAll()
		if err != nil {
			ts.Logger.Errorf("Failed to get chores: %v", err)
			http.Error(w, fmt.Sprintf("Failed to get chores: %v", err), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		for _, chore := range chores {
			fmt.Fprintf(w, "%s\n\n", chore.String())
		}
	}
}

func Create(ts *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "Expect Method PUT", http.StatusMethodNotAllowed)
			return
		}
	}
}

// Хэндлер для получения конкретной chore по ID
func Get(ts *services.TaskService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
			ts.Logger.Errorf("Failed to get chore %d: %v", id, err)
			http.Error(w, fmt.Sprintf("Chore not found: %v", err), http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, chore.String())
	}

}
