package db

import (
	"GST/billlingSystem/model"
)

type Storer interface {
	Login(model.User) (bool, error)
	InsertTask(model.Task) error
	GetTasks(page, limit int, status, dueDateAfter, dueDateBefore, sortBy, sortOrder string) ([]model.Task, error)
	GetTaskByID(id int) (*model.Task, error)
	UpdateTaskByID(id int, input model.UpdateTaskInput) (*model.Task, error)
	DeleteTaskByID(id int) (bool, error)
}
