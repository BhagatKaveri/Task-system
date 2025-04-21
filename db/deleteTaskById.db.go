package db

import (
	"log"
)

func (s *pgStore) DeleteTaskByID(id int) (bool, error) {
	query := `DELETE FROM tasks WHERE id = $1`
	result, err := s.db.Exec(query, id)
	if err != nil {
		log.Println("Error deleting task:", err)
		return false, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error checking rows affected:", err)
		return false, err
	}

	if rowsAffected == 0 {
		return false, nil // No task found
	}

	return true, nil
}
