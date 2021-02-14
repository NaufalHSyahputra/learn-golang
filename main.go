package main

import (
	"fmt"
	"learn-go-fiber/config"
	"learn-go-fiber/database"
	"learn-go-fiber/routes"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	//env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	//Database
	config.NewDB()
	db := config.DB
	// sqlDB := db.DB()
	database.InitMigration(db)
	//Fiber
	app := fiber.New()
	routes.ApiRoutes(app)
	app.Listen(fmt.Sprintf(":%v", os.Getenv("APP_PORT")))
}
