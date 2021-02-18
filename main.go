package main

import (
	"fmt"
	"learn-go-fiber/config"
	"learn-go-fiber/database"
	"learn-go-fiber/routes"
	"log"
	"os"

	"github.com/go-redis/redis/v7"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var client *redis.Client

func init() {
	//Initializing redis
	dsn := os.Getenv("REDIS_DSN")
	if len(dsn) == 0 {
		dsn = "localhost:6379"
	}
	client = redis.NewClient(&redis.Options{
		Addr: dsn, //redis port
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
}
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
