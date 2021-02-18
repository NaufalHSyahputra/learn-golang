package routes

import (
	"learn-go-fiber/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app fiber.Router) {
	apiRoutes := app.Group("api")

	todoRoutes := apiRoutes.Group("todo")

	todoRoutes.Get("/", controllers.GetTodo)
	todoRoutes.Post("/create", controllers.InsertTodo)
	todoRoutes.Get("/:id", controllers.GetSingleTodo)

}
