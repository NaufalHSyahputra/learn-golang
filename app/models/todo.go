package models

import "gorm.io/gorm"

type Todo struct {
	gorm.Model
	Id          int  `json:"id";gorm:"primaryKey"`
	UserID      uint `json:"user_id"`
	User        User
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
}
