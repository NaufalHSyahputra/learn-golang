package models

import (
	"learn-go-fiber/config"

	"gorm.io/gorm"
)

type Todo struct {
	Base
	UserID      uint `json:"user_id"`
	User        User
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
}

func (todo *Todo) CreateTodo() error {
	result := config.DB.Create(todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (todo *Todo) GetTodo(args ...string) *gorm.DB {
	query := config.DB.Model(todo).Select(args).Joins("User")
	return query
}

func (todo *Todo) GetTodoById(id int, args ...string) *gorm.DB {
	query := config.DB.Model(todo).Select(args).Joins("User")
	return query
}
