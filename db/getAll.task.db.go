package db

import (
	"GST/billlingSystem/model"
	"log"
)

// Helper function to validate allowed fields
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

func (s *pgStore) GetTasks(page, limit int, status, dueDateAfter, dueDateBefore, sortBy, sortOrder string) ([]model.Task, error) {
	allowedSortFields := []string{"id", "due_date"}
	allowedSortOrders := []string{"ASC", "DESC"}

	if !contains(allowedSortFields, sortBy) {
		sortBy = "id"
	}
	if !contains(allowedSortOrders, sortOrder) {
		sortOrder = "ASC"
	}

	query := `
		SELECT id, title, description, status, due_date, created_at, updated_at
		FROM tasks
		WHERE ($1::text IS NULL OR status = $1::task_status)
		AND ($2::text IS NULL OR due_date >= $2::timestamp)
		AND ($3::text IS NULL OR due_date <= $3::timestamp)
		ORDER BY ` + sortBy + ` ` + sortOrder + `
		LIMIT $4 OFFSET $5
	`

	rows, err := s.db.Query(query,
		status,
		dueDateAfter,
		dueDateBefore,
		limit,
		(page-1)*limit,
	)
	if err != nil {
		log.Println("Error fetching tasks:", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
		if err != nil {
			log.Println("Error scanning task:", err)
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}
