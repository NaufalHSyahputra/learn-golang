package routes

import (
	"learn-go-fiber/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func ApiRoutes(app fiber.Router) {
	apiRoutes := app.Group("api")

	todoRoutes := apiRoutes.Group("todo")
	// userRoutes := apiRoutes.Group("user")

	// app.Use(jwtware.New(jwtware.Config{
	// 	SigningKey: os.Getenv("JWT_SECRET"),
	// }))

	todoRoutes.Get("/", controllers.GetTodo)
	todoRoutes.Post("/create", controllers.InsertTodo)
	todoRoutes.Get("/get/:id", controllers.GetSingleTodo)

	apiRoutes.Get("/restrict", func(c *fiber.Ctx) error {
		return c.SendString("Eh Kok Bisa Masuk?")
	})

}
