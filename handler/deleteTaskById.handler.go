package handler

import (
	"GST/billlingSystem/service"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func DeleteTaskByID(deps service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr := vars["id"]

		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid task ID", http.StatusBadRequest)
			return
		}

		deleted, err := deps.BillerService.DeleteTaskByID(id)
		if err != nil {
			http.Error(w, "Failed to delete task", http.StatusInternalServerError)
			return
		}
		if !deleted {
			http.Error(w, "Task not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message":"Task successfully deleted"}`))
	}
}
