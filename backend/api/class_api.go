package api

import (
	"backend/driver"
	"backend/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

type responseObject struct {
	Classname   string `json:"class_name"`
	Subjectname string `json:"subject_name"`
	Lecturer    string `json:"lecturer"`
	Room        string `json:"room"`
	Lesson      string `json:"lesson"`
}

func GetClasses(ctx *fiber.Ctx) error {
	now := ctx.Query("date")

	db := driver.ConnectMongoDB().Client.Database(os.Getenv("DB_NAME"))
	classRepository := repository.NewClassRepository(db)
	studentRepository := repository.NewStudentRepository(db)

	user := ctx.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	allClassesOfStudent := studentRepository.FindAll(username)
	requiredClasses, _ := classRepository.GetScheduleOfClassesByDay(allClassesOfStudent, now)

	var resObjs []responseObject

	for _, c := range requiredClasses {
		resObj := responseObject{
			Classname:   c.ClassName,
			Subjectname: c.SubjectName,
			Lecturer:    c.Lecturer,
			Room:        c.Room,
		}

		for _, s := range c.Schedule {
			resObj.Lesson = s
		}

		resObjs = append(resObjs, resObj)
	}

	return ctx.Status(fiber.StatusOK).JSON(resObjs)
}
