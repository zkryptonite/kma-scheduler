package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Student struct {
	ID       primitive.ObjectID `json:"-" bson:"_id,omitempty"`
	UserName string             `json:"username" bson:"username"`
	Password string             `json:"password" bson:"password"`
	Email    string             `json:"email" bson:"email"`
	Class    []string           `json:"class" bson:"class"`
}

func NewStudent(username, password, email string, class []string) *Student {
	return &Student{
		UserName: username,
		Password: password,
		Email: email,
		Class: class,
	}
}