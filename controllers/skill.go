package controllers

import "github.com/gofiber/fiber/v2"

func RegisterSkillRoutes(api fiber.Router) {
	r := api.Group("/skill")

	r.Get("", func(c *fiber.Ctx) error {
		return nil
	})
}
