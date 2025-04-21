package handler

import (
	"GST/billlingSystem/model"
	"GST/billlingSystem/service"
	"encoding/json"
	"net/http"
	"time"
)

func AddTask(deps service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var task model.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		if task.Title == "" {
			http.Error(w, "Please provide task title", http.StatusBadRequest)
			return
		}

		// Set default status if not provided
		if task.Status == "" {
			task.Status = "Pending"
		}

		// Set CreatedAt and UpdatedAt timestamps
		now := time.Now()
		task.CreatedAt = now
		task.UpdatedAt = now

		// Save the task
		err = deps.BillerService.AddTask(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"message": "Task added successfully"})
	}
}
