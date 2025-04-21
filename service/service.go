package service

import (
	"GST/billlingSystem/db"
	"GST/billlingSystem/model"
)

type Service interface {
	Login(model.User) (string, error)
	AddTask(model.Task) error
	GetTasks(page, limit int, status, dueDateAfter, dueDateBefore, sortBy, sortOrder string) ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTaskByID(id int, input model.UpdateTaskInput) (*model.Task, error)
	DeleteTaskByID(id int) (bool, error)
}

type billerService struct {
	store db.Storer
}

func NewBillerService(s db.Storer) Service {
	return &billerService{
		store: s,
	}

}
