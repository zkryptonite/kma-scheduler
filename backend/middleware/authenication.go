package middleware

import (
	"backend/collector"
	"backend/driver"
	"backend/entity"
	"backend/repository"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"time"
)

type request struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

var (
	accessTokenSecret = os.Getenv("ACCESS_TOKEN_SECRET")
	dbName            = os.Getenv("DB_NAME")
)

func Authenicate(ctx *fiber.Ctx) error {
	var r request

	err := ctx.BodyParser(&r)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse json",
		})
	}

	db := driver.ConnectMongoDB().Client.Database(dbName)
	studentRepository := repository.NewStudentRepository(db)
	student, err := studentRepository.CheckLoginInfo(r.UserName)

	if err != nil {
		student, err = loginToHomePage(r.UserName, r.Password)
		if err != nil {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Bad credentials",
			})
		}
	}

	if !checkPasswordHash(r.Password, student.Password) {
		return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad credentials",
		})
	}

	tokenString, err := generateAccessToken(student)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Server error",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"access_token": tokenString,
	})
}

func loginToHomePage(username, password string) (entity.Student, error) {
	var student entity.Student
	cookies, err := collector.GetCookies(username, password)

	if err != nil || len(cookies) == 1 {
		return student, fmt.Errorf("LOGIN FAILED")
	}

	encryptedPassword, _ := hashPassword(password)

	student = entity.Student{
		UserName: username,
		Password: encryptedPassword,
	}

	db := driver.ConnectMongoDB().Client.Database(dbName)
	studentRepository := repository.NewStudentRepository(db)
	classRepository := repository.NewClassRepository(db)

	html, err := collector.GetRequiredHtmlFromAccount(cookies)
	if err != nil {
		return student, fmt.Errorf("HTML GET FAILED")
	}
	classes := collector.ParseHtmlToClassObjects(html)

	classRepository.Insert(classes...)

	for _, cl := range classes {
		student.Class = append(student.Class, cl.ClassName)
	}

	studentRepository.Insert(student)
	return student, nil
}

func generateAccessToken(student entity.Student) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["username"] = student.UserName
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()

	accessToken, err := token.SignedString([]byte(accessTokenSecret))
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
