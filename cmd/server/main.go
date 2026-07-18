package main

import (
	"log"
	centralroutes "navora/packages/central_routes"
	"navora/packages/config"
	"navora/packages/redis"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg := config.LeadConfig()

	//fiber engine
	app := fiber.New(fiber.Config{
		BodyLimit: 20 * 1024 * 1024, //set the limit to 20mb for image uploading
	})

	_, err := redis.NewRedisClient(cfg)
	if err != nil {
		log.Fatal("Redis connection failed:", err)
	}

	centralroutes.SetUp(app)

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	log.Println("Server running on port", port)

	log.Fatal(app.Listen(":" + port))
}