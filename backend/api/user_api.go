package api

import (
	"backend/driver"
	"backend/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"os"
)

func AddEmail(ctx *fiber.Ctx) error {
	request := struct{ Email string }{}

	err := ctx.BodyParser(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	email := request.Email
	db := driver.ConnectMongoDB().Client.Database(os.Getenv("DB_NAME"))
	studentRepository := repository.NewStudentRepository(db)

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)

	err = studentRepository.InsertEmail(username, email)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Add email failed",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": "Email added successfully",
	})
}
