package controllers

import (
	"learn-go-fiber/app/models"
	"learn-go-fiber/config"

	"github.com/gofiber/fiber/v2"
)

type TodoResponse struct {
	Id          int
	Name        string
	IsCompleted string
	UserId      int
	UserName    string
}

func GetTodo(c *fiber.Ctx) error {
	db := config.GetDBInstance()
	todos := []models.Todo{}
	if err := db.Preload("User").Find(&todos).Error; err != nil {
		return c.Status(400).JSON(err)
	}

	return c.JSON(todos)
}

func GetSingleTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.GetDBInstance()
	todos := []models.Todo{}
	if err := db.Find(&todos, id).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.JSON(todos)
}

func InsertTodo(c *fiber.Ctx) error {
	newTodo := models.Todo{
		Name:        "Learn Python",
		IsCompleted: false,
	}
	db := config.GetDBInstance()
	result := db.Create(&newTodo)
	return c.JSON(result.RowsAffected)
}
