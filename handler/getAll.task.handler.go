package handler

import (
	"GST/billlingSystem/service"
	"encoding/json"
	"net/http"
	"strconv"
)

func GetAllTasks(deps service.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query()

		pageStr := query.Get("page")
		limitStr := query.Get("limit")
		status := query.Get("status")
		dueDateAfter := query.Get("due_date_after")
		dueDateBefore := query.Get("due_date_before")
		sortBy := query.Get("sort_by")
		sortOrder := query.Get("sort_order")

		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = 1
		}

		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 1 || limit > 50 {
			limit = 10
		}

		tasks, err := deps.BillerService.GetTasks(page, limit, status, dueDateAfter, dueDateBefore, sortBy, sortOrder)
		if err != nil {
			http.Error(w, "Failed to fetch tasks", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}
