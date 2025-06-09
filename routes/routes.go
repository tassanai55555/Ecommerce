package routes

import (
	"ecommerce-backend/controllers"
	"ecommerce-backend/services"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	userController := controllers.UserController{
		Service: &services.UserService{},
	}

	app.Post("/users", userController.CreateUser)
}