package main

import (
	"ecommerce-backend/database"
	"ecommerce-backend/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	log.Fatal(app.Listen(":3000"))
}