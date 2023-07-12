package controllers

import (
	"gikslab-practical-test/dto"
	"gikslab-practical-test/helpers"
	"gikslab-practical-test/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterAuthRoutes(api fiber.Router) {
	r := api.Group("/auth")

	r.Post("/login", func(c *fiber.Ctx) error {
		var body dto.LoginBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		if err := body.Validate(); err != nil {
			return c.Status(400).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		res, err := services.Login(body)
		if err != nil {
			return c.Status(401).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		return c.Status(200).JSON(res)
	})

	r.Get("/logout", func(c *fiber.Ctx) error {
		claims, err := helpers.GetClaims(c)
		if err != nil {
			return c.Status(401).JSON(dto.BaseResponse{
				Message: "unauthorized user",
			})
		}

		helpers.BlacklistToken(claims)

		return c.Status(200).JSON(dto.BaseResponse{
			Message: "logout success",
		})
	})
}
