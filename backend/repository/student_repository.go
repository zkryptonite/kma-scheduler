package repository

import (
	"backend/entity"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type StudentRepository struct {
	Db *mongo.Database
}

func NewStudentRepository(db *mongo.Database) IStudentRepository {
	return &StudentRepository{Db: db}
}

func (mongo *StudentRepository) Insert(student entity.Student) error {
	var _, err = mongo.Db.Collection("Student").
		InsertOne(context.Background(), student)
	if err != nil {
		return err
	}
	return nil
}

func (mongo *StudentRepository) InsertEmail(username, email string) error {
	update := bson.M{"$set": bson.M{"email": email}}

	_, err := mongo.Db.Collection("Student").
		UpdateOne(context.Background(),
			bson.M{"username": username}, update)

	if err != nil {
		return err
	}

	return nil
}

func (mongo *StudentRepository) FindStudentByUsername(username string) (entity.Student, error) {
	student := entity.Student{}
	singleRes := mongo.Db.Collection("Student").
		FindOne(context.Background(),
			bson.M{"username": username})
	err := singleRes.Decode(&student)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (mongo *StudentRepository) FindStudentByEmail(email string) (entity.Student, error) {
	student := entity.Student{}
	singleRes := mongo.Db.Collection("Student").
		FindOne(context.Background(),
			bson.M{"email": email})
	err := singleRes.Decode(&student)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (mongo *StudentRepository) FindStudentsWithEmail() ([]entity.Student, error) {
	students := []entity.Student{}
	cursor, err := mongo.Db.Collection("Student").
		Find(context.Background(),
			bson.M{"email": bson.M{"$ne": ""}})

	if err != nil {
		return students, err
	}

	err = cursor.All(context.Background(), &students)

	if err != nil {
		return students, err
	}

	return students, nil
}

func (mongo *StudentRepository) CheckLoginInfo(username string) (entity.Student, error) {

	student := entity.Student{}

	singleRes := mongo.Db.Collection("Student").
		FindOne(context.Background(),
			bson.M{"username": username})

	err := singleRes.Decode(&student)

	if err != nil {
		return student, err
	}

	return student, nil
}

func (mongo *StudentRepository) FindAll(username string) []string {
	student, _ := mongo.FindStudentByUsername(username)
	return student.Class
}
