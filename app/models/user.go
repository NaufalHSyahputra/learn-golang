package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       int    `json:id;gorm:primaryKey`
	Email    string `json:email`
	Password string `json:password;gorm:"->:false;<-:create"`
	Name     string `json:name`
	Todos    []Todo
}
