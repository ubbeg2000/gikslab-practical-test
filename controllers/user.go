package controllers

import "github.com/gofiber/fiber/v2"

func RegisterUserRoutes(api fiber.Router) {
	r := api.Group("/user")

	r.Post("", func(c *fiber.Ctx) error {
		return nil
	})
}
