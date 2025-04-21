package service

import (
	"GST/billlingSystem/model"
	"fmt"
)

func (s *billerService) GetTasks(page, limit int, status, dueDateAfter, dueDateBefore, sortBy, sortOrder string) ([]model.Task, error) {
	tasks, err := s.store.GetTasks(page, limit, status, dueDateAfter, dueDateBefore, sortBy, sortOrder)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch tasks: %w", err)
	}
	return tasks, nil
}
