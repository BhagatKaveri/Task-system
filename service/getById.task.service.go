package service

import (
	"GST/billlingSystem/model"
	"fmt"
)

func (s *billerService) GetTaskByID(id int) (*model.Task, error) {
	task, err := s.store.GetTaskByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch task by id: %w", err)
	}
	return task, nil
}
