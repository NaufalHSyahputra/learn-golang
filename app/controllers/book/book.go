package book

import (
	"github.com/gofiber/fiber/v2"
)

type BookStruct struct {
	Id     uint8
	Name   string
	Author string
}

func GetBook(c *fiber.Ctx) error {
	data := BookStruct{
		Id:     1,
		Name:   "Belajar Golang",
		Author: "Andi Budiman",
	}

	return c.JSON(data)
}
