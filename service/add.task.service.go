package service

import (
	"GST/billlingSystem/model"
	"errors"
	"log"
)

func (s *billerService) AddTask(task model.Task) error {

	if task.Title == "" {
		return errors.New("missing required fields")
	}

	err := s.store.InsertTask(task)
	if err != nil {
		log.Println("Error adding task:", err)
		return err
	}

	return nil
}
