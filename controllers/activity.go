package controllers

import (
	"gikslab-practical-test/dto"
	"gikslab-practical-test/middlewares"
	"gikslab-practical-test/services"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func RegisterActivityRoutes(api fiber.Router) {
	r := api.Group("/activity")

	r.Post("", middlewares.AllowProfiles("expert"), func(c *fiber.Ctx) error {
		var body dto.RegisterActivityBody
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

		if err := services.RegisterActivity(body); err != nil {
			return c.Status(422).JSON(dto.BaseResponse{
				Message: "data cannot be processed",
			})
		}

		return c.Status(200).JSON(dto.BaseResponse{
			Message: "create success",
		})
	})

	r.Patch("/:activity_id", middlewares.AllowProfiles("expert"), func(c *fiber.Ctx) error {
		activityID, err := strconv.ParseUint(c.Params("activity_id", ""), 10, 64)
		if err != nil {
			return c.Status(400).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		var body dto.UpdateActivityBody
		if err := c.BodyParser(&body); err != nil {
			return c.Status(400).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		if err := services.UpdateActivity(activityID, body); err != nil {
			return c.Status(422).JSON(dto.BaseResponse{
				Message: "data cannot be processed",
			})
		}

		return c.Status(200).JSON(dto.BaseResponse{
			Message: "update success",
		})
	})

	r.Delete("/:activity_id", middlewares.AllowProfiles("expert"), func(c *fiber.Ctx) error {
		activityID, err := strconv.ParseUint(c.Params("activity_id", ""), 10, 64)
		if err != nil {
			return c.Status(400).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		if err := services.DeleteActivity(activityID); err != nil {
			return c.Status(422).JSON(dto.BaseResponse{
				Message: "data cannot be processed",
			})
		}

		return c.Status(200).JSON(dto.BaseResponse{
			Message: "delete success",
		})
	})

	r.Get("/:skill_id", middlewares.TokenRequired, func(c *fiber.Ctx) error {
		sortBy := c.Query("sort_by", "startdate")
		sortOrder := c.Query("sort_order", "asc")
		page := c.QueryInt("page", 1)
		limit := c.QueryInt("limit", 20)

		if sortBy == "startdate" {
			sortBy = "start_date"
		}

		if sortBy == "enddate" {
			sortBy = "end_date"
		}

		skillID, err := strconv.ParseUint(c.Params("skill_id", ""), 10, 64)
		if err != nil {
			return c.Status(400).JSON(dto.BaseResponse{
				Message: err.Error(),
			})
		}

		return c.Status(200).JSON(services.ListActivities(skillID, page, limit, sortBy, sortOrder))
	})
}
