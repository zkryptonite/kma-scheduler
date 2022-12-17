package notifier

import (
	"backend/driver"
	"backend/repository"
	"time"
)

type class struct {
	Subject string `json:"subject"`
	Room    string `json:"room"`
	Lesson  string `json:"lesson"`
}

func createBody(email string) map[string]interface{} {
	var (
		now               = time.Now().Format("02/01/2006")
		db                = driver.ConnectMongoDB().Client.Database(dbName)
		studentRepository = repository.NewStudentRepository(db)
		classRepository   = repository.NewClassRepository(db)
		std, _            = studentRepository.FindStudentByEmail(email)
		classes, _        = classRepository.GetScheduleOfClassesByDay(std.Class, now)
	)

	if len(classes) == 0 {
		return map[string]interface{}{}
	}

	var cls []class

	for _, c := range classes {
		newClass := class{
			Subject: c.SubjectName,
			Room:    c.Room,
			Lesson:  c.Schedule[now],
		}
		cls = append(cls, newClass)
	}

	return map[string]interface{}{"classes": cls}
}
