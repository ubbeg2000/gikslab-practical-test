package controllers

import (
	"gikslab-practical-test/middlewares"
	"gikslab-practical-test/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterSkillRoutes(api fiber.Router) {
	r := api.Group("/skill")

	r.Get("", middlewares.TokenRequired, func(c *fiber.Ctx) error {
		return c.Status(200).JSON(services.ListSkills())
	})
}
