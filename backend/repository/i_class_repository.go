package repository

import "backend/entity"

type IClassRepository interface {
	GetScheduleOfClassesByDay(classNames []string, day string) ([]entity.Class, error)
	Insert(classes ...entity.Class) error
	FindClass(className string) (entity.Class, error)
}
