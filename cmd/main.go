package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/joho/godotenv"
    "github.com/tassanai55555/media-city-backend/config"
    "github.com/tassanai55555/media-city-backend/routes"
)

func main() {
    godotenv.Load()
    config.ConnectDatabase()

    app := fiber.New()

    routes.SetupRoutes(app)

    app.Listen(":3000")
}
