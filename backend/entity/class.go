package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Class struct {
	ID          primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	ClassName   string             `json:"class_name" bson:"class_name"`
	SubjectName string             `json:"subject_name" bson:"subject_name"`
	Room        string             `json:"room" bson:"room"`
	Lecturer    string             `json:"lecturer" bson:"lecturer"`
	Schedule    map[string]string  `json:"schedule" bson:"schedule"`
}

func NewClass(className, subjectName, room, lecturer string, schedule map[string]string) *Class {
	return &Class{
		ClassName:   className,
		SubjectName: subjectName,
		Room:        room,
		Lecturer:    lecturer,
		Schedule:    schedule,
	}
}
