package controllers

import (
	"gikslab-practical-test/dto"
	"gikslab-practical-test/middlewares"
	"gikslab-practical-test/services"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(api fiber.Router) {
	r := api.Group("/user")

	r.Post("", middlewares.AllowProfiles("expert"), func(c *fiber.Ctx) error {
		var body dto.RegistrationBody
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

		if err := services.RegisterUser(body); err != nil {
			return c.Status(422).JSON(dto.BaseResponse{
				Message: "data cannot be processed",
			})
		}

		return c.Status(200).JSON(dto.BaseResponse{
			Message: "create success",
		})
	})
}
