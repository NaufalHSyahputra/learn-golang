package main

import (
	"fmt"
	"log"
	"os"

	"github.com/NaufalHSyahputra/learn-golang/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.ApiRoutes(app)

	app.Listen(fmt.Sprintf(":%v", os.Getenv("PORT")))
}
