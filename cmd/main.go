package main

import (
	"ecommerce-backend/database"
	"ecommerce-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// ✅ Environment variables are automatically loaded in Docker via `env_file`
	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}