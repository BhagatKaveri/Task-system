package model

import (
	"time"
)

type User struct {
	UserID   string `db:"user_id" json:"user_id"`
	Email    string `db:"username" json:"email"`
	Password string `db:"password_hash" json:"password"`
}

type LoginResponse struct {
	Message string `json:"msg"`
	Token   string `json:"token,omitempty"`
}

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type TaskSearchInput struct {
	UserID  string `json:"user_id,omitempty"`
	Status  string `json:"status,omitempty"`
	DueDate string `json:"due_date,omitempty"`
}
type UpdateTaskInput struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	DueDate     *time.Time `json:"due_date,omitempty"`
}
