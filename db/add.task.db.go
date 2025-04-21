package db

import (
	"GST/billlingSystem/model"
	"log"
)

func (s *pgStore) InsertTask(task model.Task) error {
	query := `
        INSERT INTO tasks (title, description, status, due_date)
        VALUES ($1, $2, $3, $4)
    `
	_, err := s.db.Exec(query,
		task.Title,
		task.Description,
		task.Status,
		task.DueDate,
	)
	if err != nil {
		log.Println("Error inserting task into DB:", err)
		return err
	}

	return nil
}
