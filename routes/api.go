package routes

import (
	"github.com/NaufalHSyahputra/learn-golang/app/controllers/book"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app fiber.Router) {
	r := app.Group("api")

	ApiUser(r)

	bookRoutes := r.Group("book")

	bookRoutes.Get("/", book.GetBook)
}
