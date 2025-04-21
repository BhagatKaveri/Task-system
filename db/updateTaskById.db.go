package db

import (
	"GST/billlingSystem/model"
	"database/sql"
	"log"
)

func (s *pgStore) UpdateTaskByID(id int, input model.UpdateTaskInput) (*model.Task, error) {
	query := `
		UPDATE tasks
		SET title = $1, description = $2, status = $3, due_date = $4, updated_at = CURRENT_TIMESTAMP
		WHERE id = $5
		RETURNING id, title, description, status, due_date, created_at, updated_at
	`

	row := s.db.QueryRow(query, input.Title, input.Description, input.Status, input.DueDate, id)

	var task model.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Not found
		}
		log.Println("Error updating task:", err)
		return nil, err
	}

	return &task, nil
}
