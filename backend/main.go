package main

import (
	"backend/api"
	"backend/middleware"
	"backend/notifier"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	var app = fiber.New()

	app.Use(cors.New())

	app.Post("/login", middleware.Authenicate)

	app.Post("/api/student/email", middleware.Authorize(), api.AddEmail)

	app.Get("/api/classes", middleware.Authorize(), api.GetClasses)

	app.Use("/*", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "not found",
		})
	})

	go notifier.Run()
	app.Listen(":5000")
}
