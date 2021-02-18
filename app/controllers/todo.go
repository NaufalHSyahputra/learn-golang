package controllers

import (
	"learn-go-fiber/app/models"
	request "learn-go-fiber/app/requests"
	"learn-go-fiber/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
)

type TodoResponse struct {
	ID uint `json:"id";gorm:"primarykey"`
	// UserID      uint   `json:"user_id"`
	Name        string `json:"name"`
	IsCompleted bool   `json:"is_completed"`
	UserName    string `json:"user_name"`
}

func GetTodo(c *fiber.Ctx) error {
	db := config.GetDBInstance()
	// todos := []models.Todo{}
	// if err := db.Find(&todos).Error; err != nil {
	// 	return c.Status(400).JSON(err)
	// }
	result := []TodoResponse{}
	todos := []models.Todo{}
	// db.Model(&todos).Find(&result)

	if err := db.Model(&todos).Select("todos.id", "todos.name", "user.name as user_name", "todos.user_id", "todos.is_completed").Joins("User").Where("user.id = ?", 1).Find(&result).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.JSON(result)
}

func GetSingleTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	db := config.GetDBInstance()
	todos := []models.Todo{}
	result := []TodoResponse{}
	if err := db.Model(todos).Select("todos.id", "todos.name", "user.name as user_name", "todos.user_id", "todos.is_completed").Joins("User").Where("user.id = ?", 1).Find(&result, id).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.JSON(result)
}

func InsertTodo(c *fiber.Ctx) error {

	// db := config.GetDBInstance()
	// result := db.Create(&newTodo)
	p := new(request.TodoForm)
	if err := c.BodyParser(p); err != nil {
		return err
	}
	v := validate.Struct(p)
	if !v.Validate() {
		return c.JSON(fiber.Map{
			"message": v.Errors.One(),
		})
	}
	newTodo := models.Todo{
		Name:        p.Name,
		IsCompleted: p.IsCompleted,
		UserID:      1,
	}
	db := config.GetDBInstance()
	if err := db.Create(&newTodo).Error; err != nil {
		return c.Status(400).JSON(err)
	}
	return c.JSON(fiber.Map{
		"message": "Create Todo Success!",
	})
}
