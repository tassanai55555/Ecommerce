package controllers

import (
    "github.com/gofiber/fiber/v2"
    "github.com/tassanai55555/media-city-backend/config"
    "github.com/tassanai55555/media-city-backend/models"
)

func CreateUser(c *fiber.Ctx) error {
    user := new(models.User)

    if err := c.BodyParser(user); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "cannot parse JSON"})
    }

    if err := config.DB.Create(&user).Error; err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }

    return c.Status(201).JSON(user)
}
