package main

import (
	"gikslab-practical-test/bootstrap"
	"gikslab-practical-test/controllers"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	bootstrap.Run()

	app := fiber.New()
	api := app.Group("/v1", cors.New())

	controllers.RegisterAuthRoutes(api)

	if err := app.Listen(os.Getenv("APP_HOST") + ":" + os.Getenv("APP_PORT")); err != nil {
		panic(err.Error())
	}
}
