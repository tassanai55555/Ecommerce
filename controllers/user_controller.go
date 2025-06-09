package controllers

import (
	"ecommerce-backend/dto"
	"ecommerce-backend/services"

	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	Service *services.UserService
}

func (c *UserController) CreateUser(ctx *fiber.Ctx) error {
	var body dto.CreateUserRequest

	if err := ctx.BodyParser(&body); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	user, err := c.Service.CreateUser(body)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(user)
}