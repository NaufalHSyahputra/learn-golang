package database

import (
	"learn-go-fiber/app/models"

	"gorm.io/gorm"
)

func InitMigration(db *gorm.DB) {
	db.AutoMigrate(&models.User{}, &models.Todo{})
	InitSeeding(db)
}

func InitSeeding(db *gorm.DB) {
	var userCount, todoCount int64
	db.Model(&models.User{}).Count(&userCount)
	db.Model(&models.Todo{}).Count(&todoCount)
	if userCount == 0 {
		UserSeeder(db)
	}
	if todoCount == 0 {
		TodoSeeder(db)
	}
}

func UserSeeder(db *gorm.DB) {
	user := models.User{
		Name:     "Naufal",
		Email:    "naufal@test.com",
		Password: "$2y$12$0pFrZPFPTAJWw85GgIWXD.fEkYIcdJ//RBsGzC2EqglArI6W/Ugu2", //password
	}
	db.Create(&user)
}

func TodoSeeder(db *gorm.DB) {
	todos := []models.Todo{
		{UserID: 1, Name: "Belajar Golang", IsCompleted: false},
		{UserID: 1, Name: "Belajar Docker", IsCompleted: false},
		{UserID: 1, Name: "Belajar Kubernetes", IsCompleted: false},
		{UserID: 1, Name: "Belajar Flutter", IsCompleted: false},
	}
	db.Create(&todos)
}
