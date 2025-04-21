package service

import (
	"GST/billlingSystem/model"
	"fmt"
)

func (s *billerService) UpdateTaskByID(id int, input model.UpdateTaskInput) (*model.Task, error) {
	task, err := s.store.UpdateTaskByID(id, input)
	if err != nil {
		return nil, fmt.Errorf("failed to update task: %w", err)
	}
	return task, nil
}
