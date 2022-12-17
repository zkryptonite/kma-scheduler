package repository

import (
	"backend/entity"
)

type IStudentRepository interface {
	FindStudentByUsername(username string) (entity.Student, error)
	FindStudentByEmail(username string) (entity.Student, error)
	FindStudentsWithEmail() ([]entity.Student, error)
	CheckLoginInfo(username string) (entity.Student, error)
	Insert(std entity.Student) error
	FindAll(username string) []string
	InsertEmail(username, email string) error
}
