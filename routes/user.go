package routes

import (
	"github.com/NaufalHSyahputra/learn-golang/app/controllers/book"

	"github.com/gofiber/fiber/v2"
)

func ApiUser(app fiber.Router) {
	userRoute := app.Group("user")

	userRoute.Get("/", book.GetBook)
}
