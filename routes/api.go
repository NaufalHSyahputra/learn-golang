package routes

import (
	"learn-go-fiber/app/controllers"
	"os"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
)

func ApiRoutes(app fiber.Router) {
	apiRoutes := app.Group("api")

	todoRoutes := apiRoutes.Group("todo")

	todoRoutes.Get("/", controllers.GetTodo)
	todoRoutes.Post("/create", controllers.InsertTodo)
	todoRoutes.Get("/:id", controllers.GetSingleTodo)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: os.Getenv("JWT_SECRET"),
	}))

	apiRoutes.Get("/restrict", func(c *fiber.Ctx) error {
		return c.SendString("Eh Kok Bisa Masuk?")
	})

}
