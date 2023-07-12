package middlewares

import (
	"gikslab-practical-test/dto"
	"gikslab-practical-test/helpers"

	"github.com/gofiber/fiber/v2"
)

func TokenRequired(c *fiber.Ctx) error {
	token := c.Query("token", "")
	if token == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
			Message: "unauthorized user",
		})
	}

	_, err := helpers.ParseToken(token)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
			Message: "unauthorized user",
		})
	}

	return c.Next()
}

func AllowProfiles(allowedProfiles ...string) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := c.Query("token", "")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
				Message: "unauthorized user",
			})
		}

		claims, err := helpers.ParseToken(token)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
				Message: "unauthorized user",
			})
		}

		if _, ok := helpers.TokenBlacklist[claims.Id]; ok {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
				Message: "unauthorized user",
			})
		}

		ok := false
		for _, p := range allowedProfiles {
			if p == claims.UserProfile {
				ok = true
				break
			}
		}

		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(dto.BaseResponse{
				Message: "unauthorized user",
			})
		}

		return c.Next()
	}
}
