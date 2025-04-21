package db

import (
	"GST/billlingSystem/model"
	"database/sql"
	"log"
)

func (s *pgStore) GetTaskByID(id int) (*model.Task, error) {
	query := `
		SELECT id, title, description, status, due_date, created_at, updated_at
		FROM tasks
		WHERE id = $1
	`
	row := s.db.QueryRow(query, id)

	var task model.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description, &task.Status, &task.DueDate, &task.CreatedAt, &task.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Println("Error fetching task by ID:", err)
		return nil, err
	}

	return &task, nil
}
