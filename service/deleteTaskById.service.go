package service

import (
	"fmt"
)

func (s *billerService) DeleteTaskByID(id int) (bool, error) {
	deleted, err := s.store.DeleteTaskByID(id)
	if err != nil {
		return false, fmt.Errorf("failed to delete task: %w", err)
	}
	return deleted, nil
}
