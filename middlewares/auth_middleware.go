package middleware

import "github.com/gofiber/fiber/v2"

func AllowProfiles() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return nil
	}
}
