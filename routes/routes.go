package routes

import (
    "github.com/gofiber/fiber/v2"
    "github.com/tassanai55555/media-city-backend/controllers"
)

func SetupRoutes(app *fiber.App) {
    app.Post("/users", controllers.CreateUser)
}
