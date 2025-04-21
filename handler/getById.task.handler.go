package handler

import (
	"GST/billlingSystem/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetTaskByID(deps service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get task ID from URL
		vars := mux.Vars(r) // Using gorilla/mux
		idStr := vars["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		task, err := deps.BillerService.GetTaskByID(id)
		if err != nil {
			http.Error(w, "Failed to fetch task", http.StatusInternalServerError)
			return
		}
		if task == nil {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(task)
	}
}
