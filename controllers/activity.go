package controllers

import "github.com/gofiber/fiber/v2"

func RegisterActivityRoutes(api fiber.Router) {
	r := api.Group("/activity")

	r.Post("", func(c *fiber.Ctx) error {
		return nil
	})

	r.Patch("", func(c *fiber.Ctx) error {
		return nil
	})

	r.Delete("", func(c *fiber.Ctx) error {
		return nil
	})

	r.Get("", func(c *fiber.Ctx) error {
		return nil
	})
}
