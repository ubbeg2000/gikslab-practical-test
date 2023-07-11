package controllers

import "github.com/gofiber/fiber/v2"

func RegisterAuthRoutes(api fiber.Router) {
	r := api.Group("/auth")

	r.Post("/login", func(c *fiber.Ctx) error {
		return nil
	})

	r.Get("/logout", func(c *fiber.Ctx) error {
		return nil
	})
}
