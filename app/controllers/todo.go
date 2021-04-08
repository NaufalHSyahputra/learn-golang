package controllers

import (
	"learn-go-fiber/app/models"
	request "learn-go-fiber/app/requests"
	"strconv"

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
	result := []TodoResponse{}
	todos := models.Todo{}
	if err := todos.GetTodo("todos.id", "todos.name", "user.name as user_name", "todos.user_id", "todos.is_completed").Find(&result); err.Error != nil {
		return c.Status(400).JSON(err.Error)
	}
	return c.JSON(result)
}

func GetSingleTodo(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(400).JSON(err)
	}
	result := []TodoResponse{}
	todos := models.Todo{}
	if err := todos.GetTodoById(id, "todos.id", "todos.name", "user.name as user_name", "todos.user_id", "todos.is_completed").Find(&result, id); err.Error != nil {
		return c.Status(400).JSON(err.Error)
	}
	return c.JSON(result)
}

func InsertTodo(c *fiber.Ctx) error {
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

	err := newTodo.CreateTodo()
	if err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}
	return c.JSON(fiber.Map{
		"message": "Create Todo Success!",
	})
}
